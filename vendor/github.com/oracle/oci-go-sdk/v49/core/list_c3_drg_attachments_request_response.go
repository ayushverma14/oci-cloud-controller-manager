// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package core

import (
	"github.com/oracle/oci-go-sdk/v49/common"
	"net/http"
)

// ListC3DrgAttachmentsRequest wrapper for the ListC3DrgAttachments operation
type ListC3DrgAttachmentsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VCN.
	VcnId *string `mandatory:"false" contributesTo:"query" name:"vcnId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DRG.
	DrgId *string `mandatory:"false" contributesTo:"query" name:"drgId"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List"
	// call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resource (virtual circuit, VCN, IPSec tunnel, or remote peering connection) attached to the DRG.
	NetworkId *string `mandatory:"false" contributesTo:"query" name:"networkId"`

	// The type for the network resource attached to the DRG.
	AttachmentType ListC3DrgAttachmentsAttachmentTypeEnum `mandatory:"false" contributesTo:"query" name:"attachmentType" omitEmpty:"true"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DRG route table assigned to the DRG attachment.
	DrgRouteTableId *string `mandatory:"false" contributesTo:"query" name:"drgRouteTableId"`

	// A filter to return only resources that match the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The field to sort by. You can provide one sort order (`sortOrder`). Default order for
	// TIMECREATED is descending. Default order for DISPLAYNAME is ascending. The DISPLAYNAME
	// sort order is case sensitive.
	// **Note:** In general, some "List" operations (for example, `ListInstances`) let you
	// optionally filter by availability domain if the scope of the resource type is within a
	// single availability domain. If you call one of these "List" operations without specifying
	// an availability domain, the resources are grouped by availability domain, then sorted.
	SortBy ListC3DrgAttachmentsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`). The DISPLAYNAME sort order
	// is case sensitive.
	SortOrder ListC3DrgAttachmentsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to return only resources that match the specified lifecycle
	// state. The value is case insensitive.
	LifecycleState DrgAttachmentLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListC3DrgAttachmentsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListC3DrgAttachmentsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListC3DrgAttachmentsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListC3DrgAttachmentsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListC3DrgAttachmentsResponse wrapper for the ListC3DrgAttachments operation
type ListC3DrgAttachmentsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []DrgAttachment instances
	Items []DrgAttachment `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListC3DrgAttachmentsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListC3DrgAttachmentsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListC3DrgAttachmentsAttachmentTypeEnum Enum with underlying type: string
type ListC3DrgAttachmentsAttachmentTypeEnum string

// Set of constants representing the allowable values for ListC3DrgAttachmentsAttachmentTypeEnum
const (
	ListC3DrgAttachmentsAttachmentTypeVcn                     ListC3DrgAttachmentsAttachmentTypeEnum = "VCN"
	ListC3DrgAttachmentsAttachmentTypeVirtualCircuit          ListC3DrgAttachmentsAttachmentTypeEnum = "VIRTUAL_CIRCUIT"
	ListC3DrgAttachmentsAttachmentTypeRemotePeeringConnection ListC3DrgAttachmentsAttachmentTypeEnum = "REMOTE_PEERING_CONNECTION"
	ListC3DrgAttachmentsAttachmentTypeIpsecTunnel             ListC3DrgAttachmentsAttachmentTypeEnum = "IPSEC_TUNNEL"
	ListC3DrgAttachmentsAttachmentTypeAll                     ListC3DrgAttachmentsAttachmentTypeEnum = "ALL"
)

var mappingListC3DrgAttachmentsAttachmentType = map[string]ListC3DrgAttachmentsAttachmentTypeEnum{
	"VCN":                       ListC3DrgAttachmentsAttachmentTypeVcn,
	"VIRTUAL_CIRCUIT":           ListC3DrgAttachmentsAttachmentTypeVirtualCircuit,
	"REMOTE_PEERING_CONNECTION": ListC3DrgAttachmentsAttachmentTypeRemotePeeringConnection,
	"IPSEC_TUNNEL":              ListC3DrgAttachmentsAttachmentTypeIpsecTunnel,
	"ALL":                       ListC3DrgAttachmentsAttachmentTypeAll,
}

// GetListC3DrgAttachmentsAttachmentTypeEnumValues Enumerates the set of values for ListC3DrgAttachmentsAttachmentTypeEnum
func GetListC3DrgAttachmentsAttachmentTypeEnumValues() []ListC3DrgAttachmentsAttachmentTypeEnum {
	values := make([]ListC3DrgAttachmentsAttachmentTypeEnum, 0)
	for _, v := range mappingListC3DrgAttachmentsAttachmentType {
		values = append(values, v)
	}
	return values
}

// ListC3DrgAttachmentsSortByEnum Enum with underlying type: string
type ListC3DrgAttachmentsSortByEnum string

// Set of constants representing the allowable values for ListC3DrgAttachmentsSortByEnum
const (
	ListC3DrgAttachmentsSortByTimecreated ListC3DrgAttachmentsSortByEnum = "TIMECREATED"
	ListC3DrgAttachmentsSortByDisplayname ListC3DrgAttachmentsSortByEnum = "DISPLAYNAME"
)

var mappingListC3DrgAttachmentsSortBy = map[string]ListC3DrgAttachmentsSortByEnum{
	"TIMECREATED": ListC3DrgAttachmentsSortByTimecreated,
	"DISPLAYNAME": ListC3DrgAttachmentsSortByDisplayname,
}

// GetListC3DrgAttachmentsSortByEnumValues Enumerates the set of values for ListC3DrgAttachmentsSortByEnum
func GetListC3DrgAttachmentsSortByEnumValues() []ListC3DrgAttachmentsSortByEnum {
	values := make([]ListC3DrgAttachmentsSortByEnum, 0)
	for _, v := range mappingListC3DrgAttachmentsSortBy {
		values = append(values, v)
	}
	return values
}

// ListC3DrgAttachmentsSortOrderEnum Enum with underlying type: string
type ListC3DrgAttachmentsSortOrderEnum string

// Set of constants representing the allowable values for ListC3DrgAttachmentsSortOrderEnum
const (
	ListC3DrgAttachmentsSortOrderAsc  ListC3DrgAttachmentsSortOrderEnum = "ASC"
	ListC3DrgAttachmentsSortOrderDesc ListC3DrgAttachmentsSortOrderEnum = "DESC"
)

var mappingListC3DrgAttachmentsSortOrder = map[string]ListC3DrgAttachmentsSortOrderEnum{
	"ASC":  ListC3DrgAttachmentsSortOrderAsc,
	"DESC": ListC3DrgAttachmentsSortOrderDesc,
}

// GetListC3DrgAttachmentsSortOrderEnumValues Enumerates the set of values for ListC3DrgAttachmentsSortOrderEnum
func GetListC3DrgAttachmentsSortOrderEnumValues() []ListC3DrgAttachmentsSortOrderEnum {
	values := make([]ListC3DrgAttachmentsSortOrderEnum, 0)
	for _, v := range mappingListC3DrgAttachmentsSortOrder {
		values = append(values, v)
	}
	return values
}
