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

// AuthToken An `AuthToken` is an Oracle-generated token string that you can use to authenticate with third-party APIs
// that do not support Oracle Cloud Infrastructure's signature-based authentication. For example, use an `AuthToken`
// to authenticate with a Swift client with the Object Storage Service.
// The auth token is associated with the user's Console login. Auth tokens never expire. A user can have up to two
// auth tokens at a time.
// **Note:** The token is always an Oracle-generated string; you can't change it to a string of your choice.
// For more information, see Managing User Credentials (https://docs.cloud.oracle.com/Content/Identity/Tasks/managingcredentials.htm).
type AuthToken struct {

	// The auth token. The value is available only in the response for `CreateAuthToken`, and not
	// for `ListAuthTokens` or `UpdateAuthToken`.
	Token *string `mandatory:"false" json:"token"`

	// The OCID of the auth token.
	Id *string `mandatory:"false" json:"id"`

	// The OCID of the user the auth token belongs to.
	UserId *string `mandatory:"false" json:"userId"`

	// The description you assign to the auth token. Does not have to be unique, and it's changeable.
	Description *string `mandatory:"false" json:"description"`

	// Date and time the `AuthToken` object was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Date and time when this auth token will expire, in the format defined by RFC3339.
	// Null if it never expires.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeExpires *common.SDKTime `mandatory:"false" json:"timeExpires"`

	// The token's current state. After creating an auth token, make sure its `lifecycleState` changes from
	// CREATING to ACTIVE before using it.
	LifecycleState AuthTokenLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The detailed status of INACTIVE lifecycleState.
	InactiveStatus *int64 `mandatory:"false" json:"inactiveStatus"`
}

func (m AuthToken) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AuthToken) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingAuthTokenLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetAuthTokenLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AuthTokenLifecycleStateEnum Enum with underlying type: string
type AuthTokenLifecycleStateEnum string

// Set of constants representing the allowable values for AuthTokenLifecycleStateEnum
const (
	AuthTokenLifecycleStateCreating AuthTokenLifecycleStateEnum = "CREATING"
	AuthTokenLifecycleStateActive   AuthTokenLifecycleStateEnum = "ACTIVE"
	AuthTokenLifecycleStateInactive AuthTokenLifecycleStateEnum = "INACTIVE"
	AuthTokenLifecycleStateDeleting AuthTokenLifecycleStateEnum = "DELETING"
	AuthTokenLifecycleStateDeleted  AuthTokenLifecycleStateEnum = "DELETED"
)

var mappingAuthTokenLifecycleStateEnum = map[string]AuthTokenLifecycleStateEnum{
	"CREATING": AuthTokenLifecycleStateCreating,
	"ACTIVE":   AuthTokenLifecycleStateActive,
	"INACTIVE": AuthTokenLifecycleStateInactive,
	"DELETING": AuthTokenLifecycleStateDeleting,
	"DELETED":  AuthTokenLifecycleStateDeleted,
}

// GetAuthTokenLifecycleStateEnumValues Enumerates the set of values for AuthTokenLifecycleStateEnum
func GetAuthTokenLifecycleStateEnumValues() []AuthTokenLifecycleStateEnum {
	values := make([]AuthTokenLifecycleStateEnum, 0)
	for _, v := range mappingAuthTokenLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetAuthTokenLifecycleStateEnumStringValues Enumerates the set of values in String for AuthTokenLifecycleStateEnum
func GetAuthTokenLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"INACTIVE",
		"DELETING",
		"DELETED",
	}
}

// GetMappingAuthTokenLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAuthTokenLifecycleStateEnum(val string) (AuthTokenLifecycleStateEnum, bool) {
	mappingAuthTokenLifecycleStateEnumIgnoreCase := make(map[string]AuthTokenLifecycleStateEnum)
	for k, v := range mappingAuthTokenLifecycleStateEnum {
		mappingAuthTokenLifecycleStateEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingAuthTokenLifecycleStateEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
