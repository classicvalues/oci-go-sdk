// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// TaggingWorkRequestSummary The work request summary. Tracks the status of the asynchronous operation.
type TaggingWorkRequestSummary struct {

	// The OCID of the work request.
	Id *string `mandatory:"true" json:"id"`

	// An enum-like description of the type of work the work request is doing.
	OperationType TaggingWorkRequestSummaryOperationTypeEnum `mandatory:"true" json:"operationType"`

	// The current status of the work request.
	Status TaggingWorkRequestSummaryStatusEnum `mandatory:"true" json:"status"`

	// The OCID of the compartment that contains the work request.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The resources this work request affects.
	Resources []WorkRequestResource `mandatory:"false" json:"resources"`

	// Date and time the work was accepted, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeAccepted *common.SDKTime `mandatory:"false" json:"timeAccepted"`

	// Date and time the work started, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// Date and time the work completed, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`

	// How much progress the operation has made.
	PercentComplete *float32 `mandatory:"false" json:"percentComplete"`
}

func (m TaggingWorkRequestSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TaggingWorkRequestSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingTaggingWorkRequestSummaryOperationTypeEnum(string(m.OperationType)); !ok && m.OperationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OperationType: %s. Supported values are: %s.", m.OperationType, strings.Join(GetTaggingWorkRequestSummaryOperationTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingTaggingWorkRequestSummaryStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetTaggingWorkRequestSummaryStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TaggingWorkRequestSummaryOperationTypeEnum Enum with underlying type: string
type TaggingWorkRequestSummaryOperationTypeEnum string

// Set of constants representing the allowable values for TaggingWorkRequestSummaryOperationTypeEnum
const (
	TaggingWorkRequestSummaryOperationTypeDeleteTagDefinition        TaggingWorkRequestSummaryOperationTypeEnum = "DELETE_TAG_DEFINITION"
	TaggingWorkRequestSummaryOperationTypeDeleteNonEmptyTagNamespace TaggingWorkRequestSummaryOperationTypeEnum = "DELETE_NON_EMPTY_TAG_NAMESPACE"
	TaggingWorkRequestSummaryOperationTypeBulkDeleteTagDefinition    TaggingWorkRequestSummaryOperationTypeEnum = "BULK_DELETE_TAG_DEFINITION"
	TaggingWorkRequestSummaryOperationTypeBulkEditOfTags             TaggingWorkRequestSummaryOperationTypeEnum = "BULK_EDIT_OF_TAGS"
	TaggingWorkRequestSummaryOperationTypeImportStandardTags         TaggingWorkRequestSummaryOperationTypeEnum = "IMPORT_STANDARD_TAGS"
)

var mappingTaggingWorkRequestSummaryOperationTypeEnum = map[string]TaggingWorkRequestSummaryOperationTypeEnum{
	"DELETE_TAG_DEFINITION":          TaggingWorkRequestSummaryOperationTypeDeleteTagDefinition,
	"DELETE_NON_EMPTY_TAG_NAMESPACE": TaggingWorkRequestSummaryOperationTypeDeleteNonEmptyTagNamespace,
	"BULK_DELETE_TAG_DEFINITION":     TaggingWorkRequestSummaryOperationTypeBulkDeleteTagDefinition,
	"BULK_EDIT_OF_TAGS":              TaggingWorkRequestSummaryOperationTypeBulkEditOfTags,
	"IMPORT_STANDARD_TAGS":           TaggingWorkRequestSummaryOperationTypeImportStandardTags,
}

// GetTaggingWorkRequestSummaryOperationTypeEnumValues Enumerates the set of values for TaggingWorkRequestSummaryOperationTypeEnum
func GetTaggingWorkRequestSummaryOperationTypeEnumValues() []TaggingWorkRequestSummaryOperationTypeEnum {
	values := make([]TaggingWorkRequestSummaryOperationTypeEnum, 0)
	for _, v := range mappingTaggingWorkRequestSummaryOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetTaggingWorkRequestSummaryOperationTypeEnumStringValues Enumerates the set of values in String for TaggingWorkRequestSummaryOperationTypeEnum
func GetTaggingWorkRequestSummaryOperationTypeEnumStringValues() []string {
	return []string{
		"DELETE_TAG_DEFINITION",
		"DELETE_NON_EMPTY_TAG_NAMESPACE",
		"BULK_DELETE_TAG_DEFINITION",
		"BULK_EDIT_OF_TAGS",
		"IMPORT_STANDARD_TAGS",
	}
}

// GetMappingTaggingWorkRequestSummaryOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTaggingWorkRequestSummaryOperationTypeEnum(val string) (TaggingWorkRequestSummaryOperationTypeEnum, bool) {
	mappingTaggingWorkRequestSummaryOperationTypeEnumIgnoreCase := make(map[string]TaggingWorkRequestSummaryOperationTypeEnum)
	for k, v := range mappingTaggingWorkRequestSummaryOperationTypeEnum {
		mappingTaggingWorkRequestSummaryOperationTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingTaggingWorkRequestSummaryOperationTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// TaggingWorkRequestSummaryStatusEnum Enum with underlying type: string
type TaggingWorkRequestSummaryStatusEnum string

// Set of constants representing the allowable values for TaggingWorkRequestSummaryStatusEnum
const (
	TaggingWorkRequestSummaryStatusAccepted           TaggingWorkRequestSummaryStatusEnum = "ACCEPTED"
	TaggingWorkRequestSummaryStatusInProgress         TaggingWorkRequestSummaryStatusEnum = "IN_PROGRESS"
	TaggingWorkRequestSummaryStatusFailed             TaggingWorkRequestSummaryStatusEnum = "FAILED"
	TaggingWorkRequestSummaryStatusSucceeded          TaggingWorkRequestSummaryStatusEnum = "SUCCEEDED"
	TaggingWorkRequestSummaryStatusPartiallySucceeded TaggingWorkRequestSummaryStatusEnum = "PARTIALLY_SUCCEEDED"
	TaggingWorkRequestSummaryStatusCanceling          TaggingWorkRequestSummaryStatusEnum = "CANCELING"
	TaggingWorkRequestSummaryStatusCanceled           TaggingWorkRequestSummaryStatusEnum = "CANCELED"
)

var mappingTaggingWorkRequestSummaryStatusEnum = map[string]TaggingWorkRequestSummaryStatusEnum{
	"ACCEPTED":            TaggingWorkRequestSummaryStatusAccepted,
	"IN_PROGRESS":         TaggingWorkRequestSummaryStatusInProgress,
	"FAILED":              TaggingWorkRequestSummaryStatusFailed,
	"SUCCEEDED":           TaggingWorkRequestSummaryStatusSucceeded,
	"PARTIALLY_SUCCEEDED": TaggingWorkRequestSummaryStatusPartiallySucceeded,
	"CANCELING":           TaggingWorkRequestSummaryStatusCanceling,
	"CANCELED":            TaggingWorkRequestSummaryStatusCanceled,
}

// GetTaggingWorkRequestSummaryStatusEnumValues Enumerates the set of values for TaggingWorkRequestSummaryStatusEnum
func GetTaggingWorkRequestSummaryStatusEnumValues() []TaggingWorkRequestSummaryStatusEnum {
	values := make([]TaggingWorkRequestSummaryStatusEnum, 0)
	for _, v := range mappingTaggingWorkRequestSummaryStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetTaggingWorkRequestSummaryStatusEnumStringValues Enumerates the set of values in String for TaggingWorkRequestSummaryStatusEnum
func GetTaggingWorkRequestSummaryStatusEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"IN_PROGRESS",
		"FAILED",
		"SUCCEEDED",
		"PARTIALLY_SUCCEEDED",
		"CANCELING",
		"CANCELED",
	}
}

// GetMappingTaggingWorkRequestSummaryStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTaggingWorkRequestSummaryStatusEnum(val string) (TaggingWorkRequestSummaryStatusEnum, bool) {
	mappingTaggingWorkRequestSummaryStatusEnumIgnoreCase := make(map[string]TaggingWorkRequestSummaryStatusEnum)
	for k, v := range mappingTaggingWorkRequestSummaryStatusEnum {
		mappingTaggingWorkRequestSummaryStatusEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingTaggingWorkRequestSummaryStatusEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
