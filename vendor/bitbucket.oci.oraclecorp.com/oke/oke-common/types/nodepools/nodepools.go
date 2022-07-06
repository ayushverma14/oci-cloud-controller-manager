package nodepools

import (
	fmt "fmt"
	"net/http"
	"regexp"
	"strings"

	"bitbucket.oci.oraclecorp.com/oke/oke-common/types/apierrors"
	"golang.org/x/crypto/ssh"
)

const (
	// UpdateFieldName is the affected fields key to indicate that the name needs updating
	UpdateFieldName = "Name"
	// UpdateFieldK8SVersion is the affected fields key to indicate that the K8SVersion needs updating
	UpdateFieldK8SVersion = "K8SVersion"
	// UpdateFieldInitialNodeLabels is the affected fields key to indicate that the InitialNodeLabels needs updating
	UpdateFieldInitialNodeLabels = "InitialNodeLabels"
	// UpdateFieldQuantityPerSubnet is the affected fields key to indicate that the QuantityPerSubnet needs updating
	UpdateFieldQuantityPerSubnet = "QuantityPerSubnet"
	// UpdateFieldSubnetsInfo is the affected fields key to indicate that the SubnetsInfo needs updating
	UpdateFieldSubnetsInfo = "SubnetsInfo"
	// UpdateFieldSize is the affected fields key to indicate that the Size needs updating
	UpdateFieldSize = "Size"
	// UpdateFieldNsgIds is the affected fields key to indicate that the NsgIds needs updating
	UpdateFieldNsgIds = "NsgIds"
	// UpdateFieldNodeImageID is the affected fields key to indicate that the NodeImageID needs updating
	UpdateFieldNodeImageID = "NodeImageID"
	// UpdateFieldNodeImageName is the affected fields key to indicate that the NodeImageName needs updating
	UpdateFieldNodeImageName = "NodeImageName"
	// UpdateFieldNodeMetadata is the affected fields key to indicate that the NodeMetadata needs updating
	UpdateFieldNodeMetadata = "NodeMetadata"
	// UpdateFieldNodeBootVolumeSizeInGBs is the affected fields key to indicate that the NodeBootVolumeSizeInGBs needs updating
	UpdateFieldNodeBootVolumeSizeInGBs = "NodeBootVolumeSizeInGBs"
	// UpdateFieldSSHPublicKey is the affected fields key to indicate that the SSHPublicKey needs updating
	UpdateFieldSSHPublicKey = "SSHPublicKey"
	// UpdateFieldNodeShape is the affected fields key to indicate that the NodeShape needs updating
	UpdateFieldNodeShape = "NodeShape"
	// UpdateFieldNodeOcpus is the affected fields key to indicate that the NodeOcpus needs updating
	UpdateFieldNodeOcpus = "NodeOcpus"
	// UpdateFieldNodeOcpus is the affected fields key to indicate that the NodeMemoryInGBs needs updating
	UpdateFieldNodeMemoryInGBs = "NodeMemoryInGBs"

	// k8s labels https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#syntax-and-character-set
	maxLabelKeyPrefixLength = 253
	minLabelKeyNameLength   = 1
	maxLabelKeyNameLength   = 63
	maxLabelValueLength     = 63
	// MaxInitialNodePoolsLength is the max length of initial_node_pools
	MaxInitialNodePoolsLength = 65535
)

// NodePoolNewResponseV1 is the response from creating a new nodepool
type NodePoolNewResponseV1 struct {
	JobID      string `json:"workItemId"`
	NodePoolID string `json:"nodePoolId"`
}

// ToV1 converts a nodepool.NewResponse object to a NodePoolNewResponseV1 object understood by the higher layers
func (src *NewResponse) ToV1() NodePoolNewResponseV1 {
	var dst NodePoolNewResponseV1
	if src == nil {
		return dst
	}
	if src != nil {
		dst.JobID = src.JobID
		dst.NodePoolID = src.NodePool.ID
	}

	return dst
}

// ValidateInitialNodeLabels validates the initial node labels, e.g. k1=v1,k2=v2
func ValidateInitialNodeLabels(labels string, disallowedPrefixes []string) (int, *apierrors.ErrorV3) {
	if len(labels) <= 0 {
		return http.StatusOK, nil
	}

	// disallow invalid characters
	// https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/
	var validCharactersRegex = regexp.MustCompile("^[0-9a-zA-Z/,=._-]+$").MatchString
	if !validCharactersRegex(labels) {
		return http.StatusBadRequest, apierrors.NewErrorV3(
			apierrors.HTTP400InvalidParameterCode,
			apierrors.HTTP400InvalidParameterMessageInitialNodeLabelsChars)
	}

	kvs := KeyValuesFromString(labels)

	if equalCount := strings.Count(labels, "="); equalCount != len(kvs) {
		return http.StatusBadRequest, apierrors.NewErrorV3(
			apierrors.HTTP400InvalidParameterCode,
			apierrors.HTTP400InvalidParameterMessageInitialNodeLabelsEqualDelim)
	}

	for _, kv := range kvs {
		// disallow empty keys
		if len(kv.Key) <= 0 {
			return http.StatusBadRequest, apierrors.NewErrorV3(
				apierrors.HTTP400InvalidParameterCode,
				apierrors.HTTP400InvalidParameterMessageInitialNodeLabelsKeys)
		}

		// disallow invalid keys
		parts := strings.Split(kv.Key, "/")
		if len(parts) > 2 { // disallow kv.Key with more than one slash separator
			return http.StatusBadRequest, apierrors.NewErrorV3(
				apierrors.HTTP400InvalidParameterCode,
				apierrors.HTTP400InvalidParameterMessageInitialNodeLabelsKeys)
		} else if len(parts) > 1 { // kv.Key consists of prefix and name (Key: "prefix/name")
			// disallow key prefix length greater than maxLabelKeyPrefixLength
			if len(parts[0]) > maxLabelKeyPrefixLength {
				return http.StatusBadRequest, apierrors.NewErrorV3(
					apierrors.HTTP400InvalidParameterCode,
					apierrors.HTTP400InvalidParameterMessageInitialNodeLabelsKeys)
			}

			// disallow key name length less than minLabelKeyNameLength
			if len(parts[1]) < minLabelKeyNameLength {
				return http.StatusBadRequest, apierrors.NewErrorV3(
					apierrors.HTTP400InvalidParameterCode,
					apierrors.HTTP400InvalidParameterMessageInitialNodeLabelsKeys)
			}

			// disallow key name length greater than maxLabelKeyNameLength
			if len(parts[1]) > maxLabelKeyNameLength {
				return http.StatusBadRequest, apierrors.NewErrorV3(
					apierrors.HTTP400InvalidParameterCode,
					apierrors.HTTP400InvalidParameterMessageInitialNodeLabelsKeys)
			}

			// setup to disallow prefixes
			var disallowedPrefixesRegex func(b string) bool
			if len(disallowedPrefixes) > 0 {
				for _, disallowedPrefix := range disallowedPrefixes {
					disallowedPrefixesRegex = regexp.MustCompile(disallowedPrefix).MatchString
					// disallow prefixes for each Key's prefix, barring the specific ones allowed
					if disallowedPrefixesRegex != nil && disallowedPrefixesRegex(parts[0]) && !isAllowedSpecificPrefix(kv.Key) {
						return http.StatusBadRequest, apierrors.NewErrorV3(
							apierrors.HTTP400InvalidParameterCode,
							apierrors.HTTP400InvalidParameterMessageInitialNodeLabelsPrefix)
					}
				}
			}
			// disallow *node-restriction.kubernetes.io/ prefix
			if strings.HasSuffix(parts[0], "node-restriction.kubernetes.io") {
				return http.StatusBadRequest, apierrors.NewErrorV3(
					apierrors.HTTP400InvalidParameterCode,
					fmt.Sprintf(apierrors.HTTP400InvalidParameterMessageInitialNodeLabelsNodeRestrictionNotSupported, kv.Key))
			}
		} else { // kv.Key consists of name only (Key: "name")
			// disallow key name length less than minLabelKeyNameLength
			if len(kv.Key) < minLabelKeyNameLength {
				return http.StatusBadRequest, apierrors.NewErrorV3(
					apierrors.HTTP400InvalidParameterCode,
					apierrors.HTTP400InvalidParameterMessageInitialNodeLabelsKeys)
			}

			// disallow key name length greater than maxLabelKeyNameLength
			if len(kv.Key) > maxLabelKeyNameLength {
				return http.StatusBadRequest, apierrors.NewErrorV3(
					apierrors.HTTP400InvalidParameterCode,
					apierrors.HTTP400InvalidParameterMessageInitialNodeLabelsKeys)
			}
		}

		// disallow value length greater than maxLabelValueLength
		if len(kv.Value) > maxLabelValueLength {
			return http.StatusBadRequest, apierrors.NewErrorV3(
				apierrors.HTTP400InvalidParameterCode,
				apierrors.HTTP400InvalidParameterMessageInitialNodeLabelsValues)
		}
	}

	// let kubernetes deal with other issues like needing labels to start with alpha, ...

	return http.StatusOK, nil
}

// https://github.com/kubernetes/enhancements/blob/master/keps/sig-auth/0000-20170814-bounding-self-labeling-kubelets.md
func isAllowedSpecificPrefix(key string) bool {

	if regexp.MustCompile("^(.*[.])?kubelet.kubernetes.io/.*$").MatchString(key) || regexp.MustCompile("^(.*[.])?node.kubernetes.io/.*$").MatchString(key) {
		return true
	}

	switch key {
	case
		"kubernetes.io/hostname",
		"kubernetes.io/instance-type",
		"kubernetes.io/os",
		"kubernetes.io/arch",
		"beta.kubernetes.io/instance-type",
		"beta.kubernetes.io/os",
		"beta.kubernetes.io/arch",
		"failure-domain.beta.kubernetes.io/zone",
		"failure-domain.beta.kubernetes.io/region",
		"failure-domain.kubernetes.io/zone",
		"failure-domain.kubernetes.io/region":
		return true
	}
	return false
}

// ValidateSSHKey will return nil if the passed string is a valid
// key from an authorized_keys file used in OpenSSH
func ValidateSSHKey(s string) (int, *apierrors.ErrorV3) {
	if len(s) == 0 {
		return http.StatusOK, nil
	}
	bs := []byte(s)
	_, _, _, _, err := ssh.ParseAuthorizedKey(bs)
	if err != nil {
		return http.StatusBadRequest, apierrors.NewErrorV3(
			apierrors.HTTP400InvalidParameterCode,
			apierrors.HTTP400InvalidParameterMessageSSHKey)

	}
	return http.StatusOK, nil
}
