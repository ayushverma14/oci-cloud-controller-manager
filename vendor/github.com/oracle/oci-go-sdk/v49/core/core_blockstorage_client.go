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
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v49/common"
	"github.com/oracle/oci-go-sdk/v49/common/auth"
	"net/http"
)

//BlockstorageClient a client for Blockstorage
type BlockstorageClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewBlockstorageClientWithConfigurationProvider Creates a new default Blockstorage client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewBlockstorageClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client BlockstorageClient, err error) {
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newBlockstorageClientFromBaseClient(baseClient, provider)
}

// NewBlockstorageClientWithOboToken Creates a new default Blockstorage client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewBlockstorageClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client BlockstorageClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newBlockstorageClientFromBaseClient(baseClient, configProvider)
}

func newBlockstorageClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client BlockstorageClient, err error) {
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = BlockstorageClient{BaseClient: baseClient}
	client.BasePath = "20160918"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *BlockstorageClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("iaas", "https://iaas.{region}.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *BlockstorageClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
	if ok, err := common.IsConfigurationProviderValid(configProvider); !ok {
		return err
	}

	// Error has been checked already
	region, _ := configProvider.Region()
	client.SetRegion(region)
	client.config = &configProvider
	return nil
}

// ConfigurationProvider the ConfigurationProvider used in this client, or null if none set
func (client *BlockstorageClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangeBootVolumeBackupCompartment Moves a boot volume backup into a different compartment within the same tenancy.
// For information about moving resources between compartments,
// see Moving Resources to a Different Compartment (https://docs.cloud.oracle.com/iaas/Content/Identity/Tasks/managingcompartments.htm#moveRes).
func (client BlockstorageClient) ChangeBootVolumeBackupCompartment(ctx context.Context, request ChangeBootVolumeBackupCompartmentRequest) (response ChangeBootVolumeBackupCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.changeBootVolumeBackupCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeBootVolumeBackupCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeBootVolumeBackupCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeBootVolumeBackupCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeBootVolumeBackupCompartmentResponse")
	}
	return
}

// changeBootVolumeBackupCompartment implements the OCIOperation interface (enables retrying operations)
func (client BlockstorageClient) changeBootVolumeBackupCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/bootVolumeBackups/{bootVolumeBackupId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeBootVolumeBackupCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeBootVolumeCompartment Moves a boot volume into a different compartment within the same tenancy.
// For information about moving resources between compartments,
// see Moving Resources to a Different Compartment (https://docs.cloud.oracle.com/iaas/Content/Identity/Tasks/managingcompartments.htm#moveRes).
func (client BlockstorageClient) ChangeBootVolumeCompartment(ctx context.Context, request ChangeBootVolumeCompartmentRequest) (response ChangeBootVolumeCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.changeBootVolumeCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeBootVolumeCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeBootVolumeCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeBootVolumeCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeBootVolumeCompartmentResponse")
	}
	return
}

// changeBootVolumeCompartment implements the OCIOperation interface (enables retrying operations)
func (client BlockstorageClient) changeBootVolumeCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/bootVolumes/{bootVolumeId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeBootVolumeCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeVolumeBackupCompartment Moves a volume backup into a different compartment within the same tenancy.
// For information about moving resources between compartments,
// see Moving Resources to a Different Compartment (https://docs.cloud.oracle.com/iaas/Content/Identity/Tasks/managingcompartments.htm#moveRes).
func (client BlockstorageClient) ChangeVolumeBackupCompartment(ctx context.Context, request ChangeVolumeBackupCompartmentRequest) (response ChangeVolumeBackupCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.changeVolumeBackupCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeVolumeBackupCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeVolumeBackupCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeVolumeBackupCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeVolumeBackupCompartmentResponse")
	}
	return
}

// changeVolumeBackupCompartment implements the OCIOperation interface (enables retrying operations)
func (client BlockstorageClient) changeVolumeBackupCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/volumeBackups/{volumeBackupId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeVolumeBackupCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeVolumeCompartment Moves a volume into a different compartment within the same tenancy.
// For information about moving resources between compartments,
// see Moving Resources to a Different Compartment (https://docs.cloud.oracle.com/iaas/Content/Identity/Tasks/managingcompartments.htm#moveRes).
func (client BlockstorageClient) ChangeVolumeCompartment(ctx context.Context, request ChangeVolumeCompartmentRequest) (response ChangeVolumeCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.changeVolumeCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeVolumeCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeVolumeCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeVolumeCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeVolumeCompartmentResponse")
	}
	return
}

// changeVolumeCompartment implements the OCIOperation interface (enables retrying operations)
func (client BlockstorageClient) changeVolumeCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/volumes/{volumeId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeVolumeCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeVolumeGroupBackupCompartment Moves a volume group backup into a different compartment within the same tenancy.
// For information about moving resources between compartments,
// see Moving Resources to a Different Compartment (https://docs.cloud.oracle.com/iaas/Content/Identity/Tasks/managingcompartments.htm#moveRes).
func (client BlockstorageClient) ChangeVolumeGroupBackupCompartment(ctx context.Context, request ChangeVolumeGroupBackupCompartmentRequest) (response ChangeVolumeGroupBackupCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.changeVolumeGroupBackupCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeVolumeGroupBackupCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeVolumeGroupBackupCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeVolumeGroupBackupCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeVolumeGroupBackupCompartmentResponse")
	}
	return
}

// changeVolumeGroupBackupCompartment implements the OCIOperation interface (enables retrying operations)
func (client BlockstorageClient) changeVolumeGroupBackupCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/volumeGroupBackups/{volumeGroupBackupId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeVolumeGroupBackupCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeVolumeGroupCompartment Moves a volume group into a different compartment within the same tenancy.
// For information about moving resources between compartments,
// see Moving Resources to a Different Compartment (https://docs.cloud.oracle.com/iaas/Content/Identity/Tasks/managingcompartments.htm#moveRes).
func (client BlockstorageClient) ChangeVolumeGroupCompartment(ctx context.Context, request ChangeVolumeGroupCompartmentRequest) (response ChangeVolumeGroupCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.changeVolumeGroupCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeVolumeGroupCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeVolumeGroupCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeVolumeGroupCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeVolumeGroupCompartmentResponse")
	}
	return
}

// changeVolumeGroupCompartment implements the OCIOperation interface (enables retrying operations)
func (client BlockstorageClient) changeVolumeGroupCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/volumeGroups/{volumeGroupId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeVolumeGroupCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CopyBootVolumeBackup Creates a boot volume backup copy in specified region. For general information about volume backups,
// see Overview of Boot Volume Backups (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/bootvolumebackups.htm)
func (client BlockstorageClient) CopyBootVolumeBackup(ctx context.Context, request CopyBootVolumeBackupRequest) (response CopyBootVolumeBackupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.copyBootVolumeBackup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CopyBootVolumeBackupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CopyBootVolumeBackupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CopyBootVolumeBackupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CopyBootVolumeBackupResponse")
	}
	return
}

// copyBootVolumeBackup implements the OCIOperation interface (enables retrying operations)
func (client BlockstorageClient) copyBootVolumeBackup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/bootVolumeBackups/{bootVolumeBackupId}/actions/copy", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CopyBootVolumeBackupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CopyVolumeBackup Creates a volume backup copy in specified region. For general information about volume backups,
// see Overview of Block Volume Service Backups (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/blockvolumebackups.htm)
func (client BlockstorageClient) CopyVolumeBackup(ctx context.Context, request CopyVolumeBackupRequest) (response CopyVolumeBackupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.copyVolumeBackup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CopyVolumeBackupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CopyVolumeBackupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CopyVolumeBackupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CopyVolumeBackupResponse")
	}
	return
}

// copyVolumeBackup implements the OCIOperation interface (enables retrying operations)
func (client BlockstorageClient) copyVolumeBackup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/volumeBackups/{volumeBackupId}/actions/copy", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CopyVolumeBackupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CopyVolumeGroupBackup Creates a volume group backup copy in specified region. For general information about volume group backups,
// see Overview of Block Volume Service Backups (https://docs.cloud.oracle.com/Content/Block/Concepts/blockvolumebackups.htm)
func (client BlockstorageClient) CopyVolumeGroupBackup(ctx context.Context, request CopyVolumeGroupBackupRequest) (response CopyVolumeGroupBackupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.copyVolumeGroupBackup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CopyVolumeGroupBackupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CopyVolumeGroupBackupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CopyVolumeGroupBackupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CopyVolumeGroupBackupResponse")
	}
	return
}

// copyVolumeGroupBackup implements the OCIOperation interface (enables retrying operations)
func (client BlockstorageClient) copyVolumeGroupBackup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/volumeGroupBackups/{volumeGroupBackupId}/actions/copy", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CopyVolumeGroupBackupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateBootVolume Creates a new boot volume in the specified compartment from an existing boot volume or a boot volume backup.
// For general information about boot volumes, see Boot Volumes (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/bootvolumes.htm).
// You may optionally specify a *display name* for the volume, which is simply a friendly name or
// description. It does not have to be unique, and you can change it. Avoid entering confidential information.
func (client BlockstorageClient) CreateBootVolume(ctx context.Context, request CreateBootVolumeRequest) (response CreateBootVolumeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createBootVolume, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateBootVolumeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateBootVolumeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateBootVolumeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateBootVolumeResponse")
	}
	return
}

// createBootVolume implements the OCIOperation interface (enables retrying operations)
func (client BlockstorageClient) createBootVolume(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/bootVolumes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateBootVolumeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateBootVolumeBackup Creates a new boot volume backup of the specified boot volume. For general information about boot volume backups,
// see Overview of Boot Volume Backups (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/bootvolumebackups.htm)
// When the request is received, the backup object is in a REQUEST_RECEIVED state.
// When the data is imaged, it goes into a CREATING state.
// After the backup is fully uploaded to the cloud, it goes into an AVAILABLE state.
func (client BlockstorageClient) CreateBootVolumeBackup(ctx context.Context, request CreateBootVolumeBackupRequest) (response CreateBootVolumeBackupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createBootVolumeBackup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateBootVolumeBackupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateBootVolumeBackupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateBootVolumeBackupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateBootVolumeBackupResponse")
	}
	return
}

// createBootVolumeBackup implements the OCIOperation interface (enables retrying operations)
func (client BlockstorageClient) createBootVolumeBackup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/bootVolumeBackups", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateBootVolumeBackupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateVolume Creates a new volume in the specified compartment. Volumes can be created in sizes ranging from
// 50 GB (51200 MB) to 32 TB (33554432 MB), in 1 GB (1024 MB) increments. By default, volumes are 1 TB (1048576 MB).
// For general information about block volumes, see
// Overview of Block Volume Service (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm).
// A volume and instance can be in separate compartments but must be in the same availability domain.
// For information about access control and compartments, see
// Overview of the IAM Service (https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/overview.htm). For information about
// availability domains, see Regions and Availability Domains (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/regions.htm).
// To get a list of availability domains, use the `ListAvailabilityDomains` operation
// in the Identity and Access Management Service API.
// You may optionally specify a *display name* for the volume, which is simply a friendly name or
// description. It does not have to be unique, and you can change it. Avoid entering confidential information.
func (client BlockstorageClient) CreateVolume(ctx context.Context, request CreateVolumeRequest) (response CreateVolumeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createVolume, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateVolumeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateVolumeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateVolumeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateVolumeResponse")
	}
	return
}

// createVolume implements the OCIOperation interface (enables retrying operations)
func (client BlockstorageClient) createVolume(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/volumes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateVolumeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateVolumeBackup Creates a new backup of the specified volume. For general information about volume backups,
// see Overview of Block Volume Service Backups (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/blockvolumebackups.htm)
// When the request is received, the backup object is in a REQUEST_RECEIVED state.
// When the data is imaged, it goes into a CREATING state.
// After the backup is fully uploaded to the cloud, it goes into an AVAILABLE state.
func (client BlockstorageClient) CreateVolumeBackup(ctx context.Context, request CreateVolumeBackupRequest) (response CreateVolumeBackupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createVolumeBackup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateVolumeBackupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateVolumeBackupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateVolumeBackupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateVolumeBackupResponse")
	}
	return
}

// createVolumeBackup implements the OCIOperation interface (enables retrying operations)
func (client BlockstorageClient) createVolumeBackup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/volumeBackups", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateVolumeBackupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateVolumeBackupPolicy Creates a new user defined backup policy.
// For more information about Oracle defined backup policies and user defined backup policies,
// see Policy-Based Backups (https://docs.cloud.oracle.com/iaas/Content/Block/Tasks/schedulingvolumebackups.htm).
func (client BlockstorageClient) CreateVolumeBackupPolicy(ctx context.Context, request CreateVolumeBackupPolicyRequest) (response CreateVolumeBackupPolicyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createVolumeBackupPolicy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateVolumeBackupPolicyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateVolumeBackupPolicyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateVolumeBackupPolicyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateVolumeBackupPolicyResponse")
	}
	return
}

// createVolumeBackupPolicy implements the OCIOperation interface (enables retrying operations)
func (client BlockstorageClient) createVolumeBackupPolicy(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/volumeBackupPolicies", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateVolumeBackupPolicyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateVolumeBackupPolicyAssignment Assigns a volume backup policy to the specified volume. Note that a given volume can
// only have one backup policy assigned to it. If this operation is used for a volume that already
// has a different backup policy assigned, the prior backup policy will be silently unassigned.
func (client BlockstorageClient) CreateVolumeBackupPolicyAssignment(ctx context.Context, request CreateVolumeBackupPolicyAssignmentRequest) (response CreateVolumeBackupPolicyAssignmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.createVolumeBackupPolicyAssignment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateVolumeBackupPolicyAssignmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateVolumeBackupPolicyAssignmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateVolumeBackupPolicyAssignmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateVolumeBackupPolicyAssignmentResponse")
	}
	return
}

// createVolumeBackupPolicyAssignment implements the OCIOperation interface (enables retrying operations)
func (client BlockstorageClient) createVolumeBackupPolicyAssignment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/volumeBackupPolicyAssignments", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateVolumeBackupPolicyAssignmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateVolumeGroup Creates a new volume group in the specified compartment.
// A volume group is a collection of volumes and may be created from a list of volumes, cloning an existing
// volume group, or by restoring a volume group backup.
// You may optionally specify a *display name* for the volume group, which is simply a friendly name or
// description. It does not have to be unique, and you can change it. Avoid entering confidential information.
// For more information, see Volume Groups (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/volumegroups.htm).
func (client BlockstorageClient) CreateVolumeGroup(ctx context.Context, request CreateVolumeGroupRequest) (response CreateVolumeGroupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createVolumeGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateVolumeGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateVolumeGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateVolumeGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateVolumeGroupResponse")
	}
	return
}

// createVolumeGroup implements the OCIOperation interface (enables retrying operations)
func (client BlockstorageClient) createVolumeGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/volumeGroups", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateVolumeGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateVolumeGroupBackup Creates a new backup volume group of the specified volume group.
// For more information, see Volume Groups (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/volumegroups.htm).
func (client BlockstorageClient) CreateVolumeGroupBackup(ctx context.Context, request CreateVolumeGroupBackupRequest) (response CreateVolumeGroupBackupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createVolumeGroupBackup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateVolumeGroupBackupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateVolumeGroupBackupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateVolumeGroupBackupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateVolumeGroupBackupResponse")
	}
	return
}

// createVolumeGroupBackup implements the OCIOperation interface (enables retrying operations)
func (client BlockstorageClient) createVolumeGroupBackup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/volumeGroupBackups", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateVolumeGroupBackupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteBootVolume Deletes the specified boot volume. The volume cannot have an active connection to an instance.
// To disconnect the boot volume from a connected instance, see
// Disconnecting From a Boot Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Tasks/deletingbootvolume.htm).
// **Warning:** All data on the boot volume will be permanently lost when the boot volume is deleted.
func (client BlockstorageClient) DeleteBootVolume(ctx context.Context, request DeleteBootVolumeRequest) (response DeleteBootVolumeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteBootVolume, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteBootVolumeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteBootVolumeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteBootVolumeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteBootVolumeResponse")
	}
	return
}

// deleteBootVolume implements the OCIOperation interface (enables retrying operations)
func (client BlockstorageClient) deleteBootVolume(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/bootVolumes/{bootVolumeId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteBootVolumeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteBootVolumeBackup Deletes a boot volume backup.
func (client BlockstorageClient) DeleteBootVolumeBackup(ctx context.Context, request DeleteBootVolumeBackupRequest) (response DeleteBootVolumeBackupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteBootVolumeBackup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteBootVolumeBackupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteBootVolumeBackupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteBootVolumeBackupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteBootVolumeBackupResponse")
	}
	return
}

// deleteBootVolumeBackup implements the OCIOperation interface (enables retrying operations)
func (client BlockstorageClient) deleteBootVolumeBackup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/bootVolumeBackups/{bootVolumeBackupId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteBootVolumeBackupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteBootVolumeKmsKey Removes the specified boot volume's assigned Key Management encryption key.
func (client BlockstorageClient) DeleteBootVolumeKmsKey(ctx context.Context, request DeleteBootVolumeKmsKeyRequest) (response DeleteBootVolumeKmsKeyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteBootVolumeKmsKey, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteBootVolumeKmsKeyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteBootVolumeKmsKeyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteBootVolumeKmsKeyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteBootVolumeKmsKeyResponse")
	}
	return
}

// deleteBootVolumeKmsKey implements the OCIOperation interface (enables retrying operations)
func (client BlockstorageClient) deleteBootVolumeKmsKey(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/bootVolumes/{bootVolumeId}/kmsKey", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteBootVolumeKmsKeyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteVolume Deletes the specified volume. The volume cannot have an active connection to an instance.
// To disconnect the volume from a connected instance, see
// Disconnecting From a Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Tasks/disconnectingfromavolume.htm).
// **Warning:** All data on the volume will be permanently lost when the volume is deleted.
func (client BlockstorageClient) DeleteVolume(ctx context.Context, request DeleteVolumeRequest) (response DeleteVolumeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteVolume, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteVolumeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteVolumeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteVolumeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteVolumeResponse")
	}
	return
}

// deleteVolume implements the OCIOperation interface (enables retrying operations)
func (client BlockstorageClient) deleteVolume(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/volumes/{volumeId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteVolumeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteVolumeBackup Deletes a volume backup.
func (client BlockstorageClient) DeleteVolumeBackup(ctx context.Context, request DeleteVolumeBackupRequest) (response DeleteVolumeBackupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteVolumeBackup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteVolumeBackupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteVolumeBackupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteVolumeBackupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteVolumeBackupResponse")
	}
	return
}

// deleteVolumeBackup implements the OCIOperation interface (enables retrying operations)
func (client BlockstorageClient) deleteVolumeBackup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/volumeBackups/{volumeBackupId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteVolumeBackupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteVolumeBackupPolicy Deletes a user defined backup policy.
//  For more information about user defined backup policies,
//  see Policy-Based Backups (https://docs.cloud.oracle.com/iaas/Content/Block/Tasks/schedulingvolumebackups.htm#UserDefinedBackupPolicies).
//  Avoid entering confidential information.
func (client BlockstorageClient) DeleteVolumeBackupPolicy(ctx context.Context, request DeleteVolumeBackupPolicyRequest) (response DeleteVolumeBackupPolicyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteVolumeBackupPolicy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteVolumeBackupPolicyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteVolumeBackupPolicyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteVolumeBackupPolicyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteVolumeBackupPolicyResponse")
	}
	return
}

// deleteVolumeBackupPolicy implements the OCIOperation interface (enables retrying operations)
func (client BlockstorageClient) deleteVolumeBackupPolicy(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/volumeBackupPolicies/{policyId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteVolumeBackupPolicyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteVolumeBackupPolicyAssignment Deletes a volume backup policy assignment.
func (client BlockstorageClient) DeleteVolumeBackupPolicyAssignment(ctx context.Context, request DeleteVolumeBackupPolicyAssignmentRequest) (response DeleteVolumeBackupPolicyAssignmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteVolumeBackupPolicyAssignment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteVolumeBackupPolicyAssignmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteVolumeBackupPolicyAssignmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteVolumeBackupPolicyAssignmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteVolumeBackupPolicyAssignmentResponse")
	}
	return
}

// deleteVolumeBackupPolicyAssignment implements the OCIOperation interface (enables retrying operations)
func (client BlockstorageClient) deleteVolumeBackupPolicyAssignment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/volumeBackupPolicyAssignments/{policyAssignmentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteVolumeBackupPolicyAssignmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteVolumeGroup Deletes the specified volume group. Individual volumes are not deleted, only the volume group is deleted.
// For more information, see Volume Groups (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/volumegroups.htm).
func (client BlockstorageClient) DeleteVolumeGroup(ctx context.Context, request DeleteVolumeGroupRequest) (response DeleteVolumeGroupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteVolumeGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteVolumeGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteVolumeGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteVolumeGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteVolumeGroupResponse")
	}
	return
}

// deleteVolumeGroup implements the OCIOperation interface (enables retrying operations)
func (client BlockstorageClient) deleteVolumeGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/volumeGroups/{volumeGroupId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteVolumeGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteVolumeGroupBackup Deletes a volume group backup. This operation deletes all the backups in
// the volume group. For more information, see Volume Groups (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/volumegroups.htm).
func (client BlockstorageClient) DeleteVolumeGroupBackup(ctx context.Context, request DeleteVolumeGroupBackupRequest) (response DeleteVolumeGroupBackupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteVolumeGroupBackup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteVolumeGroupBackupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteVolumeGroupBackupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteVolumeGroupBackupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteVolumeGroupBackupResponse")
	}
	return
}

// deleteVolumeGroupBackup implements the OCIOperation interface (enables retrying operations)
func (client BlockstorageClient) deleteVolumeGroupBackup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/volumeGroupBackups/{volumeGroupBackupId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteVolumeGroupBackupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteVolumeKmsKey Removes the specified volume's assigned Key Management encryption key.
func (client BlockstorageClient) DeleteVolumeKmsKey(ctx context.Context, request DeleteVolumeKmsKeyRequest) (response DeleteVolumeKmsKeyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteVolumeKmsKey, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteVolumeKmsKeyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteVolumeKmsKeyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteVolumeKmsKeyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteVolumeKmsKeyResponse")
	}
	return
}

// deleteVolumeKmsKey implements the OCIOperation interface (enables retrying operations)
func (client BlockstorageClient) deleteVolumeKmsKey(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/volumes/{volumeId}/kmsKey", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteVolumeKmsKeyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetBlockVolumeReplica Gets information for the specified block volume replica.
func (client BlockstorageClient) GetBlockVolumeReplica(ctx context.Context, request GetBlockVolumeReplicaRequest) (response GetBlockVolumeReplicaResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getBlockVolumeReplica, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetBlockVolumeReplicaResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetBlockVolumeReplicaResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetBlockVolumeReplicaResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetBlockVolumeReplicaResponse")
	}
	return
}

// getBlockVolumeReplica implements the OCIOperation interface (enables retrying operations)
func (client BlockstorageClient) getBlockVolumeReplica(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/blockVolumeReplicas/{blockVolumeReplicaId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetBlockVolumeReplicaResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetBootVolume Gets information for the specified boot volume.
func (client BlockstorageClient) GetBootVolume(ctx context.Context, request GetBootVolumeRequest) (response GetBootVolumeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getBootVolume, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetBootVolumeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetBootVolumeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetBootVolumeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetBootVolumeResponse")
	}
	return
}

// getBootVolume implements the OCIOperation interface (enables retrying operations)
func (client BlockstorageClient) getBootVolume(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/bootVolumes/{bootVolumeId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetBootVolumeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetBootVolumeBackup Gets information for the specified boot volume backup.
func (client BlockstorageClient) GetBootVolumeBackup(ctx context.Context, request GetBootVolumeBackupRequest) (response GetBootVolumeBackupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getBootVolumeBackup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetBootVolumeBackupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetBootVolumeBackupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetBootVolumeBackupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetBootVolumeBackupResponse")
	}
	return
}

// getBootVolumeBackup implements the OCIOperation interface (enables retrying operations)
func (client BlockstorageClient) getBootVolumeBackup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/bootVolumeBackups/{bootVolumeBackupId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetBootVolumeBackupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetBootVolumeKmsKey Gets the Key Management encryption key assigned to the specified boot volume.
func (client BlockstorageClient) GetBootVolumeKmsKey(ctx context.Context, request GetBootVolumeKmsKeyRequest) (response GetBootVolumeKmsKeyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getBootVolumeKmsKey, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetBootVolumeKmsKeyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetBootVolumeKmsKeyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetBootVolumeKmsKeyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetBootVolumeKmsKeyResponse")
	}
	return
}

// getBootVolumeKmsKey implements the OCIOperation interface (enables retrying operations)
func (client BlockstorageClient) getBootVolumeKmsKey(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/bootVolumes/{bootVolumeId}/kmsKey", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetBootVolumeKmsKeyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetBootVolumeReplica Gets information for the specified boot volume replica.
func (client BlockstorageClient) GetBootVolumeReplica(ctx context.Context, request GetBootVolumeReplicaRequest) (response GetBootVolumeReplicaResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getBootVolumeReplica, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetBootVolumeReplicaResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetBootVolumeReplicaResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetBootVolumeReplicaResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetBootVolumeReplicaResponse")
	}
	return
}

// getBootVolumeReplica implements the OCIOperation interface (enables retrying operations)
func (client BlockstorageClient) getBootVolumeReplica(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/bootVolumeReplicas/{bootVolumeReplicaId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetBootVolumeReplicaResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetVolume Gets information for the specified volume.
func (client BlockstorageClient) GetVolume(ctx context.Context, request GetVolumeRequest) (response GetVolumeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getVolume, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetVolumeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetVolumeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetVolumeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetVolumeResponse")
	}
	return
}

// getVolume implements the OCIOperation interface (enables retrying operations)
func (client BlockstorageClient) getVolume(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/volumes/{volumeId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetVolumeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetVolumeBackup Gets information for the specified volume backup.
func (client BlockstorageClient) GetVolumeBackup(ctx context.Context, request GetVolumeBackupRequest) (response GetVolumeBackupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getVolumeBackup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetVolumeBackupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetVolumeBackupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetVolumeBackupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetVolumeBackupResponse")
	}
	return
}

// getVolumeBackup implements the OCIOperation interface (enables retrying operations)
func (client BlockstorageClient) getVolumeBackup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/volumeBackups/{volumeBackupId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetVolumeBackupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetVolumeBackupPolicy Gets information for the specified volume backup policy.
func (client BlockstorageClient) GetVolumeBackupPolicy(ctx context.Context, request GetVolumeBackupPolicyRequest) (response GetVolumeBackupPolicyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getVolumeBackupPolicy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetVolumeBackupPolicyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetVolumeBackupPolicyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetVolumeBackupPolicyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetVolumeBackupPolicyResponse")
	}
	return
}

// getVolumeBackupPolicy implements the OCIOperation interface (enables retrying operations)
func (client BlockstorageClient) getVolumeBackupPolicy(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/volumeBackupPolicies/{policyId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetVolumeBackupPolicyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetVolumeBackupPolicyAssetAssignment Gets the volume backup policy assignment for the specified volume. The
// `assetId` query parameter is required, and the returned list will contain at most
// one item, since volume can only have one volume backup policy assigned at a time.
func (client BlockstorageClient) GetVolumeBackupPolicyAssetAssignment(ctx context.Context, request GetVolumeBackupPolicyAssetAssignmentRequest) (response GetVolumeBackupPolicyAssetAssignmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getVolumeBackupPolicyAssetAssignment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetVolumeBackupPolicyAssetAssignmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetVolumeBackupPolicyAssetAssignmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetVolumeBackupPolicyAssetAssignmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetVolumeBackupPolicyAssetAssignmentResponse")
	}
	return
}

// getVolumeBackupPolicyAssetAssignment implements the OCIOperation interface (enables retrying operations)
func (client BlockstorageClient) getVolumeBackupPolicyAssetAssignment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/volumeBackupPolicyAssignments", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetVolumeBackupPolicyAssetAssignmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetVolumeBackupPolicyAssignment Gets information for the specified volume backup policy assignment.
func (client BlockstorageClient) GetVolumeBackupPolicyAssignment(ctx context.Context, request GetVolumeBackupPolicyAssignmentRequest) (response GetVolumeBackupPolicyAssignmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getVolumeBackupPolicyAssignment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetVolumeBackupPolicyAssignmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetVolumeBackupPolicyAssignmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetVolumeBackupPolicyAssignmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetVolumeBackupPolicyAssignmentResponse")
	}
	return
}

// getVolumeBackupPolicyAssignment implements the OCIOperation interface (enables retrying operations)
func (client BlockstorageClient) getVolumeBackupPolicyAssignment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/volumeBackupPolicyAssignments/{policyAssignmentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetVolumeBackupPolicyAssignmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetVolumeGroup Gets information for the specified volume group. For more information, see Volume Groups (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/volumegroups.htm).
func (client BlockstorageClient) GetVolumeGroup(ctx context.Context, request GetVolumeGroupRequest) (response GetVolumeGroupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getVolumeGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetVolumeGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetVolumeGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetVolumeGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetVolumeGroupResponse")
	}
	return
}

// getVolumeGroup implements the OCIOperation interface (enables retrying operations)
func (client BlockstorageClient) getVolumeGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/volumeGroups/{volumeGroupId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetVolumeGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetVolumeGroupBackup Gets information for the specified volume group backup. For more information, see Volume Groups (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/volumegroups.htm).
func (client BlockstorageClient) GetVolumeGroupBackup(ctx context.Context, request GetVolumeGroupBackupRequest) (response GetVolumeGroupBackupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getVolumeGroupBackup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetVolumeGroupBackupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetVolumeGroupBackupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetVolumeGroupBackupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetVolumeGroupBackupResponse")
	}
	return
}

// getVolumeGroupBackup implements the OCIOperation interface (enables retrying operations)
func (client BlockstorageClient) getVolumeGroupBackup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/volumeGroupBackups/{volumeGroupBackupId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetVolumeGroupBackupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetVolumeGroupReplica Gets information for the specified volume group replica.
func (client BlockstorageClient) GetVolumeGroupReplica(ctx context.Context, request GetVolumeGroupReplicaRequest) (response GetVolumeGroupReplicaResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getVolumeGroupReplica, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetVolumeGroupReplicaResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetVolumeGroupReplicaResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetVolumeGroupReplicaResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetVolumeGroupReplicaResponse")
	}
	return
}

// getVolumeGroupReplica implements the OCIOperation interface (enables retrying operations)
func (client BlockstorageClient) getVolumeGroupReplica(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/volumeGroupReplicas/{volumeGroupReplicaId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetVolumeGroupReplicaResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetVolumeKmsKey Gets the Key Management encryption key assigned to the specified volume.
func (client BlockstorageClient) GetVolumeKmsKey(ctx context.Context, request GetVolumeKmsKeyRequest) (response GetVolumeKmsKeyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getVolumeKmsKey, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetVolumeKmsKeyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetVolumeKmsKeyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetVolumeKmsKeyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetVolumeKmsKeyResponse")
	}
	return
}

// getVolumeKmsKey implements the OCIOperation interface (enables retrying operations)
func (client BlockstorageClient) getVolumeKmsKey(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/volumes/{volumeId}/kmsKey", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetVolumeKmsKeyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListBlockVolumeReplicas Lists the block volume replicas in the specified compartment and availability domain.
func (client BlockstorageClient) ListBlockVolumeReplicas(ctx context.Context, request ListBlockVolumeReplicasRequest) (response ListBlockVolumeReplicasResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listBlockVolumeReplicas, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListBlockVolumeReplicasResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListBlockVolumeReplicasResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListBlockVolumeReplicasResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListBlockVolumeReplicasResponse")
	}
	return
}

// listBlockVolumeReplicas implements the OCIOperation interface (enables retrying operations)
func (client BlockstorageClient) listBlockVolumeReplicas(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/blockVolumeReplicas", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListBlockVolumeReplicasResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListBootVolumeBackups Lists the boot volume backups in the specified compartment. You can filter the results by boot volume.
func (client BlockstorageClient) ListBootVolumeBackups(ctx context.Context, request ListBootVolumeBackupsRequest) (response ListBootVolumeBackupsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listBootVolumeBackups, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListBootVolumeBackupsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListBootVolumeBackupsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListBootVolumeBackupsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListBootVolumeBackupsResponse")
	}
	return
}

// listBootVolumeBackups implements the OCIOperation interface (enables retrying operations)
func (client BlockstorageClient) listBootVolumeBackups(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/bootVolumeBackups", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListBootVolumeBackupsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListBootVolumeReplicas Lists the boot volume replicas in the specified compartment and availability domain.
func (client BlockstorageClient) ListBootVolumeReplicas(ctx context.Context, request ListBootVolumeReplicasRequest) (response ListBootVolumeReplicasResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listBootVolumeReplicas, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListBootVolumeReplicasResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListBootVolumeReplicasResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListBootVolumeReplicasResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListBootVolumeReplicasResponse")
	}
	return
}

// listBootVolumeReplicas implements the OCIOperation interface (enables retrying operations)
func (client BlockstorageClient) listBootVolumeReplicas(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/bootVolumeReplicas", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListBootVolumeReplicasResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListBootVolumes Lists the boot volumes in the specified compartment and availability domain.
func (client BlockstorageClient) ListBootVolumes(ctx context.Context, request ListBootVolumesRequest) (response ListBootVolumesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listBootVolumes, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListBootVolumesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListBootVolumesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListBootVolumesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListBootVolumesResponse")
	}
	return
}

// listBootVolumes implements the OCIOperation interface (enables retrying operations)
func (client BlockstorageClient) listBootVolumes(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/bootVolumes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListBootVolumesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListVolumeBackupPolicies Lists all the volume backup policies available in the specified compartment.
// For more information about Oracle defined backup policies and user defined backup policies,
// see Policy-Based Backups (https://docs.cloud.oracle.com/iaas/Content/Block/Tasks/schedulingvolumebackups.htm).
func (client BlockstorageClient) ListVolumeBackupPolicies(ctx context.Context, request ListVolumeBackupPoliciesRequest) (response ListVolumeBackupPoliciesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listVolumeBackupPolicies, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListVolumeBackupPoliciesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListVolumeBackupPoliciesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListVolumeBackupPoliciesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListVolumeBackupPoliciesResponse")
	}
	return
}

// listVolumeBackupPolicies implements the OCIOperation interface (enables retrying operations)
func (client BlockstorageClient) listVolumeBackupPolicies(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/volumeBackupPolicies", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListVolumeBackupPoliciesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListVolumeBackups Lists the volume backups in the specified compartment. You can filter the results by volume.
func (client BlockstorageClient) ListVolumeBackups(ctx context.Context, request ListVolumeBackupsRequest) (response ListVolumeBackupsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listVolumeBackups, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListVolumeBackupsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListVolumeBackupsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListVolumeBackupsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListVolumeBackupsResponse")
	}
	return
}

// listVolumeBackups implements the OCIOperation interface (enables retrying operations)
func (client BlockstorageClient) listVolumeBackups(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/volumeBackups", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListVolumeBackupsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListVolumeGroupBackups Lists the volume group backups in the specified compartment. You can filter the results by volume group.
// For more information, see Volume Groups (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/volumegroups.htm).
func (client BlockstorageClient) ListVolumeGroupBackups(ctx context.Context, request ListVolumeGroupBackupsRequest) (response ListVolumeGroupBackupsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listVolumeGroupBackups, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListVolumeGroupBackupsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListVolumeGroupBackupsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListVolumeGroupBackupsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListVolumeGroupBackupsResponse")
	}
	return
}

// listVolumeGroupBackups implements the OCIOperation interface (enables retrying operations)
func (client BlockstorageClient) listVolumeGroupBackups(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/volumeGroupBackups", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListVolumeGroupBackupsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListVolumeGroupReplicas Lists the volume group replicas in the specified compartment. You can filter the results by volume group.
// For more information, see Volume Group Replication (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/volumegroupreplication.htm).
func (client BlockstorageClient) ListVolumeGroupReplicas(ctx context.Context, request ListVolumeGroupReplicasRequest) (response ListVolumeGroupReplicasResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listVolumeGroupReplicas, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListVolumeGroupReplicasResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListVolumeGroupReplicasResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListVolumeGroupReplicasResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListVolumeGroupReplicasResponse")
	}
	return
}

// listVolumeGroupReplicas implements the OCIOperation interface (enables retrying operations)
func (client BlockstorageClient) listVolumeGroupReplicas(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/volumeGroupReplicas", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListVolumeGroupReplicasResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListVolumeGroups Lists the volume groups in the specified compartment and availability domain.
// For more information, see Volume Groups (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/volumegroups.htm).
func (client BlockstorageClient) ListVolumeGroups(ctx context.Context, request ListVolumeGroupsRequest) (response ListVolumeGroupsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listVolumeGroups, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListVolumeGroupsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListVolumeGroupsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListVolumeGroupsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListVolumeGroupsResponse")
	}
	return
}

// listVolumeGroups implements the OCIOperation interface (enables retrying operations)
func (client BlockstorageClient) listVolumeGroups(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/volumeGroups", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListVolumeGroupsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListVolumes Lists the volumes in the specified compartment and availability domain.
func (client BlockstorageClient) ListVolumes(ctx context.Context, request ListVolumesRequest) (response ListVolumesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listVolumes, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListVolumesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListVolumesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListVolumesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListVolumesResponse")
	}
	return
}

// listVolumes implements the OCIOperation interface (enables retrying operations)
func (client BlockstorageClient) listVolumes(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/volumes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListVolumesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateBootVolume Updates the specified boot volume's display name, defined tags, and free-form tags.
func (client BlockstorageClient) UpdateBootVolume(ctx context.Context, request UpdateBootVolumeRequest) (response UpdateBootVolumeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateBootVolume, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateBootVolumeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateBootVolumeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateBootVolumeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateBootVolumeResponse")
	}
	return
}

// updateBootVolume implements the OCIOperation interface (enables retrying operations)
func (client BlockstorageClient) updateBootVolume(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/bootVolumes/{bootVolumeId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateBootVolumeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateBootVolumeBackup Updates the display name for the specified boot volume backup.
// Avoid entering confidential information.
func (client BlockstorageClient) UpdateBootVolumeBackup(ctx context.Context, request UpdateBootVolumeBackupRequest) (response UpdateBootVolumeBackupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateBootVolumeBackup, policy)
	if err != nil {
		if ociResponse != nil {
			response = UpdateBootVolumeBackupResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateBootVolumeBackupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateBootVolumeBackupResponse")
	}
	return
}

// updateBootVolumeBackup implements the OCIOperation interface (enables retrying operations)
func (client BlockstorageClient) updateBootVolumeBackup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/bootVolumeBackups/{bootVolumeBackupId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateBootVolumeBackupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateBootVolumeKmsKey Updates the specified volume with a new Key Management master encryption key.
func (client BlockstorageClient) UpdateBootVolumeKmsKey(ctx context.Context, request UpdateBootVolumeKmsKeyRequest) (response UpdateBootVolumeKmsKeyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateBootVolumeKmsKey, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateBootVolumeKmsKeyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateBootVolumeKmsKeyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateBootVolumeKmsKeyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateBootVolumeKmsKeyResponse")
	}
	return
}

// updateBootVolumeKmsKey implements the OCIOperation interface (enables retrying operations)
func (client BlockstorageClient) updateBootVolumeKmsKey(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/bootVolumes/{bootVolumeId}/kmsKey", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateBootVolumeKmsKeyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateVolume Updates the specified volume's display name.
// Avoid entering confidential information.
func (client BlockstorageClient) UpdateVolume(ctx context.Context, request UpdateVolumeRequest) (response UpdateVolumeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateVolume, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateVolumeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateVolumeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateVolumeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateVolumeResponse")
	}
	return
}

// updateVolume implements the OCIOperation interface (enables retrying operations)
func (client BlockstorageClient) updateVolume(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/volumes/{volumeId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateVolumeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateVolumeBackup Updates the display name for the specified volume backup.
// Avoid entering confidential information.
func (client BlockstorageClient) UpdateVolumeBackup(ctx context.Context, request UpdateVolumeBackupRequest) (response UpdateVolumeBackupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateVolumeBackup, policy)
	if err != nil {
		if ociResponse != nil {
			response = UpdateVolumeBackupResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateVolumeBackupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateVolumeBackupResponse")
	}
	return
}

// updateVolumeBackup implements the OCIOperation interface (enables retrying operations)
func (client BlockstorageClient) updateVolumeBackup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/volumeBackups/{volumeBackupId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateVolumeBackupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateVolumeBackupPolicy Updates a user defined backup policy.
//  For more information about user defined backup policies,
//  see Policy-Based Backups (https://docs.cloud.oracle.com/iaas/Content/Block/Tasks/schedulingvolumebackups.htm#UserDefinedBackupPolicies).
//  Avoid entering confidential information.
func (client BlockstorageClient) UpdateVolumeBackupPolicy(ctx context.Context, request UpdateVolumeBackupPolicyRequest) (response UpdateVolumeBackupPolicyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.updateVolumeBackupPolicy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateVolumeBackupPolicyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateVolumeBackupPolicyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateVolumeBackupPolicyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateVolumeBackupPolicyResponse")
	}
	return
}

// updateVolumeBackupPolicy implements the OCIOperation interface (enables retrying operations)
func (client BlockstorageClient) updateVolumeBackupPolicy(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/volumeBackupPolicies/{policyId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateVolumeBackupPolicyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateVolumeGroup Updates the set of volumes in a volume group along with the display name. Use this operation
// to add or remove volumes in a volume group. Specify the full list of volume IDs to include in the
// volume group. If the volume ID is not specified in the call, it will be removed from the volume group.
// Avoid entering confidential information.
// For more information, see Volume Groups (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/volumegroups.htm).
func (client BlockstorageClient) UpdateVolumeGroup(ctx context.Context, request UpdateVolumeGroupRequest) (response UpdateVolumeGroupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateVolumeGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateVolumeGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateVolumeGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateVolumeGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateVolumeGroupResponse")
	}
	return
}

// updateVolumeGroup implements the OCIOperation interface (enables retrying operations)
func (client BlockstorageClient) updateVolumeGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/volumeGroups/{volumeGroupId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateVolumeGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateVolumeGroupBackup Updates the display name for the specified volume group backup. For more information, see Volume Groups (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/volumegroups.htm).
func (client BlockstorageClient) UpdateVolumeGroupBackup(ctx context.Context, request UpdateVolumeGroupBackupRequest) (response UpdateVolumeGroupBackupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateVolumeGroupBackup, policy)
	if err != nil {
		if ociResponse != nil {
			response = UpdateVolumeGroupBackupResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateVolumeGroupBackupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateVolumeGroupBackupResponse")
	}
	return
}

// updateVolumeGroupBackup implements the OCIOperation interface (enables retrying operations)
func (client BlockstorageClient) updateVolumeGroupBackup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/volumeGroupBackups/{volumeGroupBackupId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateVolumeGroupBackupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateVolumeKmsKey Updates the specified volume with a new Key Management master encryption key.
func (client BlockstorageClient) UpdateVolumeKmsKey(ctx context.Context, request UpdateVolumeKmsKeyRequest) (response UpdateVolumeKmsKeyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateVolumeKmsKey, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateVolumeKmsKeyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateVolumeKmsKeyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateVolumeKmsKeyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateVolumeKmsKeyResponse")
	}
	return
}

// updateVolumeKmsKey implements the OCIOperation interface (enables retrying operations)
func (client BlockstorageClient) updateVolumeKmsKey(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/volumes/{volumeId}/kmsKey", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateVolumeKmsKeyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
