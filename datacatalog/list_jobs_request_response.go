// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datacatalog

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListJobsRequest wrapper for the ListJobs operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/ListJobs.go.html to see an example of how to use ListJobsRequest.
type ListJobsRequest struct {

	// Unique catalog identifier.
	CatalogId *string `mandatory:"true" contributesTo:"path" name:"catalogId"`

	// A filter to return only resources that match the entire display name given. The match is not case sensitive.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only resources that match display name pattern given. The match is not case sensitive.
	// For Example : /folders?displayNameContains=Cu.*
	// The above would match all folders with display name that starts with "Cu" or has the pattern "Cu" anywhere in between.
	DisplayNameContains *string `mandatory:"false" contributesTo:"query" name:"displayNameContains"`

	// Job lifecycle state.
	LifecycleState ListJobsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Time that the resource was created. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeCreated"`

	// Time that the resource was updated. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeUpdated"`

	// OCID of the user who created the resource.
	CreatedById *string `mandatory:"false" contributesTo:"query" name:"createdById"`

	// OCID of the user who updated the resource.
	UpdatedById *string `mandatory:"false" contributesTo:"query" name:"updatedById"`

	// Job type.
	JobType ListJobsJobTypeEnum `mandatory:"false" contributesTo:"query" name:"jobType" omitEmpty:"true"`

	// Unique job definition key.
	JobDefinitionKey *string `mandatory:"false" contributesTo:"query" name:"jobDefinitionKey"`

	// Unique data asset key.
	DataAssetKey *string `mandatory:"false" contributesTo:"query" name:"dataAssetKey"`

	// Schedule specified in the cron expression format that has seven fields for second, minute, hour, day-of-month, month, day-of-week, year.
	// It can also include special characters like * for all and ? for any. There are also pre-defined schedules that can be specified using
	// special strings. For example, @hourly will run the job every hour.
	ScheduleCronExpression *string `mandatory:"false" contributesTo:"query" name:"scheduleCronExpression"`

	// Date that the schedule should be operational. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeScheduleBegin *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeScheduleBegin"`

	// Date that the schedule should end from being operational. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeScheduleEnd *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeScheduleEnd"`

	// Type of the job schedule.
	ScheduleType ListJobsScheduleTypeEnum `mandatory:"false" contributesTo:"query" name:"scheduleType" omitEmpty:"true"`

	// Unique connection key.
	ConnectionKey *string `mandatory:"false" contributesTo:"query" name:"connectionKey"`

	// Specifies the fields to return in a job summary response.
	Fields []ListJobsFieldsEnum `contributesTo:"query" name:"fields" omitEmpty:"true" collectionFormat:"multi"`

	// The total number of executions for this job schedule.
	ExecutionCount *int `mandatory:"false" contributesTo:"query" name:"executionCount"`

	// The date and time the most recent execution for this job ,in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2019-03-25T21:10:29.600Z`
	TimeOfLatestExecution *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeOfLatestExecution"`

	// The field to sort by. Only one sort order may be provided. Default order for TIMECREATED is descending. Default order for DISPLAYNAME is ascending. If no value is specified TIMECREATED is default.
	SortBy ListJobsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListJobsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListJobsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListJobsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListJobsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListJobsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListJobsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListJobsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListJobsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListJobsJobTypeEnum(string(request.JobType)); !ok && request.JobType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for JobType: %s. Supported values are: %s.", request.JobType, strings.Join(GetListJobsJobTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListJobsScheduleTypeEnum(string(request.ScheduleType)); !ok && request.ScheduleType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ScheduleType: %s. Supported values are: %s.", request.ScheduleType, strings.Join(GetListJobsScheduleTypeEnumStringValues(), ",")))
	}
	for _, val := range request.Fields {
		if _, ok := GetMappingListJobsFieldsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Fields: %s. Supported values are: %s.", val, strings.Join(GetListJobsFieldsEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingListJobsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListJobsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListJobsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListJobsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListJobsResponse wrapper for the ListJobs operation
type ListJobsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of JobCollection instances
	JobCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListJobsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListJobsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListJobsLifecycleStateEnum Enum with underlying type: string
type ListJobsLifecycleStateEnum string

// Set of constants representing the allowable values for ListJobsLifecycleStateEnum
const (
	ListJobsLifecycleStateActive   ListJobsLifecycleStateEnum = "ACTIVE"
	ListJobsLifecycleStateInactive ListJobsLifecycleStateEnum = "INACTIVE"
	ListJobsLifecycleStateExpired  ListJobsLifecycleStateEnum = "EXPIRED"
)

var mappingListJobsLifecycleStateEnum = map[string]ListJobsLifecycleStateEnum{
	"ACTIVE":   ListJobsLifecycleStateActive,
	"INACTIVE": ListJobsLifecycleStateInactive,
	"EXPIRED":  ListJobsLifecycleStateExpired,
}

// GetListJobsLifecycleStateEnumValues Enumerates the set of values for ListJobsLifecycleStateEnum
func GetListJobsLifecycleStateEnumValues() []ListJobsLifecycleStateEnum {
	values := make([]ListJobsLifecycleStateEnum, 0)
	for _, v := range mappingListJobsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListJobsLifecycleStateEnumStringValues Enumerates the set of values in String for ListJobsLifecycleStateEnum
func GetListJobsLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
		"EXPIRED",
	}
}

// GetMappingListJobsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListJobsLifecycleStateEnum(val string) (ListJobsLifecycleStateEnum, bool) {
	mappingListJobsLifecycleStateEnumIgnoreCase := make(map[string]ListJobsLifecycleStateEnum)
	for k, v := range mappingListJobsLifecycleStateEnum {
		mappingListJobsLifecycleStateEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListJobsLifecycleStateEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListJobsJobTypeEnum Enum with underlying type: string
type ListJobsJobTypeEnum string

// Set of constants representing the allowable values for ListJobsJobTypeEnum
const (
	ListJobsJobTypeHarvest                    ListJobsJobTypeEnum = "HARVEST"
	ListJobsJobTypeProfiling                  ListJobsJobTypeEnum = "PROFILING"
	ListJobsJobTypeSampling                   ListJobsJobTypeEnum = "SAMPLING"
	ListJobsJobTypePreview                    ListJobsJobTypeEnum = "PREVIEW"
	ListJobsJobTypeImport                     ListJobsJobTypeEnum = "IMPORT"
	ListJobsJobTypeExport                     ListJobsJobTypeEnum = "EXPORT"
	ListJobsJobTypeImportGlossary             ListJobsJobTypeEnum = "IMPORT_GLOSSARY"
	ListJobsJobTypeExportGlossary             ListJobsJobTypeEnum = "EXPORT_GLOSSARY"
	ListJobsJobTypeInternal                   ListJobsJobTypeEnum = "INTERNAL"
	ListJobsJobTypePurge                      ListJobsJobTypeEnum = "PURGE"
	ListJobsJobTypeImmediate                  ListJobsJobTypeEnum = "IMMEDIATE"
	ListJobsJobTypeScheduled                  ListJobsJobTypeEnum = "SCHEDULED"
	ListJobsJobTypeImmediateExecution         ListJobsJobTypeEnum = "IMMEDIATE_EXECUTION"
	ListJobsJobTypeScheduledExecution         ListJobsJobTypeEnum = "SCHEDULED_EXECUTION"
	ListJobsJobTypeScheduledExecutionInstance ListJobsJobTypeEnum = "SCHEDULED_EXECUTION_INSTANCE"
	ListJobsJobTypeAsyncDelete                ListJobsJobTypeEnum = "ASYNC_DELETE"
	ListJobsJobTypeImportDataAsset            ListJobsJobTypeEnum = "IMPORT_DATA_ASSET"
)

var mappingListJobsJobTypeEnum = map[string]ListJobsJobTypeEnum{
	"HARVEST":                      ListJobsJobTypeHarvest,
	"PROFILING":                    ListJobsJobTypeProfiling,
	"SAMPLING":                     ListJobsJobTypeSampling,
	"PREVIEW":                      ListJobsJobTypePreview,
	"IMPORT":                       ListJobsJobTypeImport,
	"EXPORT":                       ListJobsJobTypeExport,
	"IMPORT_GLOSSARY":              ListJobsJobTypeImportGlossary,
	"EXPORT_GLOSSARY":              ListJobsJobTypeExportGlossary,
	"INTERNAL":                     ListJobsJobTypeInternal,
	"PURGE":                        ListJobsJobTypePurge,
	"IMMEDIATE":                    ListJobsJobTypeImmediate,
	"SCHEDULED":                    ListJobsJobTypeScheduled,
	"IMMEDIATE_EXECUTION":          ListJobsJobTypeImmediateExecution,
	"SCHEDULED_EXECUTION":          ListJobsJobTypeScheduledExecution,
	"SCHEDULED_EXECUTION_INSTANCE": ListJobsJobTypeScheduledExecutionInstance,
	"ASYNC_DELETE":                 ListJobsJobTypeAsyncDelete,
	"IMPORT_DATA_ASSET":            ListJobsJobTypeImportDataAsset,
}

// GetListJobsJobTypeEnumValues Enumerates the set of values for ListJobsJobTypeEnum
func GetListJobsJobTypeEnumValues() []ListJobsJobTypeEnum {
	values := make([]ListJobsJobTypeEnum, 0)
	for _, v := range mappingListJobsJobTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListJobsJobTypeEnumStringValues Enumerates the set of values in String for ListJobsJobTypeEnum
func GetListJobsJobTypeEnumStringValues() []string {
	return []string{
		"HARVEST",
		"PROFILING",
		"SAMPLING",
		"PREVIEW",
		"IMPORT",
		"EXPORT",
		"IMPORT_GLOSSARY",
		"EXPORT_GLOSSARY",
		"INTERNAL",
		"PURGE",
		"IMMEDIATE",
		"SCHEDULED",
		"IMMEDIATE_EXECUTION",
		"SCHEDULED_EXECUTION",
		"SCHEDULED_EXECUTION_INSTANCE",
		"ASYNC_DELETE",
		"IMPORT_DATA_ASSET",
	}
}

// GetMappingListJobsJobTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListJobsJobTypeEnum(val string) (ListJobsJobTypeEnum, bool) {
	mappingListJobsJobTypeEnumIgnoreCase := make(map[string]ListJobsJobTypeEnum)
	for k, v := range mappingListJobsJobTypeEnum {
		mappingListJobsJobTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListJobsJobTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListJobsScheduleTypeEnum Enum with underlying type: string
type ListJobsScheduleTypeEnum string

// Set of constants representing the allowable values for ListJobsScheduleTypeEnum
const (
	ListJobsScheduleTypeScheduled ListJobsScheduleTypeEnum = "SCHEDULED"
	ListJobsScheduleTypeImmediate ListJobsScheduleTypeEnum = "IMMEDIATE"
)

var mappingListJobsScheduleTypeEnum = map[string]ListJobsScheduleTypeEnum{
	"SCHEDULED": ListJobsScheduleTypeScheduled,
	"IMMEDIATE": ListJobsScheduleTypeImmediate,
}

// GetListJobsScheduleTypeEnumValues Enumerates the set of values for ListJobsScheduleTypeEnum
func GetListJobsScheduleTypeEnumValues() []ListJobsScheduleTypeEnum {
	values := make([]ListJobsScheduleTypeEnum, 0)
	for _, v := range mappingListJobsScheduleTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListJobsScheduleTypeEnumStringValues Enumerates the set of values in String for ListJobsScheduleTypeEnum
func GetListJobsScheduleTypeEnumStringValues() []string {
	return []string{
		"SCHEDULED",
		"IMMEDIATE",
	}
}

// GetMappingListJobsScheduleTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListJobsScheduleTypeEnum(val string) (ListJobsScheduleTypeEnum, bool) {
	mappingListJobsScheduleTypeEnumIgnoreCase := make(map[string]ListJobsScheduleTypeEnum)
	for k, v := range mappingListJobsScheduleTypeEnum {
		mappingListJobsScheduleTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListJobsScheduleTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListJobsFieldsEnum Enum with underlying type: string
type ListJobsFieldsEnum string

// Set of constants representing the allowable values for ListJobsFieldsEnum
const (
	ListJobsFieldsKey                    ListJobsFieldsEnum = "key"
	ListJobsFieldsDisplayname            ListJobsFieldsEnum = "displayName"
	ListJobsFieldsDescription            ListJobsFieldsEnum = "description"
	ListJobsFieldsCatalogid              ListJobsFieldsEnum = "catalogId"
	ListJobsFieldsJobdefinitionkey       ListJobsFieldsEnum = "jobDefinitionKey"
	ListJobsFieldsLifecyclestate         ListJobsFieldsEnum = "lifecycleState"
	ListJobsFieldsTimecreated            ListJobsFieldsEnum = "timeCreated"
	ListJobsFieldsTimeupdated            ListJobsFieldsEnum = "timeUpdated"
	ListJobsFieldsCreatedbyid            ListJobsFieldsEnum = "createdById"
	ListJobsFieldsUpdatedbyid            ListJobsFieldsEnum = "updatedById"
	ListJobsFieldsJobtype                ListJobsFieldsEnum = "jobType"
	ListJobsFieldsSchedulecronexpression ListJobsFieldsEnum = "scheduleCronExpression"
	ListJobsFieldsTimeschedulebegin      ListJobsFieldsEnum = "timeScheduleBegin"
	ListJobsFieldsScheduletype           ListJobsFieldsEnum = "scheduleType"
	ListJobsFieldsExecutioncount         ListJobsFieldsEnum = "executionCount"
	ListJobsFieldsTimeoflatestexecution  ListJobsFieldsEnum = "timeOfLatestExecution"
	ListJobsFieldsExecutions             ListJobsFieldsEnum = "executions"
	ListJobsFieldsUri                    ListJobsFieldsEnum = "uri"
	ListJobsFieldsJobdefinitionname      ListJobsFieldsEnum = "jobDefinitionName"
	ListJobsFieldsErrorcode              ListJobsFieldsEnum = "errorCode"
	ListJobsFieldsErrormessage           ListJobsFieldsEnum = "errorMessage"
)

var mappingListJobsFieldsEnum = map[string]ListJobsFieldsEnum{
	"key":                    ListJobsFieldsKey,
	"displayName":            ListJobsFieldsDisplayname,
	"description":            ListJobsFieldsDescription,
	"catalogId":              ListJobsFieldsCatalogid,
	"jobDefinitionKey":       ListJobsFieldsJobdefinitionkey,
	"lifecycleState":         ListJobsFieldsLifecyclestate,
	"timeCreated":            ListJobsFieldsTimecreated,
	"timeUpdated":            ListJobsFieldsTimeupdated,
	"createdById":            ListJobsFieldsCreatedbyid,
	"updatedById":            ListJobsFieldsUpdatedbyid,
	"jobType":                ListJobsFieldsJobtype,
	"scheduleCronExpression": ListJobsFieldsSchedulecronexpression,
	"timeScheduleBegin":      ListJobsFieldsTimeschedulebegin,
	"scheduleType":           ListJobsFieldsScheduletype,
	"executionCount":         ListJobsFieldsExecutioncount,
	"timeOfLatestExecution":  ListJobsFieldsTimeoflatestexecution,
	"executions":             ListJobsFieldsExecutions,
	"uri":                    ListJobsFieldsUri,
	"jobDefinitionName":      ListJobsFieldsJobdefinitionname,
	"errorCode":              ListJobsFieldsErrorcode,
	"errorMessage":           ListJobsFieldsErrormessage,
}

// GetListJobsFieldsEnumValues Enumerates the set of values for ListJobsFieldsEnum
func GetListJobsFieldsEnumValues() []ListJobsFieldsEnum {
	values := make([]ListJobsFieldsEnum, 0)
	for _, v := range mappingListJobsFieldsEnum {
		values = append(values, v)
	}
	return values
}

// GetListJobsFieldsEnumStringValues Enumerates the set of values in String for ListJobsFieldsEnum
func GetListJobsFieldsEnumStringValues() []string {
	return []string{
		"key",
		"displayName",
		"description",
		"catalogId",
		"jobDefinitionKey",
		"lifecycleState",
		"timeCreated",
		"timeUpdated",
		"createdById",
		"updatedById",
		"jobType",
		"scheduleCronExpression",
		"timeScheduleBegin",
		"scheduleType",
		"executionCount",
		"timeOfLatestExecution",
		"executions",
		"uri",
		"jobDefinitionName",
		"errorCode",
		"errorMessage",
	}
}

// GetMappingListJobsFieldsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListJobsFieldsEnum(val string) (ListJobsFieldsEnum, bool) {
	mappingListJobsFieldsEnumIgnoreCase := make(map[string]ListJobsFieldsEnum)
	for k, v := range mappingListJobsFieldsEnum {
		mappingListJobsFieldsEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListJobsFieldsEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListJobsSortByEnum Enum with underlying type: string
type ListJobsSortByEnum string

// Set of constants representing the allowable values for ListJobsSortByEnum
const (
	ListJobsSortByTimecreated ListJobsSortByEnum = "TIMECREATED"
	ListJobsSortByDisplayname ListJobsSortByEnum = "DISPLAYNAME"
)

var mappingListJobsSortByEnum = map[string]ListJobsSortByEnum{
	"TIMECREATED": ListJobsSortByTimecreated,
	"DISPLAYNAME": ListJobsSortByDisplayname,
}

// GetListJobsSortByEnumValues Enumerates the set of values for ListJobsSortByEnum
func GetListJobsSortByEnumValues() []ListJobsSortByEnum {
	values := make([]ListJobsSortByEnum, 0)
	for _, v := range mappingListJobsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListJobsSortByEnumStringValues Enumerates the set of values in String for ListJobsSortByEnum
func GetListJobsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListJobsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListJobsSortByEnum(val string) (ListJobsSortByEnum, bool) {
	mappingListJobsSortByEnumIgnoreCase := make(map[string]ListJobsSortByEnum)
	for k, v := range mappingListJobsSortByEnum {
		mappingListJobsSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListJobsSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListJobsSortOrderEnum Enum with underlying type: string
type ListJobsSortOrderEnum string

// Set of constants representing the allowable values for ListJobsSortOrderEnum
const (
	ListJobsSortOrderAsc  ListJobsSortOrderEnum = "ASC"
	ListJobsSortOrderDesc ListJobsSortOrderEnum = "DESC"
)

var mappingListJobsSortOrderEnum = map[string]ListJobsSortOrderEnum{
	"ASC":  ListJobsSortOrderAsc,
	"DESC": ListJobsSortOrderDesc,
}

// GetListJobsSortOrderEnumValues Enumerates the set of values for ListJobsSortOrderEnum
func GetListJobsSortOrderEnumValues() []ListJobsSortOrderEnum {
	values := make([]ListJobsSortOrderEnum, 0)
	for _, v := range mappingListJobsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListJobsSortOrderEnumStringValues Enumerates the set of values in String for ListJobsSortOrderEnum
func GetListJobsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListJobsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListJobsSortOrderEnum(val string) (ListJobsSortOrderEnum, bool) {
	mappingListJobsSortOrderEnumIgnoreCase := make(map[string]ListJobsSortOrderEnum)
	for k, v := range mappingListJobsSortOrderEnum {
		mappingListJobsSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListJobsSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
