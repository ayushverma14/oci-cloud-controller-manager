// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Load Balancing API
//
// API for the Load Balancing service. Use this API to manage load balancers, backend sets, and related items. For more
// information, see Overview of Load Balancing (https://docs.cloud.oracle.com/iaas/Content/Balance/Concepts/balanceoverview.htm).
//

package loadbalancer

import (
	"github.com/oracle/oci-go-sdk/v49/common"
)

// ConnectionConfiguration Configuration details for the connection between the client and backend servers.
type ConnectionConfiguration struct {

	// The maximum idle time, in seconds, allowed between two successive receive or two successive send operations
	// between the client and backend servers. A send operation does not reset the timer for receive operations. A
	// receive operation does not reset the timer for send operations.
	// For more information, see Connection Configuration (https://docs.cloud.oracle.com/Content/Balance/Reference/connectionreuse.htm#ConnectionConfiguration).
	// Example: `1200`
	IdleTimeout *int64 `mandatory:"true" json:"idleTimeout"`

	// The backend TCP Proxy Protocol version.
	// Example: `1`
	BackendTcpProxyProtocolVersion *int `mandatory:"false" json:"backendTcpProxyProtocolVersion"`

	// An array that represents the PPV2 Options that can be enabled on TCP Listeners.
	// Example: ["PP2_TYPE_AUTHORITY"]
	BackendTcpProxyProtocolOptions []ConnectionConfigurationBackendTcpProxyProtocolOptionsEnum `mandatory:"false" json:"backendTcpProxyProtocolOptions,omitempty"`
}

func (m ConnectionConfiguration) String() string {
	return common.PointerString(m)
}

// ConnectionConfigurationBackendTcpProxyProtocolOptionsEnum Enum with underlying type: string
type ConnectionConfigurationBackendTcpProxyProtocolOptionsEnum string

// Set of constants representing the allowable values for ConnectionConfigurationBackendTcpProxyProtocolOptionsEnum
const (
	ConnectionConfigurationBackendTcpProxyProtocolOptionsPp2TypeAuthority ConnectionConfigurationBackendTcpProxyProtocolOptionsEnum = "PP2_TYPE_AUTHORITY"
)

var mappingConnectionConfigurationBackendTcpProxyProtocolOptions = map[string]ConnectionConfigurationBackendTcpProxyProtocolOptionsEnum{
	"PP2_TYPE_AUTHORITY": ConnectionConfigurationBackendTcpProxyProtocolOptionsPp2TypeAuthority,
}

// GetConnectionConfigurationBackendTcpProxyProtocolOptionsEnumValues Enumerates the set of values for ConnectionConfigurationBackendTcpProxyProtocolOptionsEnum
func GetConnectionConfigurationBackendTcpProxyProtocolOptionsEnumValues() []ConnectionConfigurationBackendTcpProxyProtocolOptionsEnum {
	values := make([]ConnectionConfigurationBackendTcpProxyProtocolOptionsEnum, 0)
	for _, v := range mappingConnectionConfigurationBackendTcpProxyProtocolOptions {
		values = append(values, v)
	}
	return values
}
