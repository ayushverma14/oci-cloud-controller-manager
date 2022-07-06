package versions

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/coreos/go-semver/semver"
	"github.com/pkg/errors"
)

// List of supported master and worker kubernetes versions
type K8sV1 struct {
	MastersK8sVersions []string `json:"masters" yaml:"masters"`
	NodesK8sVersions   []string `json:"nodes" yaml:"nodes"`
}

type EtcdKubernetesMapGetter interface {
	GetKubernetesEtcdVersionMap(ctx context.Context, tenancyID string) (map[string]string, error)
}

type KubernetesGetter interface {
	GetSupportedMasterKubernetes(ctx context.Context, tenancyID string) ([]string, error)
	GetSupportedWorkerKubernetes(ctx context.Context, tenancyID string) ([]string, error)
	// Makes no guarantee on sort order
	GetAllKubernetes(ctx context.Context, tenancyID string) ([]string, error)
	EtcdKubernetesMapGetter
}

// Sorts the output from the underlying KubernetesGetter call
func GetAllKubernetesSorted(ctx context.Context, tenancyID string, versionsGetter KubernetesGetter) ([]string, error) {
	versions, err := versionsGetter.GetAllKubernetes(ctx, tenancyID)
	if err != nil {
		return nil, errors.Wrap(err, "error fetching all k8 versions")
	}
	sort.Slice(versions, func(i, j int) bool {
		iKeySemVer, err := semver.NewVersion(strings.TrimLeft(versions[i], "v"))
		if err != nil {
			return false
		}
		jKeySemVer, err := semver.NewVersion(strings.TrimLeft(versions[j], "v"))
		if err != nil {
			return false
		}

		return iKeySemVer.LessThan(*jKeySemVer)
	})
	return versions, nil
}

// GetSupportedK8s returns lists of supported master and worker kubernetes versions
func GetSupportedK8s(ctx context.Context, tenancyID string, versionsGetter KubernetesGetter) (*K8sV1, error) {
	masterVersions, err := versionsGetter.GetSupportedMasterKubernetes(ctx, tenancyID)
	if err != nil {
		return nil, errors.Wrap(err, "error getting master versions")
	}
	workerVersions, err := versionsGetter.GetSupportedWorkerKubernetes(ctx, tenancyID)
	if err != nil {
		return nil, errors.Wrap(err, "error getting worker versions")
	}

	return &K8sV1{
		MastersK8sVersions: masterVersions,
		NodesK8sVersions:   workerVersions,
	}, nil
}

// Returns available master k8 versions, given a current version
//	* If current version equals the candidate version, don't include it
//	* If the the version is supported and adheres to the skew policy, include it
//	* If the version is deprecated and adheres to the skew policy, only include the latest patch version
func GetAvailableMasterK8s(ctx context.Context, tenancyID string, currentVersion string, versionsGetter KubernetesGetter) (*[]string, error) {
	supportedMasterVersions, err := versionsGetter.GetSupportedMasterKubernetes(ctx, tenancyID)
	if err != nil {
		return nil, errors.Wrap(err, "error getting master versions")
	}
	allKubernetesVersions, err := GetAllKubernetesSorted(ctx, tenancyID, versionsGetter)
	if err != nil {
		return nil, errors.Wrap(err, "error getting all versions")
	}

	availableMasterK8sVersions := []string{}
	currentSemver, err := semver.NewVersion(strings.TrimLeft(currentVersion, "v"))
	if err != nil {
		return nil, errors.Wrapf(err, "Unable to parse cluster K8s version %s as semantic version", currentVersion)
	}
	latestPatchVersions, err := getFilteredPatchVersions(allKubernetesVersions)
	if err != nil {
		return nil, errors.Wrapf(err, "Unable to get latest patch versions")
	}
	newestSupportedVersion, err := newestVersionInArray(supportedMasterVersions)
	if err != nil {
		return nil, errors.Wrapf(err, "Unable to get newest supported version")
	}
	newestSupportedSemVer, _ := semver.NewVersion(strings.TrimLeft(newestSupportedVersion, "v"))
	for _, candidateStrVer := range allKubernetesVersions {
		candidateSemver, err := semver.NewVersion(strings.TrimLeft(candidateStrVer, "v"))
		if err != nil {
			return nil, errors.Wrapf(err, "Unable to parse candidate version %s as semantic version", candidateStrVer)
		}

		// If we support the version and it adheres to the skew policy, include it
		if candidateSemver.Equal(*currentSemver) || newestSupportedSemVer.LessThan(*candidateSemver) {
			// Do not include
		} else if valueInArray(candidateStrVer, supportedMasterVersions) &&
			adheresToMasterK8Skew(*currentSemver, *candidateSemver) {
			availableMasterK8sVersions = append(availableMasterK8sVersions, candidateStrVer)
		} else {
			// Unsupported/older versions, do some additional checks/filters
			// Only pick up the latest patch versions that adhere to the skew policy and are greater than current version
			if currentSemver.LessThan(*candidateSemver) &&
				valueInArray(candidateStrVer, latestPatchVersions) &&
				adheresToMasterK8Skew(*currentSemver, *candidateSemver) {
				availableMasterK8sVersions = append(availableMasterK8sVersions, candidateStrVer)
			}
		}
	}
	return &availableMasterK8sVersions, nil
}

// Returns all available worker K8s with the following behaviors:
// 		* currentMasterVersion required - if only currentMasterVersion provided it will return a list of all worker versions compatible (N-2->N minor versions)
//			If any of the minor versions are deprecated, it will only return the latest patch version
//		* currentWorkerVersion optional - if provided it will filter out any versions that are less than or equal to that version
func GetAvailableWorkerK8s(ctx context.Context, tenancyID string, currentMasterVersion string, optCurrentWorkerVersion string, versionsGetter KubernetesGetter) (*[]string, error) {
	supportedWorkerVersions, err := versionsGetter.GetSupportedWorkerKubernetes(ctx, tenancyID)
	if err != nil {
		return nil, errors.Wrap(err, "error getting supported worker versions")
	}
	allKubernetesVersions, err := GetAllKubernetesSorted(ctx, tenancyID, versionsGetter)
	if err != nil {
		return nil, errors.Wrap(err, "error getting all versions")
	}

	currentMasterSemver, err := semver.NewVersion(strings.TrimLeft(currentMasterVersion, "v"))
	if err != nil {
		return nil, errors.Wrapf(err, "Unable to parse cluster K8s version %s as semantic version", currentMasterVersion)
	}

	latestPatchVersions, err := getFilteredPatchVersions(allKubernetesVersions)
	if err != nil {
		return nil, errors.Wrapf(err, "Unable to get latest patch versions")
	}

	availableWorkerK8sVersions := []string{}
	for _, candidateStrVer := range allKubernetesVersions {
		candidateSemver, err := semver.NewVersion(strings.TrimLeft(candidateStrVer, "v"))
		if err != nil {
			return nil, errors.Wrapf(err, "Unable to parse candidate version %s as semantic version", candidateStrVer)
		}

		// If we support the version and it adheres to the skew policy, include it
		if valueInArray(candidateStrVer, supportedWorkerVersions) &&
			adheresToWorkerK8Skew(*currentMasterSemver, *candidateSemver) {
			availableWorkerK8sVersions = append(availableWorkerK8sVersions, candidateStrVer)
		} else {
			// Unsupported/older versions, do some additional checks/filters
			// Only pick up the latest patch versions that adhere to the skew policy and are greater than current version
			if (candidateSemver.LessThan(*currentMasterSemver) || candidateSemver.Equal(*currentMasterSemver)) &&
				valueInArray(candidateStrVer, latestPatchVersions) &&
				adheresToWorkerK8Skew(*currentMasterSemver, *candidateSemver) {
				availableWorkerK8sVersions = append(availableWorkerK8sVersions, candidateStrVer)
			}
		}
	}
	if len(optCurrentWorkerVersion) > 0 {
		filteredWorkerVersions := []string{}
		currentWorkerSemVer, _ := semver.NewVersion(strings.TrimLeft(optCurrentWorkerVersion, "v"))
		for _, candidateStrVer := range availableWorkerK8sVersions {
			candidateSemver, err := semver.NewVersion(strings.TrimLeft(candidateStrVer, "v"))
			if err != nil {
				return nil, errors.Wrapf(err, "Unable to parse candidate version %s as semantic version", candidateStrVer)
			}
			if currentWorkerSemVer.LessThan(*candidateSemver) {
				filteredWorkerVersions = append(filteredWorkerVersions, candidateStrVer)
			}
		}
		return &filteredWorkerVersions, nil
	}
	return &availableWorkerK8sVersions, nil
}

func oldestVersionInArray(versionArray []string) (string, error) {
	if len(versionArray) == 0 {
		return "", errors.New("No versions were passed to oldestVersionInArray()")
	} else if len(versionArray) == 1 {
		return versionArray[0], nil
	}
	oldestVersion := versionArray[0]
	oldestSemVersion, err := semver.NewVersion(strings.TrimLeft(oldestVersion, "v"))
	if err != nil {
		return "", errors.Wrapf(err, "Unable to parse candidate version %s as semantic version", oldestVersion)
	}
	for i := 1; i < len(versionArray); i++ {
		candidateSemVersion, err := semver.NewVersion(strings.TrimLeft(versionArray[i], "v"))
		if err != nil {
			return "", errors.Wrapf(err, "Unable to parse candidate version %s as semantic version", versionArray[i])
		}
		if candidateSemVersion.LessThan(*oldestSemVersion) {
			oldestVersion = versionArray[i]
			oldestSemVersion = candidateSemVersion
		}
	}
	return oldestVersion, nil
}

func newestVersionInArray(versionArray []string) (string, error) {
	if len(versionArray) == 0 {
		return "", errors.New("No versions were passed to newestVersionInArray()")
	} else if len(versionArray) == 1 {
		return versionArray[0], nil
	}
	newestVersion := versionArray[0]
	newestSemVersion, err := semver.NewVersion(strings.TrimLeft(newestVersion, "v"))
	if err != nil {
		return "", errors.Wrapf(err, "Unable to parse candidate version %s as semantic version", newestVersion)
	}
	for i := 1; i < len(versionArray); i++ {
		candidateSemVersion, err := semver.NewVersion(strings.TrimLeft(versionArray[i], "v"))
		if err != nil {
			return "", errors.Wrapf(err, "Unable to parse candidate version %s as semantic version", versionArray[i])
		}
		if newestSemVersion.LessThan(*candidateSemVersion) {
			newestVersion = versionArray[i]
			newestSemVersion = candidateSemVersion
		}
	}
	return newestVersion, nil
}

func valueInArray(value string, array []string) bool {
	for _, v := range array {
		if v == value {
			return true
		}
	}
	return false
}

func adheresToMasterK8Skew(currentSemver semver.Version, candidateSemver semver.Version) bool {
	minorVersionSkew := candidateSemver.Minor - currentSemver.Minor
	majorVersionSkew := candidateSemver.Major - currentSemver.Major
	if (currentSemver.LessThan(candidateSemver) || currentSemver.Equal(candidateSemver)) &&
		minorVersionSkew <= 1 &&
		majorVersionSkew == 0 {
		return true
	}
	return false
}

func adheresToWorkerK8Skew(clusterSemVer semver.Version, candidateWorkerSemver semver.Version) bool {
	minorVersionSkew := clusterSemVer.Minor - candidateWorkerSemver.Minor
	majorVersionSkew := clusterSemVer.Major - candidateWorkerSemver.Major
	if (candidateWorkerSemver.LessThan(clusterSemVer) || candidateWorkerSemver.Equal(clusterSemVer)) &&
		minorVersionSkew <= 2 &&
		majorVersionSkew == 0 {
		return true
	}

	return false
}

// Given a set of versions, filter down to only one version per minor version, taking the latest
// patch version. So if you had versions like:
// 		v1.11.1, v1.11.9, v1.12.6, v1.12.7
// it would only return:
//		v1.11.9, v1.12.7
func getFilteredPatchVersions(versions []string) ([]string, error) {
	filterPatchVersionsMap := make(map[string]string)
	for _, version := range versions {
		candidateVersion, err := semver.NewVersion(strings.TrimLeft(version, "v"))
		if err != nil {
			return nil, errors.Wrapf(err, "Unable to parse candidate version %s as semantic version", version)
		}
		minorVersionKey := fmt.Sprintf("%d.%d", candidateVersion.Major, candidateVersion.Minor)
		if val, isInMap := filterPatchVersionsMap[minorVersionKey]; isInMap {
			existingVersion, err := semver.NewVersion(strings.TrimLeft(val, "v"))
			if err != nil {
				return nil, errors.Wrapf(err, "Unable to parse existing version %s as semantic version", version)
			} else if existingVersion.LessThan(*candidateVersion) {
				filterPatchVersionsMap[minorVersionKey] = version
			}
		} else {
			filterPatchVersionsMap[minorVersionKey] = version
		}
	}
	filterPatchVersions := []string{}
	for _, version := range filterPatchVersionsMap {
		filterPatchVersions = append(filterPatchVersions, version)
	}

	return filterPatchVersions, nil
}

func ValidateMasterK8(ctx context.Context, tenancyID string, versionsGetter KubernetesGetter, v string) error {
	versions, err := GetSupportedK8s(ctx, tenancyID, versionsGetter)
	if err != nil {
		return errors.Wrap(err, "unable to validate master kubernetes version")
	}
	if err := validateK8(v, versions.MastersK8sVersions); err != nil {
		return errors.Wrap(err, "unsupported master kubernetes version")
	}
	return nil
}

func ValidateWorkerK8(ctx context.Context, tenancyID string, versionsGetter KubernetesGetter, pool, v string) error {
	versions, err := GetSupportedK8s(ctx, tenancyID, versionsGetter)
	if err != nil {
		return errors.Wrap(err, "unable to validate worker kubernetes version")
	}
	if err := validateK8(v, versions.NodesK8sVersions); err != nil {
		return errors.Wrapf(err, "unsupported kubernetes version on pool '%s'", pool)
	}
	return nil
}

func validateK8(v string, versions []string) error {
	for _, version := range versions {
		if v == version {
			return nil
		}
	}
	return fmt.Errorf("Unknown version, %s not one of %s", v, strings.Join(versions, ", "))
}

// HardcodedKubernetesGetter gets passed in a set of kubernetes versions and always returns those versions
type HardcodedKubernetesGetter struct {
	supportedVersions []string
	allVersions       []string
	etcdK8sMap        map[string]string
	KubernetesGetter
}

func NewHardcodedKubernetesGetter(supportedVersions []string, allVersions []string, etcdK8sMap map[string]string) *HardcodedKubernetesGetter {
	return &HardcodedKubernetesGetter{supportedVersions: supportedVersions, allVersions: allVersions, etcdK8sMap: etcdK8sMap}
}

func (hk *HardcodedKubernetesGetter) GetSupportedMasterKubernetes(ctx context.Context, tenancyID string) ([]string, error) {
	return hk.supportedVersions, nil
}

func (hk *HardcodedKubernetesGetter) GetSupportedWorkerKubernetes(ctx context.Context, tenancyID string) ([]string, error) {
	// Currently, this function returns the same as the master versions
	return hk.GetSupportedMasterKubernetes(ctx, tenancyID)
}

func (hk *HardcodedKubernetesGetter) GetAllKubernetes(ctx context.Context, tenancyID string) ([]string, error) {
	// Currently, this function returns the same as the master versions
	return hk.allVersions, nil
}

func (hk *HardcodedKubernetesGetter) GetKubernetesEtcdVersionMap(ctx context.Context, tenancyID string) (map[string]string, error) {
	return hk.etcdK8sMap, nil
}
