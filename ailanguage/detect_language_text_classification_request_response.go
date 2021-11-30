// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package ailanguage

import (
	"github.com/oracle/oci-go-sdk/v53/common"
	"net/http"
)

// DetectLanguageTextClassificationRequest wrapper for the DetectLanguageTextClassification operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ailanguage/DetectLanguageTextClassification.go.html to see an example of how to use DetectLanguageTextClassificationRequest.
type DetectLanguageTextClassificationRequest struct {

	// The details to make text classification detect call.
	// Example: `{"text": "If an emerging growth company, indicate by check mark if the registrant has elected not to use the extended transition period for complying"}`
	DetectLanguageTextClassificationDetails `contributesTo:"body"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request DetectLanguageTextClassificationRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request DetectLanguageTextClassificationRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request DetectLanguageTextClassificationRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request DetectLanguageTextClassificationRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// DetectLanguageTextClassificationResponse wrapper for the DetectLanguageTextClassification operation
type DetectLanguageTextClassificationResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The DetectLanguageTextClassificationResult instance
	DetectLanguageTextClassificationResult `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response DetectLanguageTextClassificationResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response DetectLanguageTextClassificationResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
