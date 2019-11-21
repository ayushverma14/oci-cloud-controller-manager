// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package datascience

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// UpdateNotebookSessionRequest wrapper for the UpdateNotebookSession operation
type UpdateNotebookSessionRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/identifiers.htm) of the notebook session.
	NotebookSessionId *string `mandatory:"true" contributesTo:"path" name:"notebookSessionId"`

	// Details for updating a notebook session. `notebookSessionConfigurationDetails` can only be updated while the notebook session is in the `INACTIVE` state.
	// Changes to the `notebookSessionConfigurationDetails` will take effect the next time the `ActivateNotebookSession` action is invoked on the notebook session resource.
	UpdateNotebookSessionDetails `contributesTo:"body"`

	// For optimistic concurrency control. In the PUT or DELETE call
	// for a resource, set the `if-match` parameter to the value of the
	// etag from a previous GET or POST response for that resource.
	// The resource will be updated or deleted only if the `etag` you
	// provide matches the resource's current `etag` value.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request UpdateNotebookSessionRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request UpdateNotebookSessionRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request UpdateNotebookSessionRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// UpdateNotebookSessionResponse wrapper for the UpdateNotebookSession operation
type UpdateNotebookSessionResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The NotebookSession instance
	NotebookSession `presentIn:"body"`

	// For optimistic concurrency control. See ETags for Optimistic Concurrency Control (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#eleven).
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response UpdateNotebookSessionResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response UpdateNotebookSessionResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
