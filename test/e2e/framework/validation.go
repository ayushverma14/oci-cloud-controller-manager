package framework

import (
	"fmt"
	"strings"

	. "github.com/onsi/gomega"

	oke "github.com/oracle/oci-go-sdk/v49/containerengine"
	"github.com/oracle/oci-go-sdk/v49/core"
)

// CrossValidateCluster checks a Cluster is consistent internally, and
// consistent with the associated ClusterSummary, and, that the underlying
// nodepools are consistent.
func (f *Framework) CrossValidateCluster(clusterID string, validateChildResources bool) {
	Logf("cross validating cluster: %q", clusterID)
	cluster := f.GetCluster(clusterID)
	clusterSummary := f.GetClusterSummary(clusterID)
	Expect(clusterSummary).ToNot(BeNil())
	// Cross-validate the cluster with the cluster summary.
	Expect(cluster.Id).To(Equal(clusterSummary.Id))
	Expect(cluster.Name).To(Equal(clusterSummary.Name))
	Expect(cluster.CompartmentId).To(Equal(clusterSummary.CompartmentId))
	Expect(cluster.Endpoints.Kubernetes).To(Equal(clusterSummary.Endpoints.Kubernetes))
	Expect(fmt.Sprintf("%s", cluster.LifecycleState)).To(Equal(fmt.Sprintf("%s", clusterSummary.LifecycleState)))
	// Cross-validate the clusters node.
	if validateChildResources {
		for _, nodepool := range f.ListNodePools(clusterID) {
			//nodepool k8s version may be different from cluster k8s version
			//f.CrossValidateNodePool(*nodepool.Id, *cluster.KubernetesVersion)
			f.CrossValidateNodePool(*nodepool.Id, f.OkeNodePoolK8sVersion)
		}
	}

}

// CrossValidateNodePool checks a NodePool is consistent internally, consistent
// with the associated NodePoolSummary, and, consistent with corresponding
// OCI instances.
// NB: This method should only be used to check nodepools and nodes that
// have an 'ACTIVE" state.
func (f *Framework) CrossValidateNodePool(nodepoolID, kubernetesVersion string) {
	Logf("cross validating nodepool: %q", nodepoolID)
	nodepool := f.GetNodePool(nodepoolID)
	nodepoolSummary := f.GetNodePoolSummary(*nodepool.ClusterId, nodepoolID)
	Expect(nodepoolSummary).ToNot(BeNil())

	// The number of nodes should meet the specified invariant.
	var expectedNumNodes int
	if IsVersion2NodePool(nodepool) {
		expectedNumNodes = *nodepool.NodeConfigDetails.Size
	} else {
		expectedNumNodes = *nodepool.QuantityPerSubnet * len(nodepool.SubnetIds)
	}

	// Each (ACTIVE) node should be unique.
	uniqueNodesIds := make(map[string]oke.NodeLifecycleStateEnum)

	for _, node := range nodepool.Nodes {
		Expect(*nodepool.Id).To(Equal(*node.NodePoolId))
		if node.LifecycleState == oke.NodeLifecycleStateActive {
			uniqueNodesIds[*node.Id] = node.LifecycleState
		}
	}
	Expect(expectedNumNodes).To(Equal(len(uniqueNodesIds)))

	// Each subnet should have the correct number of nodes if using
	// QuantityPerSubnet model
	if !IsVersion2NodePool(nodepool) {
		nodesPerSubnet := make(map[string]int)

		for _, node := range nodepool.Nodes {
			Expect(*nodepool.Id).To(Equal(*node.NodePoolId))
			if node.LifecycleState == oke.NodeLifecycleStateActive {
				nodesPerSubnet[*node.SubnetId]++
			}
		}
		for _, numNodes := range nodesPerSubnet {
			Expect(numNodes).To(Equal(*nodepool.QuantityPerSubnet))
		}
	}

	// Check that nodepool record matches its summary.
	Expect(*nodepool.ClusterId).To(Equal(*nodepoolSummary.ClusterId))
	Expect(*nodepool.Id).To(Equal(*nodepoolSummary.Id))
	Expect(*nodepool.Name).To(Equal(*nodepoolSummary.Name))
	for idx, nodeLabel := range nodepool.InitialNodeLabels {
		// NB: Assumes key-value pairs maintain order.
		nodepoolSummaryLabel := nodepoolSummary.InitialNodeLabels[idx]
		Expect(nodeLabel.Key).To(Equal(nodepoolSummaryLabel.Key))
		Expect(nodeLabel.Value).To(Equal(nodepoolSummaryLabel.Value))
	}

	if IsVersion2NodePool(nodepool) {
		Expect(*nodepool.NodeConfigDetails.Size).To(Equal(*nodepoolSummary.NodeConfigDetails.Size))
		foundMatch := false
		for _, nodePlacement := range nodepoolSummary.NodeConfigDetails.PlacementConfigs {
			for _, val := range nodepool.NodeConfigDetails.PlacementConfigs {
				if strings.Compare(*val.AvailabilityDomain, *nodePlacement.AvailabilityDomain) == 0 &&
					strings.Compare(*val.SubnetId, *nodePlacement.SubnetId) == 0 {
					foundMatch = true
					break
				}
			}

			if !foundMatch {
				Failf("Failed to find matching node placement %s,%s", *nodePlacement.AvailabilityDomain, *nodePlacement.SubnetId)
			}

			foundMatch = false
		}
	} else {
		Expect(*nodepool.QuantityPerSubnet).To(Equal(*nodepoolSummary.QuantityPerSubnet))
		for idx, subnet := range nodepoolSummary.SubnetIds {
			// NB: Assumes subnets maintain order.
			Expect(subnet).To(Equal(nodepoolSummary.SubnetIds[idx]))
		}
	}

	Expect(kubernetesVersion).To(Equal(*nodepoolSummary.KubernetesVersion))
	// TODO
	// Expect(*nodepoolSummary.SshPublicKey).To(Equal(???))

	// Check the OCI node listed should be active and match the node specification.
	for _, node := range nodepool.Nodes {
		Expect(*nodepool.Id).To(Equal(*node.NodePoolId))
		if node.LifecycleState == oke.NodeLifecycleStateActive {
			instance := f.GetInstance(*node.Id)
			Expect(instance.LifecycleState).To(Equal(core.InstanceLifecycleStateRunning))
			Expect(*instance.Id).To(Equal(*node.Id))
			Expect(*instance.DisplayName).To(Equal(*node.Name))
			// TODO: Fetch VnicAttachment/Vnic fand Subnets.
			// Expect(*instance.XXX).To(Equal(*node.PublicIp))
			// Expect(*instance.XXX).To(Equal(*node.SubnetId))
			Expect(*instance.CompartmentId).To(Equal(*nodepoolSummary.CompartmentId))
			Expect(*instance.ImageId).To(Equal(*nodepoolSummary.NodeImageId))
			// TODO: Fetch image name
			// Expect(*instance.XXX).To(Equal(*nodepoolSummary.NodeImageName))
			Expect(*instance.Shape).To(Equal(*nodepoolSummary.NodeShape))
		}
	}
}

// findNodesInState gets the nodes in the specified state.
func findNodesInState(nodepool *oke.NodePool, state oke.NodeLifecycleStateEnum) []oke.Node {
	nodesInState := make([]oke.Node, 0)
	for _, node := range nodepool.Nodes {
		if node.LifecycleState == state {
			nodesInState = append(nodesInState, node)
		}
	}
	return nodesInState
}

// findNodesWithErrors gets the nodes in the specified state with error messages.
func findNodesWithErrors(nodepool *oke.NodePool, state oke.NodeLifecycleStateEnum) []oke.Node {
	nodesWithErrors := make([]oke.Node, 0)
	for _, node := range nodepool.Nodes {
		if node.LifecycleState == state {
			nodesWithErrors = append(nodesWithErrors, node)
		}
	}
	return nodesWithErrors
}

// countNodesInState return the number of nodes in the specified state.
func countNodesInState(nodepool oke.NodePool, state oke.NodeLifecycleStateEnum) int {
	numInState := 0
	for _, node := range nodepool.Nodes {
		if node.LifecycleState == state {
			numInState++
		}
	}
	return numInState
}

// hasExpectedNodesInState returns true if the nodes in the nodepool have the
// expected number of nodes in the specified state.
func hasExpectedNodesInState(nodepool oke.NodePool, state oke.NodeLifecycleStateEnum) bool {
	var expectedInState int
	if IsVersion2NodePool(&nodepool) {
		expectedInState = *nodepool.NodeConfigDetails.Size
	} else {
		expectedInState = *nodepool.QuantityPerSubnet * len(nodepool.SubnetIds)
	}
	numInState := countNodesInState(nodepool, state)
	return expectedInState == numInState
}

// hasInstance returns true if the specified instance exists in the nodepool.
func hasInstance(nodepool oke.NodePool, instanceID string) bool {
	for _, node := range nodepool.Nodes {
		if *node.Id == instanceID {
			Logf("Deleting nodepool '%s' instance: '%s' .", *nodepool.Name, instanceID)
			return true
		}
	}
	Logf("Deleted nodepool '%s' instance: '%s' .", *nodepool.Name, instanceID)
	return false
}

// json raw data must start with object "{}" or array "[]"
func (f *Framework) IsNotJsonFormatStr(data string) bool {
	trimmedData := strings.TrimSpace(data)
	return !(strings.HasPrefix(trimmedData, "{") && strings.HasSuffix(trimmedData, "}")) &&
		!(strings.HasPrefix(trimmedData, "[") && strings.HasSuffix(trimmedData, "]"))
}
