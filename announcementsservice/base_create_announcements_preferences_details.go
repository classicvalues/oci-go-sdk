// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Announcements Service API
//
// Manage Oracle Cloud Infrastructure console announcements.
//

package announcementsservice

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// BaseCreateAnnouncementsPreferencesDetails The model for the parameters of announcement email preferences configured for the tenancy (root compartment).
type BaseCreateAnnouncementsPreferencesDetails interface {

	// The string representing the user's preference, whether to opt in to only required announcements, to opt in to all announcements, including informational announcements, or to opt out of all announcements.
	GetPreferenceType() BaseCreateAnnouncementsPreferencesDetailsPreferenceTypeEnum

	// A Boolean value to indicate whether the specified compartment chooses to not to receive informational announcements by email.
	// (Manage preferences for receiving announcements by email by specifying the `preferenceType` attribute instead.)
	GetIsUnsubscribed() *bool

	// The OCID of the compartment for which you want to manage announcement email preferences. (Specify the tenancy by providing the
	// root compartment OCID.)
	GetCompartmentId() *string
}

type basecreateannouncementspreferencesdetails struct {
	JsonData       []byte
	PreferenceType BaseCreateAnnouncementsPreferencesDetailsPreferenceTypeEnum `mandatory:"true" json:"preferenceType"`
	IsUnsubscribed *bool                                                       `mandatory:"false" json:"isUnsubscribed"`
	CompartmentId  *string                                                     `mandatory:"false" json:"compartmentId"`
	Type           string                                                      `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *basecreateannouncementspreferencesdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerbasecreateannouncementspreferencesdetails basecreateannouncementspreferencesdetails
	s := struct {
		Model Unmarshalerbasecreateannouncementspreferencesdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.PreferenceType = s.Model.PreferenceType
	m.IsUnsubscribed = s.Model.IsUnsubscribed
	m.CompartmentId = s.Model.CompartmentId
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *basecreateannouncementspreferencesdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "CreateAnnouncementsPreferencesDetails":
		mm := CreateAnnouncementsPreferencesDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "UpdateAnnouncementsPreferencesDetails":
		mm := UpdateAnnouncementsPreferencesDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetPreferenceType returns PreferenceType
func (m basecreateannouncementspreferencesdetails) GetPreferenceType() BaseCreateAnnouncementsPreferencesDetailsPreferenceTypeEnum {
	return m.PreferenceType
}

//GetIsUnsubscribed returns IsUnsubscribed
func (m basecreateannouncementspreferencesdetails) GetIsUnsubscribed() *bool {
	return m.IsUnsubscribed
}

//GetCompartmentId returns CompartmentId
func (m basecreateannouncementspreferencesdetails) GetCompartmentId() *string {
	return m.CompartmentId
}

func (m basecreateannouncementspreferencesdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m basecreateannouncementspreferencesdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingBaseCreateAnnouncementsPreferencesDetailsPreferenceTypeEnum(string(m.PreferenceType)); !ok && m.PreferenceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PreferenceType: %s. Supported values are: %s.", m.PreferenceType, strings.Join(GetBaseCreateAnnouncementsPreferencesDetailsPreferenceTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// BaseCreateAnnouncementsPreferencesDetailsPreferenceTypeEnum Enum with underlying type: string
type BaseCreateAnnouncementsPreferencesDetailsPreferenceTypeEnum string

// Set of constants representing the allowable values for BaseCreateAnnouncementsPreferencesDetailsPreferenceTypeEnum
const (
	BaseCreateAnnouncementsPreferencesDetailsPreferenceTypeInTenantAnnouncements                 BaseCreateAnnouncementsPreferencesDetailsPreferenceTypeEnum = "OPT_IN_TENANT_ANNOUNCEMENTS"
	BaseCreateAnnouncementsPreferencesDetailsPreferenceTypeInTenantAndInformationalAnnouncements BaseCreateAnnouncementsPreferencesDetailsPreferenceTypeEnum = "OPT_IN_TENANT_AND_INFORMATIONAL_ANNOUNCEMENTS"
	BaseCreateAnnouncementsPreferencesDetailsPreferenceTypeOutAllAnnouncements                   BaseCreateAnnouncementsPreferencesDetailsPreferenceTypeEnum = "OPT_OUT_ALL_ANNOUNCEMENTS"
)

var mappingBaseCreateAnnouncementsPreferencesDetailsPreferenceTypeEnum = map[string]BaseCreateAnnouncementsPreferencesDetailsPreferenceTypeEnum{
	"OPT_IN_TENANT_ANNOUNCEMENTS":                   BaseCreateAnnouncementsPreferencesDetailsPreferenceTypeInTenantAnnouncements,
	"OPT_IN_TENANT_AND_INFORMATIONAL_ANNOUNCEMENTS": BaseCreateAnnouncementsPreferencesDetailsPreferenceTypeInTenantAndInformationalAnnouncements,
	"OPT_OUT_ALL_ANNOUNCEMENTS":                     BaseCreateAnnouncementsPreferencesDetailsPreferenceTypeOutAllAnnouncements,
}

// GetBaseCreateAnnouncementsPreferencesDetailsPreferenceTypeEnumValues Enumerates the set of values for BaseCreateAnnouncementsPreferencesDetailsPreferenceTypeEnum
func GetBaseCreateAnnouncementsPreferencesDetailsPreferenceTypeEnumValues() []BaseCreateAnnouncementsPreferencesDetailsPreferenceTypeEnum {
	values := make([]BaseCreateAnnouncementsPreferencesDetailsPreferenceTypeEnum, 0)
	for _, v := range mappingBaseCreateAnnouncementsPreferencesDetailsPreferenceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetBaseCreateAnnouncementsPreferencesDetailsPreferenceTypeEnumStringValues Enumerates the set of values in String for BaseCreateAnnouncementsPreferencesDetailsPreferenceTypeEnum
func GetBaseCreateAnnouncementsPreferencesDetailsPreferenceTypeEnumStringValues() []string {
	return []string{
		"OPT_IN_TENANT_ANNOUNCEMENTS",
		"OPT_IN_TENANT_AND_INFORMATIONAL_ANNOUNCEMENTS",
		"OPT_OUT_ALL_ANNOUNCEMENTS",
	}
}

// GetMappingBaseCreateAnnouncementsPreferencesDetailsPreferenceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBaseCreateAnnouncementsPreferencesDetailsPreferenceTypeEnum(val string) (BaseCreateAnnouncementsPreferencesDetailsPreferenceTypeEnum, bool) {
	mappingBaseCreateAnnouncementsPreferencesDetailsPreferenceTypeEnumIgnoreCase := make(map[string]BaseCreateAnnouncementsPreferencesDetailsPreferenceTypeEnum)
	for k, v := range mappingBaseCreateAnnouncementsPreferencesDetailsPreferenceTypeEnum {
		mappingBaseCreateAnnouncementsPreferencesDetailsPreferenceTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingBaseCreateAnnouncementsPreferencesDetailsPreferenceTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
