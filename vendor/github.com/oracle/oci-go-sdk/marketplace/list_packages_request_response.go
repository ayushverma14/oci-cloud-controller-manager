// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package marketplace

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListPackagesRequest wrapper for the ListPackages operation
type ListPackagesRequest struct {

	// The unique identifier of the listing.
	ListingId *string `mandatory:"true" contributesTo:"path" name:"listingId"`

	// The version of the package. Package versions are unique within a listing.
	PackageVersion *string `mandatory:"false" contributesTo:"query" name:"packageVersion"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// How many records to return. Specify a value greater than zero and less than or equal to 1000. The default is 30.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field that is used to sort listed results. You can only specify one field to sort by.
	// `TIMERELEASED` displays results in descending order by default. `NAME` displays results in
	// ascending order by default. You can change your preference by specifying a different sort order.
	SortBy ListPackagesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListPackagesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListPackagesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListPackagesRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListPackagesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListPackagesResponse wrapper for the ListPackages operation
type ListPackagesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ListingPackageSummary instances
	Items []ListingPackageSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain.
	// Include this value as the `page` parameter for the subsequent GET request. For important details about
	// how pagination works, see List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#List_Pagination).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListPackagesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListPackagesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListPackagesSortByEnum Enum with underlying type: string
type ListPackagesSortByEnum string

// Set of constants representing the allowable values for ListPackagesSortByEnum
const (
	ListPackagesSortByName         ListPackagesSortByEnum = "NAME"
	ListPackagesSortByTimereleased ListPackagesSortByEnum = "TIMERELEASED"
)

var mappingListPackagesSortBy = map[string]ListPackagesSortByEnum{
	"NAME":         ListPackagesSortByName,
	"TIMERELEASED": ListPackagesSortByTimereleased,
}

// GetListPackagesSortByEnumValues Enumerates the set of values for ListPackagesSortByEnum
func GetListPackagesSortByEnumValues() []ListPackagesSortByEnum {
	values := make([]ListPackagesSortByEnum, 0)
	for _, v := range mappingListPackagesSortBy {
		values = append(values, v)
	}
	return values
}

// ListPackagesSortOrderEnum Enum with underlying type: string
type ListPackagesSortOrderEnum string

// Set of constants representing the allowable values for ListPackagesSortOrderEnum
const (
	ListPackagesSortOrderAsc  ListPackagesSortOrderEnum = "ASC"
	ListPackagesSortOrderDesc ListPackagesSortOrderEnum = "DESC"
)

var mappingListPackagesSortOrder = map[string]ListPackagesSortOrderEnum{
	"ASC":  ListPackagesSortOrderAsc,
	"DESC": ListPackagesSortOrderDesc,
}

// GetListPackagesSortOrderEnumValues Enumerates the set of values for ListPackagesSortOrderEnum
func GetListPackagesSortOrderEnumValues() []ListPackagesSortOrderEnum {
	values := make([]ListPackagesSortOrderEnum, 0)
	for _, v := range mappingListPackagesSortOrder {
		values = append(values, v)
	}
	return values
}