package policy

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator 0.17.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// AssignmentsClient is the to manage and control access to your resources,
// you can define customized policies and assign them at a scope.
type AssignmentsClient struct {
	ManagementClient
}

// NewAssignmentsClient creates an instance of the AssignmentsClient client.
func NewAssignmentsClient(subscriptionID string) AssignmentsClient {
	return NewAssignmentsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewAssignmentsClientWithBaseURI creates an instance of the
// AssignmentsClient client.
func NewAssignmentsClientWithBaseURI(baseURI string, subscriptionID string) AssignmentsClient {
	return AssignmentsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create policy assignments are inherited by child resources. For example,
// when you apply a policy to a resource group that policy is assigned to all
// resources in the group.
//
// scope is the scope of the policy assignment. policyAssignmentName is the
// name of the policy assignment. parameters is parameters for the policy
// assignment.
func (client AssignmentsClient) Create(scope string, policyAssignmentName string, parameters Assignment) (result Assignment, err error) {
	req, err := client.CreatePreparer(scope, policyAssignmentName, parameters)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "policy.AssignmentsClient", "Create", nil, "Failure preparing request")
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "policy.AssignmentsClient", "Create", resp, "Failure sending request")
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.AssignmentsClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client AssignmentsClient) CreatePreparer(scope string, policyAssignmentName string, parameters Assignment) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"policyAssignmentName": autorest.Encode("path", policyAssignmentName),
		"scope":                scope,
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{scope}/providers/Microsoft.Authorization/policyassignments/{policyAssignmentName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client AssignmentsClient) CreateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client AssignmentsClient) CreateResponder(resp *http.Response) (result Assignment, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// CreateByID policy assignments are inherited by child resources. For
// example, when you apply a policy to a resource group that policy is
// assigned to all resources in the group. When providing a scope for the
// assigment, use '/subscriptions/{subscription-id}/' for subscriptions,
// '/subscriptions/{subscription-id}/resourceGroups/{resource-group-name}'
// for resource groups, and
// '/subscriptions/{subscription-id}/resourceGroups/{resource-group-name}/providers/{resource-provider-namespace}/{resource-type}/{resource-name}'
// for resources.
//
// policyAssignmentID is the ID of the policy assignment to create. Use the
// format
// '/{scope}/providers/Microsoft.Authorization/policyAssignments/{policy-assignment-name}'.
// parameters is parameters for policy assignment.
func (client AssignmentsClient) CreateByID(policyAssignmentID string, parameters Assignment) (result Assignment, err error) {
	req, err := client.CreateByIDPreparer(policyAssignmentID, parameters)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "policy.AssignmentsClient", "CreateByID", nil, "Failure preparing request")
	}

	resp, err := client.CreateByIDSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "policy.AssignmentsClient", "CreateByID", resp, "Failure sending request")
	}

	result, err = client.CreateByIDResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.AssignmentsClient", "CreateByID", resp, "Failure responding to request")
	}

	return
}

// CreateByIDPreparer prepares the CreateByID request.
func (client AssignmentsClient) CreateByIDPreparer(policyAssignmentID string, parameters Assignment) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"policyAssignmentId": policyAssignmentID,
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{policyAssignmentId}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// CreateByIDSender sends the CreateByID request. The method will close the
// http.Response Body if it receives an error.
func (client AssignmentsClient) CreateByIDSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// CreateByIDResponder handles the response to the CreateByID request. The method always
// closes the http.Response Body.
func (client AssignmentsClient) CreateByIDResponder(resp *http.Response) (result Assignment, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a policy assignment.
//
// scope is the scope of the policy assignment. policyAssignmentName is the
// name of the policy assignment to delete.
func (client AssignmentsClient) Delete(scope string, policyAssignmentName string) (result Assignment, err error) {
	req, err := client.DeletePreparer(scope, policyAssignmentName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "policy.AssignmentsClient", "Delete", nil, "Failure preparing request")
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "policy.AssignmentsClient", "Delete", resp, "Failure sending request")
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.AssignmentsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client AssignmentsClient) DeletePreparer(scope string, policyAssignmentName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"policyAssignmentName": autorest.Encode("path", policyAssignmentName),
		"scope":                scope,
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{scope}/providers/Microsoft.Authorization/policyassignments/{policyAssignmentName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client AssignmentsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client AssignmentsClient) DeleteResponder(resp *http.Response) (result Assignment, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// DeleteByID when providing a scope for the assigment, use
// '/subscriptions/{subscription-id}/' for subscriptions,
// '/subscriptions/{subscription-id}/resourceGroups/{resource-group-name}'
// for resource groups, and
// '/subscriptions/{subscription-id}/resourceGroups/{resource-group-name}/providers/{resource-provider-namespace}/{resource-type}/{resource-name}'
// for resources.
//
// policyAssignmentID is the ID of the policy assignment to delete. Use the
// format
// '/{scope}/providers/Microsoft.Authorization/policyAssignments/{policy-assignment-name}'.
func (client AssignmentsClient) DeleteByID(policyAssignmentID string) (result Assignment, err error) {
	req, err := client.DeleteByIDPreparer(policyAssignmentID)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "policy.AssignmentsClient", "DeleteByID", nil, "Failure preparing request")
	}

	resp, err := client.DeleteByIDSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "policy.AssignmentsClient", "DeleteByID", resp, "Failure sending request")
	}

	result, err = client.DeleteByIDResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.AssignmentsClient", "DeleteByID", resp, "Failure responding to request")
	}

	return
}

// DeleteByIDPreparer prepares the DeleteByID request.
func (client AssignmentsClient) DeleteByIDPreparer(policyAssignmentID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"policyAssignmentId": policyAssignmentID,
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{policyAssignmentId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// DeleteByIDSender sends the DeleteByID request. The method will close the
// http.Response Body if it receives an error.
func (client AssignmentsClient) DeleteByIDSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// DeleteByIDResponder handles the response to the DeleteByID request. The method always
// closes the http.Response Body.
func (client AssignmentsClient) DeleteByIDResponder(resp *http.Response) (result Assignment, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get gets a policy assignment.
//
// scope is the scope of the policy assignment. policyAssignmentName is the
// name of the policy assignment to get.
func (client AssignmentsClient) Get(scope string, policyAssignmentName string) (result Assignment, err error) {
	req, err := client.GetPreparer(scope, policyAssignmentName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "policy.AssignmentsClient", "Get", nil, "Failure preparing request")
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "policy.AssignmentsClient", "Get", resp, "Failure sending request")
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.AssignmentsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client AssignmentsClient) GetPreparer(scope string, policyAssignmentName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"policyAssignmentName": autorest.Encode("path", policyAssignmentName),
		"scope":                scope,
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{scope}/providers/Microsoft.Authorization/policyassignments/{policyAssignmentName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client AssignmentsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client AssignmentsClient) GetResponder(resp *http.Response) (result Assignment, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetByID when providing a scope for the assigment, use
// '/subscriptions/{subscription-id}/' for subscriptions,
// '/subscriptions/{subscription-id}/resourceGroups/{resource-group-name}'
// for resource groups, and
// '/subscriptions/{subscription-id}/resourceGroups/{resource-group-name}/providers/{resource-provider-namespace}/{resource-type}/{resource-name}'
// for resources.
//
// policyAssignmentID is the ID of the policy assignment to get. Use the
// format '/{scope}/providers/Microsoft.Authorization/policyAssignments/{policy-assignment-name}'.
func (client AssignmentsClient) GetByID(policyAssignmentID string) (result Assignment, err error) {
	req, err := client.GetByIDPreparer(policyAssignmentID)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "policy.AssignmentsClient", "GetByID", nil, "Failure preparing request")
	}

	resp, err := client.GetByIDSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "policy.AssignmentsClient", "GetByID", resp, "Failure sending request")
	}

	result, err = client.GetByIDResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.AssignmentsClient", "GetByID", resp, "Failure responding to request")
	}

	return
}

// GetByIDPreparer prepares the GetByID request.
func (client AssignmentsClient) GetByIDPreparer(policyAssignmentID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"policyAssignmentId": policyAssignmentID,
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{policyAssignmentId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetByIDSender sends the GetByID request. The method will close the
// http.Response Body if it receives an error.
func (client AssignmentsClient) GetByIDSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetByIDResponder handles the response to the GetByID request. The method always
// closes the http.Response Body.
func (client AssignmentsClient) GetByIDResponder(resp *http.Response) (result Assignment, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets all the policy assignments for a subscription.
//
// filter is the filter to apply on the operation.
func (client AssignmentsClient) List(filter string) (result AssignmentListResult, err error) {
	req, err := client.ListPreparer(filter)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "policy.AssignmentsClient", "List", nil, "Failure preparing request")
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "policy.AssignmentsClient", "List", resp, "Failure sending request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.AssignmentsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client AssignmentsClient) ListPreparer(filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/policyassignments", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client AssignmentsClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client AssignmentsClient) ListResponder(resp *http.Response) (result AssignmentListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListNextResults retrieves the next set of results, if any.
func (client AssignmentsClient) ListNextResults(lastResults AssignmentListResult) (result AssignmentListResult, err error) {
	req, err := lastResults.AssignmentListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "policy.AssignmentsClient", "List", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "policy.AssignmentsClient", "List", resp, "Failure sending next results request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.AssignmentsClient", "List", resp, "Failure responding to next results request")
	}

	return
}

// ListForResource gets policy assignments for a resource.
//
// resourceGroupName is the name of the resource group containing the
// resource. The name is case insensitive. resourceProviderNamespace is the
// namespace of the resource provider. parentResourcePath is the parent
// resource path. resourceType is the resource type. resourceName is the name
// of the resource with policy assignments. filter is the filter to apply on
// the operation.
func (client AssignmentsClient) ListForResource(resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, filter string) (result AssignmentListResult, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "policy.AssignmentsClient", "ListForResource")
	}

	req, err := client.ListForResourcePreparer(resourceGroupName, resourceProviderNamespace, parentResourcePath, resourceType, resourceName, filter)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "policy.AssignmentsClient", "ListForResource", nil, "Failure preparing request")
	}

	resp, err := client.ListForResourceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "policy.AssignmentsClient", "ListForResource", resp, "Failure sending request")
	}

	result, err = client.ListForResourceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.AssignmentsClient", "ListForResource", resp, "Failure responding to request")
	}

	return
}

// ListForResourcePreparer prepares the ListForResource request.
func (client AssignmentsClient) ListForResourcePreparer(resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"parentResourcePath":        parentResourcePath,
		"resourceGroupName":         autorest.Encode("path", resourceGroupName),
		"resourceName":              autorest.Encode("path", resourceName),
		"resourceProviderNamespace": autorest.Encode("path", resourceProviderNamespace),
		"resourceType":              resourceType,
		"subscriptionId":            autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{parentResourcePath}/{resourceType}/{resourceName}/providers/Microsoft.Authorization/policyassignments", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListForResourceSender sends the ListForResource request. The method will close the
// http.Response Body if it receives an error.
func (client AssignmentsClient) ListForResourceSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListForResourceResponder handles the response to the ListForResource request. The method always
// closes the http.Response Body.
func (client AssignmentsClient) ListForResourceResponder(resp *http.Response) (result AssignmentListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListForResourceNextResults retrieves the next set of results, if any.
func (client AssignmentsClient) ListForResourceNextResults(lastResults AssignmentListResult) (result AssignmentListResult, err error) {
	req, err := lastResults.AssignmentListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "policy.AssignmentsClient", "ListForResource", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListForResourceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "policy.AssignmentsClient", "ListForResource", resp, "Failure sending next results request")
	}

	result, err = client.ListForResourceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.AssignmentsClient", "ListForResource", resp, "Failure responding to next results request")
	}

	return
}

// ListForResourceGroup gets policy assignments for the resource group.
//
// resourceGroupName is the name of the resource group that contains policy
// assignments. filter is the filter to apply on the operation.
func (client AssignmentsClient) ListForResourceGroup(resourceGroupName string, filter string) (result AssignmentListResult, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "policy.AssignmentsClient", "ListForResourceGroup")
	}

	req, err := client.ListForResourceGroupPreparer(resourceGroupName, filter)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "policy.AssignmentsClient", "ListForResourceGroup", nil, "Failure preparing request")
	}

	resp, err := client.ListForResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "policy.AssignmentsClient", "ListForResourceGroup", resp, "Failure sending request")
	}

	result, err = client.ListForResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.AssignmentsClient", "ListForResourceGroup", resp, "Failure responding to request")
	}

	return
}

// ListForResourceGroupPreparer prepares the ListForResourceGroup request.
func (client AssignmentsClient) ListForResourceGroupPreparer(resourceGroupName string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = filter
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Authorization/policyAssignments", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListForResourceGroupSender sends the ListForResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client AssignmentsClient) ListForResourceGroupSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListForResourceGroupResponder handles the response to the ListForResourceGroup request. The method always
// closes the http.Response Body.
func (client AssignmentsClient) ListForResourceGroupResponder(resp *http.Response) (result AssignmentListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListForResourceGroupNextResults retrieves the next set of results, if any.
func (client AssignmentsClient) ListForResourceGroupNextResults(lastResults AssignmentListResult) (result AssignmentListResult, err error) {
	req, err := lastResults.AssignmentListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "policy.AssignmentsClient", "ListForResourceGroup", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListForResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "policy.AssignmentsClient", "ListForResourceGroup", resp, "Failure sending next results request")
	}

	result, err = client.ListForResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.AssignmentsClient", "ListForResourceGroup", resp, "Failure responding to next results request")
	}

	return
}
