// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Health Checks API
//
// API for the Health Checks service. Use this API to manage endpoint probes and monitors.
// For more information, see
// Overview of the Health Checks Service (https://docs.cloud.oracle.com/iaas/Content/HealthChecks/Concepts/healthchecks.htm).
//

package healthchecks

import (
	"github.com/oracle/oci-go-sdk/common"
)

// UpdateHttpMonitorDetails The request body used to update an HTTP monitor.
type UpdateHttpMonitorDetails struct {

	// A list of targets (hostnames or IP addresses) of the probe.
	Targets []string `mandatory:"false" json:"targets"`

	// A list of names of vantage points from which to execute the probe.
	VantagePointNames []string `mandatory:"false" json:"vantagePointNames"`

	// The port on which to probe endpoints. If unspecified, probes will use the
	// default port of their protocol.
	Port *int `mandatory:"false" json:"port"`

	// The probe timeout in seconds. Valid values: 10, 20, 30, and 60.
	// The probe timeout must be less than or equal to `intervalInSeconds` for monitors.
	TimeoutInSeconds *int `mandatory:"false" json:"timeoutInSeconds"`

	Protocol HttpProbeProtocolEnum `mandatory:"false" json:"protocol,omitempty"`

	Method HttpProbeMethodEnum `mandatory:"false" json:"method,omitempty"`

	// The optional URL path to probe, including query parameters.
	Path *string `mandatory:"false" json:"path"`

	// A dictionary of HTTP request headers.
	// *Note:* Monitors and probes do not support the use of the `Authorization` HTTP header.
	Headers map[string]string `mandatory:"false" json:"headers"`

	// A user-friendly and mutable name suitable for display in a user interface.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The monitor interval in seconds. Valid values: 10, 30, and 60.
	IntervalInSeconds *int `mandatory:"false" json:"intervalInSeconds"`

	// Enables or disables the monitor. Set to 'true' to launch monitoring.
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace.  For more information,
	// see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateHttpMonitorDetails) String() string {
	return common.PointerString(m)
}

// UpdateHttpMonitorDetailsProtocolEnum is an alias to type: HttpProbeProtocolEnum
// Consider using HttpProbeProtocolEnum instead
// Deprecated
type UpdateHttpMonitorDetailsProtocolEnum = HttpProbeProtocolEnum

// Set of constants representing the allowable values for HttpProbeProtocolEnum
// Deprecated
const (
	UpdateHttpMonitorDetailsProtocolHttp  HttpProbeProtocolEnum = "HTTP"
	UpdateHttpMonitorDetailsProtocolHttps HttpProbeProtocolEnum = "HTTPS"
)

// GetUpdateHttpMonitorDetailsProtocolEnumValues Enumerates the set of values for HttpProbeProtocolEnum
// Consider using GetHttpProbeProtocolEnumValue
// Deprecated
var GetUpdateHttpMonitorDetailsProtocolEnumValues = GetHttpProbeProtocolEnumValues

// UpdateHttpMonitorDetailsMethodEnum is an alias to type: HttpProbeMethodEnum
// Consider using HttpProbeMethodEnum instead
// Deprecated
type UpdateHttpMonitorDetailsMethodEnum = HttpProbeMethodEnum

// Set of constants representing the allowable values for HttpProbeMethodEnum
// Deprecated
const (
	UpdateHttpMonitorDetailsMethodGet  HttpProbeMethodEnum = "GET"
	UpdateHttpMonitorDetailsMethodHead HttpProbeMethodEnum = "HEAD"
)

// GetUpdateHttpMonitorDetailsMethodEnumValues Enumerates the set of values for HttpProbeMethodEnum
// Consider using GetHttpProbeMethodEnumValue
// Deprecated
var GetUpdateHttpMonitorDetailsMethodEnumValues = GetHttpProbeMethodEnumValues
