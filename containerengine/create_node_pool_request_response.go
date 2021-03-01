// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package containerengine

import (
	"github.com/oracle/oci-go-sdk/v36/common"
	"net/http"
)

// CreateNodePoolRequest wrapper for the CreateNodePool operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/containerengine/CreateNodePool.go.html to see an example of how to use CreateNodePoolRequest.
type CreateNodePoolRequest struct {

	// The details of the node pool to create.
	CreateNodePoolDetails `contributesTo:"body"`

	// A token you supply to uniquely identify the request and provide idempotency if
	// the request is retried. Idempotency tokens expire after 24 hours.
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request CreateNodePoolRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request CreateNodePoolRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request CreateNodePoolRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// CreateNodePoolResponse wrapper for the CreateNodePool operation
type CreateNodePoolResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The OCID of the work request handling the operation.
	OpcWorkRequestId *string `presentIn:"header" name:"opc-work-request-id"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response CreateNodePoolResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response CreateNodePoolResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
