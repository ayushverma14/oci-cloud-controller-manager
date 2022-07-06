#!/bin/bash

##################################################################################################
# This template can be used to tweak the environment variables needed to run the E2E tests locally #
# Default behavior:
# Runs test on an existing cluster in dev0-iad

# To run the tests:
# 1. Change the FOCUS valiable here to specify the subset of E2E tests to run
# 2. Set CLUSTER_KUBECONFIG and CLOUD_CONFIG if needed
# 3. run 'source existing-cluster-dev0-env-template.sh' to set the variables
# 4. run 'make run-ccm-e2e-tests-local`
##################################################################################################

# The test suites to run (can replace or add tags)
export FOCUS="\[cloudprovider\]"

# This variable tells the test not to install oci cli and wipe out your .oci/config
export LOCAL_RUN=1
export TC_BUILD=0

# This allows you to use your existing cluster
export ENABLE_CREATE_CLUSTER=false

# Set path to kubeconfig of existing cluster if it does not exist in default path. Defaults to $HOME/.kube/config_*
export CLUSTER_KUBECONFIG_AMD=$HOME/.kube/config_amd
export CLUSTER_KUBECONFIG_ARM=$HOME/.kube/config_arm

# Set path to cloud_config of existing cluster if it does not exist in default path. Defaults to $HOME/cloudconfig_*
export CLOUD_CONFIG_AMD=$HOME/cloudconfig_amd
export CLOUD_CONFIG_ARM=$HOME/cloudconfig_arm


export IMAGE_PULL_REPO="iad.ocir.io/okedev/e2e-tests/"
export ADLOCATION="IqDk:US-ASHBURN-AD-1"

#KMS key for CMEK testing
export CMEK_KMS_KEY="ocid1.key.oc1.iad.bbpvrcsaaaeuk.abuwcljsav7rilbt6bnu3dqoakpzdtxhfk27uixzdz3yk7jrwngptfwg5u5a"

#NSG Network security group created in cluster's VCN
export NSG_OCIDS=""

#Reserved IP created in e2e test compartment
export RESERVED_IP="144.25.98.32"

#Architecture to run tests on
export ARCHITECTURE_AMD="AMD"
export ARCHITECTURE_ARM="ARM"

#Focus the tests : ARM, AMD or BOTH
export SCOPE="BOTH"

# For debugging the tests in existing cluster, do not turn it off by default.
# export DELETE_NAMESPACE=false

# FSS volume handle
# format is FileSystemOCID:serverIP:path
export FSS_VOLUME_HANDLE="ocid1.filesystem.oc1.iad.aaaaaaaaaaa5wj2infqwillqojxwiotjmfsc2ylefuzqaaaa:10.0.10.104:/FileSystem-20210820-0454-50"
