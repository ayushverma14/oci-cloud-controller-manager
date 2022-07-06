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
	"github.com/oracle/oci-go-sdk/v49/common"
)

// RemoveAdditionalRouteRulesDetails The configuration details for the remove rules operation. Only one of the three properties should be
// supplied to call the removeAdditionalRouteRules API.
type RemoveAdditionalRouteRulesDetails struct {

	// The list of route rule identifiers used for removing route rules from route table.
	AdditionalRouteRuleIds []string `mandatory:"false" json:"additionalRouteRuleIds"`

	// The list of destinations used for removing route rules from route table. This is only supplied
	// when additionalRouteRuleIds are not provided.
	Destinations []string `mandatory:"false" json:"destinations"`

	// The pairs of <destination, target> used for removing route rules from route table. This is only
	// supplied when additionalRouteRules and destinations are not provided.
	RouteDestinationRouteTargets []RouteDestinationRouteTargetDetails `mandatory:"false" json:"routeDestinationRouteTargets"`
}

func (m RemoveAdditionalRouteRulesDetails) String() string {
	return common.PointerString(m)
}
