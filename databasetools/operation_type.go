// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools
//
// Database Tools APIs to manage Connections and Private Endpoints.
//

package databasetools

import (
	"strings"
)

// OperationTypeEnum Enum with underlying type: string
type OperationTypeEnum string

// Set of constants representing the allowable values for OperationTypeEnum
const (
	OperationTypeCreateDatabaseToolsConnection      OperationTypeEnum = "CREATE_DATABASE_TOOLS_CONNECTION"
	OperationTypeUpdateDatabaseToolsConnection      OperationTypeEnum = "UPDATE_DATABASE_TOOLS_CONNECTION"
	OperationTypeDeleteDatabaseToolsConnection      OperationTypeEnum = "DELETE_DATABASE_TOOLS_CONNECTION"
	OperationTypeCreateDatabaseToolsServiceInstance OperationTypeEnum = "CREATE_DATABASE_TOOLS_SERVICE_INSTANCE"
	OperationTypeUpdateDatabaseToolsServiceInstance OperationTypeEnum = "UPDATE_DATABASE_TOOLS_SERVICE_INSTANCE"
	OperationTypeDeleteDatabaseToolsServiceInstance OperationTypeEnum = "DELETE_DATABASE_TOOLS_SERVICE_INSTANCE"
	OperationTypeCreateDatabaseToolsPrivateEndpoint OperationTypeEnum = "CREATE_DATABASE_TOOLS_PRIVATE_ENDPOINT"
	OperationTypeUpdateDatabaseToolsPrivateEndpoint OperationTypeEnum = "UPDATE_DATABASE_TOOLS_PRIVATE_ENDPOINT"
	OperationTypeDeleteDatabaseToolsPrivateEndpoint OperationTypeEnum = "DELETE_DATABASE_TOOLS_PRIVATE_ENDPOINT"
)

var mappingOperationTypeEnum = map[string]OperationTypeEnum{
	"CREATE_DATABASE_TOOLS_CONNECTION":       OperationTypeCreateDatabaseToolsConnection,
	"UPDATE_DATABASE_TOOLS_CONNECTION":       OperationTypeUpdateDatabaseToolsConnection,
	"DELETE_DATABASE_TOOLS_CONNECTION":       OperationTypeDeleteDatabaseToolsConnection,
	"CREATE_DATABASE_TOOLS_SERVICE_INSTANCE": OperationTypeCreateDatabaseToolsServiceInstance,
	"UPDATE_DATABASE_TOOLS_SERVICE_INSTANCE": OperationTypeUpdateDatabaseToolsServiceInstance,
	"DELETE_DATABASE_TOOLS_SERVICE_INSTANCE": OperationTypeDeleteDatabaseToolsServiceInstance,
	"CREATE_DATABASE_TOOLS_PRIVATE_ENDPOINT": OperationTypeCreateDatabaseToolsPrivateEndpoint,
	"UPDATE_DATABASE_TOOLS_PRIVATE_ENDPOINT": OperationTypeUpdateDatabaseToolsPrivateEndpoint,
	"DELETE_DATABASE_TOOLS_PRIVATE_ENDPOINT": OperationTypeDeleteDatabaseToolsPrivateEndpoint,
}

// GetOperationTypeEnumValues Enumerates the set of values for OperationTypeEnum
func GetOperationTypeEnumValues() []OperationTypeEnum {
	values := make([]OperationTypeEnum, 0)
	for _, v := range mappingOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOperationTypeEnumStringValues Enumerates the set of values in String for OperationTypeEnum
func GetOperationTypeEnumStringValues() []string {
	return []string{
		"CREATE_DATABASE_TOOLS_CONNECTION",
		"UPDATE_DATABASE_TOOLS_CONNECTION",
		"DELETE_DATABASE_TOOLS_CONNECTION",
		"CREATE_DATABASE_TOOLS_SERVICE_INSTANCE",
		"UPDATE_DATABASE_TOOLS_SERVICE_INSTANCE",
		"DELETE_DATABASE_TOOLS_SERVICE_INSTANCE",
		"CREATE_DATABASE_TOOLS_PRIVATE_ENDPOINT",
		"UPDATE_DATABASE_TOOLS_PRIVATE_ENDPOINT",
		"DELETE_DATABASE_TOOLS_PRIVATE_ENDPOINT",
	}
}

// GetMappingOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypeEnum(val string) (OperationTypeEnum, bool) {
	mappingOperationTypeEnumIgnoreCase := make(map[string]OperationTypeEnum)
	for k, v := range mappingOperationTypeEnum {
		mappingOperationTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingOperationTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
