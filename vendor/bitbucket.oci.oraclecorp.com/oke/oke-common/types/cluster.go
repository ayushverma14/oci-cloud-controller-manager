package types

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	// ClusterStateForAPICreating is the creating lifecycle state used in API responses
	ClusterStateForAPICreating = "CREATING"
	// ClusterStateForAPIActive is the active lifecycle state used in API responses
	ClusterStateForAPIActive = "ACTIVE"
	// ClusterStateForAPIFailed is the failed lifecycle state used in API responses
	ClusterStateForAPIFailed = "FAILED"
	// ClusterStateForAPIDeleting is the deleting lifecycle state used in API responses
	ClusterStateForAPIDeleting = "DELETING"
	// ClusterStateForAPIDeleted is the deleted lifecycle state used in API responses
	ClusterStateForAPIDeleted = "DELETED"
	// ClusterStateForAPIUpdating is the updating lifecycle state used in API responses
	ClusterStateForAPIUpdating = "UPDATING"
)

// ClusterSummaryV3 is an array element in the response type for the ListClusters API operation.
type ClusterSummaryV3 struct {
	ID                          string                  `json:"id"`
	Name                        string                  `json:"name"`
	CompartmentID               string                  `json:"compartmentId"`
	VCNID                       string                  `json:"vcnId"`
	KubernetesVersion           string                  `json:"kubernetesVersion"`
	Options                     *ClusterCreateOptionsV3 `json:"options,omitempty"`
	Metadata                    map[string]string       `json:"metadata"`
	LifecycleState              string                  `json:"lifecycleState"`
	LifecycleDetails            string                  `json:"lifecycleDetails"`
	Endpoints                   *ClusterEndpointsV3     `json:"endpoints"`
	AvailableKubernetesUpgrades []string                `json:"availableKubernetesUpgrades"`
	ImagePolicyConfig           *ImagePolicyConfigV3    `json:"imagePolicyConfig, omitempty"`
}

// ClusterV3 is the response type for the GetCluster API operation.
type ClusterV3 struct {
	ID                          string                  `json:"id"`
	Name                        string                  `json:"name"`
	CompartmentID               string                  `json:"compartmentId"`
	VCNID                       string                  `json:"vcnId"`
	KubernetesVersion           string                  `json:"kubernetesVersion"`
	KMSKeyID                    string                  `json:"kmsKeyId,omitempty"`
	Options                     *ClusterCreateOptionsV3 `json:"options,omitempty"`
	Metadata                    map[string]string       `json:"metadata"`
	LifecycleState              string                  `json:"lifecycleState"`
	LifecycleDetails            string                  `json:"lifecycleDetails"`
	Endpoints                   *ClusterEndpointsV3     `json:"endpoints"`
	AvailableKubernetesUpgrades []string                `json:"availableKubernetesUpgrades"`
	ImagePolicyConfig           *ImagePolicyConfigV3    `json:"imagePolicyConfig, omitempty"`
}

// ToSummaryV3 a K8Instance object to a ClusterSummaryV3 object understood by the higher layers
func (src *K8Instance) ToSummaryV3(exposePSP bool) *ClusterSummaryV3 {
	var dst ClusterSummaryV3
	if src == nil {
		return &dst
	}

	dst.ID = src.ID
	dst.Name = src.Name
	dst.CompartmentID = src.CompartmentID
	dst.VCNID = src.NetworkConfig.VCNID
	dst.KubernetesVersion = src.K8Version

	if src.ImagePolicyConfig != nil {
		dst.ImagePolicyConfig = &ImagePolicyConfigV3{
			IsPolicyEnabled: src.ImagePolicyConfig.IsPolicyEnabled,
		}
		dst.ImagePolicyConfig.KeyDetails = make([]*KeyDetailsV3, len(src.ImagePolicyConfig.KeyDetails))
		for idx, srcKeyObj := range src.ImagePolicyConfig.KeyDetails {
			dst.ImagePolicyConfig.KeyDetails[idx] = &KeyDetailsV3{
				KmsKeyId:	srcKeyObj.KmsKeyId,
			}
		}
	}

	dst.Options = &ClusterCreateOptionsV3{
		KubernetesNetworkConfig: &KubernetesNetworkConfigV3{
			PodsCIDR:     src.NetworkConfig.K8SPodsCIDR,
			ServicesCIDR: src.NetworkConfig.K8SServicesCIDR,
		},
	}
	if src.InstallOptions != nil {
		dst.Options.AddOns = &AddOnOptionsV3{
			Tiller:              src.InstallOptions.HasTiller,
			KubernetesDashboard: src.InstallOptions.HasKubernetesDashboard,
		}
		if exposePSP {
			dst.Options.AdmissionControllerOptions = &AdmissionControllersV3{
				PodSecurityEnabled: src.InstallOptions.PodSecurityPolicy != PodSecurityPolicy_Disabled,
			}
		}
	}

	dst.Options.ServiceLBSubnetIDs = make([]string, len(src.NetworkConfig.ServiceLBSubnets))
	for idx, sn := range src.NetworkConfig.ServiceLBSubnets {
		dst.Options.ServiceLBSubnetIDs[idx] = sn
	}
	dst.Metadata = src.Metadata
	dst.LifecycleState = src.TKMState.ToAPIV3()
	// dst.LifecycleDetails = "FIXME"

	dst.Endpoints = &ClusterEndpointsV3{
		Kubernetes: src.K8Addr,
	}

	dst.AvailableKubernetesUpgrades = []string{}
	for _, k8Ver := range src.AvailableK8SUpgrades {
		dst.AvailableKubernetesUpgrades = append(dst.AvailableKubernetesUpgrades, k8Ver)
	}

	return &dst
}

// ToV3 converts a K8Instance object to a ClusterV3 object understood by the higher layers
func (src *K8Instance) ToV3(exposePSP bool) *ClusterV3 {
	var dst ClusterV3
	if src == nil {
		return &dst
	}

	dst.ID = src.ID
	dst.Name = src.Name
	dst.CompartmentID = src.CompartmentID
	dst.VCNID = src.NetworkConfig.VCNID
	dst.KubernetesVersion = src.K8Version
	dst.KMSKeyID = src.KMSKeyID

	if src.ImagePolicyConfig != nil {
		dst.ImagePolicyConfig = &ImagePolicyConfigV3{
			IsPolicyEnabled: src.ImagePolicyConfig.IsPolicyEnabled,
		}

		dst.ImagePolicyConfig.KeyDetails = make([]*KeyDetailsV3, len(src.ImagePolicyConfig.KeyDetails))
		for idx, srcKeyObj := range src.ImagePolicyConfig.KeyDetails {
			dst.ImagePolicyConfig.KeyDetails[idx] = &KeyDetailsV3{
				KmsKeyId:	srcKeyObj.KmsKeyId,
			}
		}
	}

	dst.Options = &ClusterCreateOptionsV3{
		KubernetesNetworkConfig: &KubernetesNetworkConfigV3{
			PodsCIDR:     src.NetworkConfig.K8SPodsCIDR,
			ServicesCIDR: src.NetworkConfig.K8SServicesCIDR,
		},
		AddOns: &AddOnOptionsV3{
			Tiller:              src.InstallOptions.HasTiller,
			KubernetesDashboard: src.InstallOptions.HasKubernetesDashboard,
		},
	}

	if exposePSP {
		dst.Options.AdmissionControllerOptions = &AdmissionControllersV3{
			PodSecurityEnabled: src.InstallOptions.PodSecurityPolicy != PodSecurityPolicy_Disabled,
		}
	}

	dst.Options.ServiceLBSubnetIDs = make([]string, len(src.NetworkConfig.ServiceLBSubnets))
	for idx, sn := range src.NetworkConfig.ServiceLBSubnets {
		dst.Options.ServiceLBSubnetIDs[idx] = sn
	}
	dst.Metadata = src.Metadata
	dst.LifecycleState = src.TKMState.ToAPIV3()
	// dst.LifecycleDetails = "FIXME"

	dst.Endpoints = &ClusterEndpointsV3{
		Kubernetes: src.K8Addr,
	}

	dst.AvailableKubernetesUpgrades = []string{}
	for _, k8Ver := range src.AvailableK8SUpgrades {
		dst.AvailableKubernetesUpgrades = append(dst.AvailableKubernetesUpgrades, k8Ver)
	}

	return &dst
}

// ToAPIV3 converts the state to valid string used by the API
func (s TKMState) ToAPIV3() string {
	switch s {
	case TKMState_Initializing:
		return ClusterStateForAPICreating
	case TKMState_Running:
		return ClusterStateForAPIActive
	case TKMState_Failed:
		return ClusterStateForAPIFailed
	case TKMState_Terminating:
		return ClusterStateForAPIDeleting
	case TKMState_Terminated:
		return ClusterStateForAPIDeleted
	case TKMState_Updating_Masters:
		return ClusterStateForAPIUpdating
	default:
		return "UNKNOWN"
	}
}

func FromClusterLifeCycleState(lifecycle string) (TKMState, error) {

	switch strings.ToUpper(lifecycle) {
	case ClusterStateForAPICreating:
		return TKMState_Initializing, nil
	case ClusterStateForAPIActive:
		return TKMState_Running, nil
	case ClusterStateForAPIFailed:
		return TKMState_Failed, nil
	case ClusterStateForAPIDeleting:
		return TKMState_Terminating, nil
	case ClusterStateForAPIDeleted:
		return TKMState_Terminated, nil
	case ClusterStateForAPIUpdating:
		return TKMState_Updating_Masters, nil
	default:
		return -1, errors.New(fmt.Sprintf("Unknown lifecycle %s", lifecycle))
	}
}

// ClusterEndpointsV3 contains the endpoint URLs for the Kubernetes API server.
type ClusterEndpointsV3 struct {
	Kubernetes string `json:"kubernetes"`
}

// CreateClusterDetailsV3 is the request body type for the CreateCluster API operation.
type CreateClusterDetailsV3 struct {
	Name              string                  `json:"name" yaml:"name"`
	KubernetesVersion string                  `json:"kubernetesVersion" yaml:"kubernetesVersion"`
	CompartmentID     string                  `json:"compartmentId" yaml:"compartmentId"`
	VCNID             string                  `json:"vcnId" yaml:"vcnId"`
	KMSKeyID          string                  `json:"kmsKeyId,omitempty" yaml:"kmsKeyId,omitempty"`
	Options           *ClusterCreateOptionsV3 `json:"options,omitempty" yaml:"options,omitempty"`
	ImagePolicyConfig *ImagePolicyConfigV3	  `json:"imagePolicyConfig, omitempty" yaml:"imagePolicyConfig,omitempty"`
}

// ImagePolicyConfigV3 defines the Image signature verification properties
type ImagePolicyConfigV3 struct {
	IsPolicyEnabled bool     `json:"isPolicyEnabled,omitempty" yaml:"isPolicyEnabled"`
	KeyDetails       []*KeyDetailsV3 `json:"keyDetails,omitempty" yaml:"keyDetails"`
}

type KeyDetailsV3 struct {
	KmsKeyId     string `json:"kmsKeyId,omitempty" yaml:"kmsKeyId"`
}

// ClusterCreateOptionsV3 defines the options that can modify how a cluster is created.
type ClusterCreateOptionsV3 struct {
	ServiceLBSubnetIDs         []string                   `json:"serviceLbSubnetIds" yaml:"serviceLbSubnetIds"`
	KubernetesNetworkConfig    *KubernetesNetworkConfigV3 `json:"kubernetesNetworkConfig,omitempty" yaml:"kubernetesNetworkConfig,omitempty"`
	AddOns                     *AddOnOptionsV3            `json:"addOns,omitempty" yaml:"addOns,omitempty"`
	AdmissionControllerOptions *AdmissionControllersV3    `json:"admissionControllerOptions,omitempty" yaml:"admissionControllerOptions,omitempty"`
}

// KubernetesNetworkConfigV3 defines the networking to use in creating a cluster
type KubernetesNetworkConfigV3 struct {
	PodsCIDR     string `json:"podsCidr" yaml:"podsCidr"`
	ServicesCIDR string `json:"servicesCidr" yaml:"servicesCidr"`
}

// AddOnOptionsV3 defines the add-ons to use in creating a cluster
type AddOnOptionsV3 struct {
	Tiller              bool `json:"isTillerEnabled" yaml:"isTillerEnabled"`
	KubernetesDashboard bool `json:"isKubernetesDashboardEnabled" yaml:"isKubernetesDashboardEnabled"`
}

// AdmissionControllersV3 defines the extended admission controllers to use in creating/updating a cluster
type AdmissionControllersV3 struct {
	PodSecurityEnabled bool `json:"isPodSecurityPolicyEnabled" yaml:"isPodSecurityPolicyEnabled"`
}

// ClusterCreateCLIResponseV3 is used by the CLI when creating a cluster
type ClusterCreateCLIResponseV3 struct {
	WorkRequestID string `json:"workRequestId"`
}

// UpdateClusterDetailsV3 is the request body type for the UpdateCluster API operation.
type UpdateClusterDetailsV3 struct {
	Name              string                `json:"name"`
	KubernetesVersion string                `json:"kubernetesVersion"`
	Options           *ClusterUpdateOptions `json:"options,omitempty" yaml:"options,omitempty"`
	ImagePolicyConfig *ImagePolicyConfigV3	`json:"imagePolicyConfig, omitempty" yaml:"imagePolicyConfig,omitempty"`
}

type ClusterUpdateOptions struct {
	AdmissionControllerOptions *AdmissionControllersV3 `json:"admissionControllerOptions,omitempty" yaml:"admissionControllerOptions,omitempty"`
}

func getK8SSemverNumbers(version string) ([]int64, error) {
	numbers := make([]int64, 3, 3)
	semvers := strings.Split(strings.TrimPrefix(version, "v"), ".")
	for i, v := range semvers {
		n, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("Incompatible version supplied")
		}
		numbers[i] = n
	}
	return numbers, nil
}

// ClusterOptionsV3 contains the options available for specific fields that can be submitted
// to the CreateCluster operation.
type ClusterOptionsV3 struct {
	KubernetesVersions []string `json:"kubernetesVersions"`
}

// ClusterDeleteCLIResponseV3 is used by the CLI when deleting a cluster
type ClusterDeleteCLIResponseV3 struct {
	WorkRequestID string `json:"workRequestId"`
}

// APIOptionsV3 is a type to hold all of the node pools api options
type APIOptionsV3 struct {
	KubernetesVersions []string `json:"kubernetesVersions"`
}

// ClusterUpdateCLIResponseV3 is used by the CLI when updating a cluster
type ClusterUpdateCLIResponseV3 struct {
	WorkRequestID string `json:"workRequestId"`
}
