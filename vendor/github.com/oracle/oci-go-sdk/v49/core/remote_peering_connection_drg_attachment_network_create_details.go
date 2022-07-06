// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// API covering the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services. Use this API
// to manage resources such as virtual cloud networks (VCNs), compute instances, and
// block storage volumes.
//

package core

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v49/common"
)

// RemotePeeringConnectionDrgAttachmentNetworkCreateDetails The representation of RemotePeeringConnectionDrgAttachmentNetworkCreateDetails
type RemotePeeringConnectionDrgAttachmentNetworkCreateDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the network attached to the DRG.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of compartment that contains the remote peering connection.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The BGP ASN to use for the remote peering connection's route target.
	RegionalOciAsn *string `mandatory:"true" json:"regionalOciAsn"`

	// Whether the RPC attachment is for a GFC DRG, indicating the mpls label should be
	// allocated from the VC label range.
	// Example: `true`
	IsGlobalFastConnect *bool `mandatory:"false" json:"isGlobalFastConnect"`
}

//GetId returns Id
func (m RemotePeeringConnectionDrgAttachmentNetworkCreateDetails) GetId() *string {
	return m.Id
}

func (m RemotePeeringConnectionDrgAttachmentNetworkCreateDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m RemotePeeringConnectionDrgAttachmentNetworkCreateDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeRemotePeeringConnectionDrgAttachmentNetworkCreateDetails RemotePeeringConnectionDrgAttachmentNetworkCreateDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeRemotePeeringConnectionDrgAttachmentNetworkCreateDetails
	}{
		"REMOTE_PEERING_CONNECTION",
		(MarshalTypeRemotePeeringConnectionDrgAttachmentNetworkCreateDetails)(m),
	}

	return json.Marshal(&s)
}
