package framework

import (
	"context"
	"time"

	. "github.com/onsi/gomega"

	"github.com/oracle/oci-go-sdk/v49/core"
)

// GetInstance return the specified instance from OCI compute.
func (f *Framework) GetInstance(instanceID string) core.Instance {
	ctx, cancel := context.WithTimeout(f.context, f.timeout)
	defer cancel()
	response, err := f.computeClient.GetInstance(ctx, core.GetInstanceRequest{
		InstanceId: &instanceID,
	})
	Expect(err).NotTo(HaveOccurred())
	return response.Instance
}

// DeleteInstance terminates the specified instance from OCI compute.
func (f *Framework) DeleteInstance(instanceID string, waitForDeleted bool) {
	Logf("Deleting instance, id: %s", instanceID)
	// Fetch the current OKE nodepool objects.
	instance := f.GetInstance(instanceID)
	// If not already 'TERMINATING' or 'TERMINATED' then issue terminate request.
	if instance.LifecycleState != core.InstanceLifecycleStateTerminating &&
		instance.LifecycleState != core.InstanceLifecycleStateTerminated {
		ctx, cancel := context.WithTimeout(f.context, f.timeout)
		defer cancel()
		_, err := f.computeClient.TerminateInstance(ctx, core.TerminateInstanceRequest{
			InstanceId: &instanceID,
		})
		Expect(err).NotTo(HaveOccurred())
	}
	// If specified wait for 'TERMINATED'.
	if waitForDeleted {
		timeout := 20 * time.Minute
		for start := time.Now(); time.Since(start) < timeout; time.Sleep(Poll) {
			instance = f.GetInstance(instanceID)
			Logf("Waiting for instance '%s' '%s' status - InstanceLifeCycleState: '%s'",
				*instance.DisplayName, core.InstanceLifecycleStateTerminated, instance.LifecycleState)
			if instance.LifecycleState == core.InstanceLifecycleStateTerminated {
				return
			}
		}
	}
}
