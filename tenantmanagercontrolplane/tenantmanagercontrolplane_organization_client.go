// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Organizations API
//
// The Organizations API allows you to consolidate multiple OCI tenancies into an organization, and centrally manage your tenancies and its resources.
//

package tenantmanagercontrolplane

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"github.com/oracle/oci-go-sdk/v58/common/auth"
	"net/http"
)

//OrganizationClient a client for Organization
type OrganizationClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewOrganizationClientWithConfigurationProvider Creates a new default Organization client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewOrganizationClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client OrganizationClient, err error) {
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newOrganizationClientFromBaseClient(baseClient, provider)
}

// NewOrganizationClientWithOboToken Creates a new default Organization client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewOrganizationClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client OrganizationClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newOrganizationClientFromBaseClient(baseClient, configProvider)
}

func newOrganizationClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client OrganizationClient, err error) {
	// Organization service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSetting())
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = OrganizationClient{BaseClient: baseClient}
	client.BasePath = "20200801"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *OrganizationClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("tenantmanagercontrolplane", "https://organizations.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *OrganizationClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *OrganizationClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ApproveOrganizationTenancyForTransfer Approve an organization's child tenancy for transfer.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/tenantmanagercontrolplane/ApproveOrganizationTenancyForTransfer.go.html to see an example of how to use ApproveOrganizationTenancyForTransfer API.
func (client OrganizationClient) ApproveOrganizationTenancyForTransfer(ctx context.Context, request ApproveOrganizationTenancyForTransferRequest) (response ApproveOrganizationTenancyForTransferResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.approveOrganizationTenancyForTransfer, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ApproveOrganizationTenancyForTransferResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ApproveOrganizationTenancyForTransferResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ApproveOrganizationTenancyForTransferResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ApproveOrganizationTenancyForTransferResponse")
	}
	return
}

// approveOrganizationTenancyForTransfer implements the OCIOperation interface (enables retrying operations)
func (client OrganizationClient) approveOrganizationTenancyForTransfer(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/organizationTenancies/{organizationTenancyId}/actions/approveForTransfer", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ApproveOrganizationTenancyForTransferResponse
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

// CreateChildTenancy Creates a child tenancy asynchronously.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/tenantmanagercontrolplane/CreateChildTenancy.go.html to see an example of how to use CreateChildTenancy API.
func (client OrganizationClient) CreateChildTenancy(ctx context.Context, request CreateChildTenancyRequest) (response CreateChildTenancyResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createChildTenancy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateChildTenancyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateChildTenancyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateChildTenancyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateChildTenancyResponse")
	}
	return
}

// createChildTenancy implements the OCIOperation interface (enables retrying operations)
func (client OrganizationClient) createChildTenancy(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/childTenancies", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateChildTenancyResponse
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

// DeleteOrganizationTenancy If certain validations are successful, initiate tenancy termination.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/tenantmanagercontrolplane/DeleteOrganizationTenancy.go.html to see an example of how to use DeleteOrganizationTenancy API.
func (client OrganizationClient) DeleteOrganizationTenancy(ctx context.Context, request DeleteOrganizationTenancyRequest) (response DeleteOrganizationTenancyResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.deleteOrganizationTenancy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteOrganizationTenancyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteOrganizationTenancyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteOrganizationTenancyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteOrganizationTenancyResponse")
	}
	return
}

// deleteOrganizationTenancy implements the OCIOperation interface (enables retrying operations)
func (client OrganizationClient) deleteOrganizationTenancy(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/organizationTenancies/{organizationTenancyId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteOrganizationTenancyResponse
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

// GetOrganization Gets information about the organization.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/tenantmanagercontrolplane/GetOrganization.go.html to see an example of how to use GetOrganization API.
func (client OrganizationClient) GetOrganization(ctx context.Context, request GetOrganizationRequest) (response GetOrganizationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getOrganization, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetOrganizationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetOrganizationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetOrganizationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetOrganizationResponse")
	}
	return
}

// getOrganization implements the OCIOperation interface (enables retrying operations)
func (client OrganizationClient) getOrganization(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/organizations/{organizationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetOrganizationResponse
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

// GetOrganizationTenancy Gets information about the organization's tenancy.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/tenantmanagercontrolplane/GetOrganizationTenancy.go.html to see an example of how to use GetOrganizationTenancy API.
func (client OrganizationClient) GetOrganizationTenancy(ctx context.Context, request GetOrganizationTenancyRequest) (response GetOrganizationTenancyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getOrganizationTenancy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetOrganizationTenancyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetOrganizationTenancyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetOrganizationTenancyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetOrganizationTenancyResponse")
	}
	return
}

// getOrganizationTenancy implements the OCIOperation interface (enables retrying operations)
func (client OrganizationClient) getOrganizationTenancy(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/organizations/{organizationId}/tenancies/{tenancyId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetOrganizationTenancyResponse
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

// ListOrganizationTenancies Gets a list of tenancies in the organization.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/tenantmanagercontrolplane/ListOrganizationTenancies.go.html to see an example of how to use ListOrganizationTenancies API.
func (client OrganizationClient) ListOrganizationTenancies(ctx context.Context, request ListOrganizationTenanciesRequest) (response ListOrganizationTenanciesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listOrganizationTenancies, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListOrganizationTenanciesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListOrganizationTenanciesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListOrganizationTenanciesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListOrganizationTenanciesResponse")
	}
	return
}

// listOrganizationTenancies implements the OCIOperation interface (enables retrying operations)
func (client OrganizationClient) listOrganizationTenancies(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/organizations/{organizationId}/tenancies", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListOrganizationTenanciesResponse
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

// ListOrganizations Lists organizations associated with the caller.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/tenantmanagercontrolplane/ListOrganizations.go.html to see an example of how to use ListOrganizations API.
func (client OrganizationClient) ListOrganizations(ctx context.Context, request ListOrganizationsRequest) (response ListOrganizationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listOrganizations, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListOrganizationsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListOrganizationsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListOrganizationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListOrganizationsResponse")
	}
	return
}

// listOrganizations implements the OCIOperation interface (enables retrying operations)
func (client OrganizationClient) listOrganizations(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/organizations", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListOrganizationsResponse
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

// UnapproveOrganizationTenancyForTransfer Cancel an organization's child tenancy for transfer.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/tenantmanagercontrolplane/UnapproveOrganizationTenancyForTransfer.go.html to see an example of how to use UnapproveOrganizationTenancyForTransfer API.
func (client OrganizationClient) UnapproveOrganizationTenancyForTransfer(ctx context.Context, request UnapproveOrganizationTenancyForTransferRequest) (response UnapproveOrganizationTenancyForTransferResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.unapproveOrganizationTenancyForTransfer, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UnapproveOrganizationTenancyForTransferResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UnapproveOrganizationTenancyForTransferResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UnapproveOrganizationTenancyForTransferResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UnapproveOrganizationTenancyForTransferResponse")
	}
	return
}

// unapproveOrganizationTenancyForTransfer implements the OCIOperation interface (enables retrying operations)
func (client OrganizationClient) unapproveOrganizationTenancyForTransfer(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/organizationTenancies/{organizationTenancyId}/actions/unapproveForTransfer", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UnapproveOrganizationTenancyForTransferResponse
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

// UpdateOrganization Assign the default subscription to the organization.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/tenantmanagercontrolplane/UpdateOrganization.go.html to see an example of how to use UpdateOrganization API.
func (client OrganizationClient) UpdateOrganization(ctx context.Context, request UpdateOrganizationRequest) (response UpdateOrganizationResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.updateOrganization, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateOrganizationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateOrganizationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateOrganizationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateOrganizationResponse")
	}
	return
}

// updateOrganization implements the OCIOperation interface (enables retrying operations)
func (client OrganizationClient) updateOrganization(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/organizations/{organizationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateOrganizationResponse
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
