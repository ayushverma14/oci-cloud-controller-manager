// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package loadbalancer

import (
	"github.com/oracle/oci-go-sdk/v49/common"
	"net/http"
)

// GetRuleSetRequest wrapper for the GetRuleSet operation
type GetRuleSetRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the specified load balancer.
	LoadBalancerId *string `mandatory:"true" contributesTo:"path" name:"loadBalancerId"`

	// The name of the rule set to retrieve.
	// Example: `example_rule_set`
	RuleSetName *string `mandatory:"true" contributesTo:"path" name:"ruleSetName"`

	// The unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// For optimistic concurrency control. In the PUT or DELETE call for a resource, set the if-match
	// parameter to the value of the ETag from a previous GET or POST response for that resource. The
	// resource is updated or deleted only if the ETag you provide matches the resource's current ETag
	// value.
	// Example: `example-etag`
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// The system returns the requested resource, with a 200 status, only if the resource has no ETag
	// matching the one specified. If the condition fails for the GET and HEAD methods, the system returns the
	// HTTP status code `304 (Not Modified)`.
	// Example: `example-etag`
	IfNoneMatch *string `mandatory:"false" contributesTo:"header" name:"if-none-match"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetRuleSetRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetRuleSetRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetRuleSetRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetRuleSetRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetRuleSetResponse wrapper for the GetRuleSet operation
type GetRuleSetResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The RuleSet instance
	RuleSet `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetRuleSetResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetRuleSetResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
