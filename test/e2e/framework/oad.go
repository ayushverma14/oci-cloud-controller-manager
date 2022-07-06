package framework

import (
	"context"

	"github.com/oracle/oci-go-sdk/v49/identity"

	. "github.com/onsi/gomega"
)

func (f *Framework) IsOAD() bool {
	response, err := f.identityClient.ListAvailabilityDomains(context.Background(), identity.ListAvailabilityDomainsRequest{
		CompartmentId: &f.Compartment1,
	})
	Expect(err).NotTo(HaveOccurred())
	return len(response.Items) == 1
}

func (f *Framework) ListADs() []identity.AvailabilityDomain {
	response, err := f.identityClient.ListAvailabilityDomains(context.Background(), identity.ListAvailabilityDomainsRequest{
		CompartmentId: &f.Compartment1,
	})
	Expect(err).NotTo(HaveOccurred())
	return response.Items
}
