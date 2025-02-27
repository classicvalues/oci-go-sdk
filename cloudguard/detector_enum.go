// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard APIs
//
// A description of the Cloud Guard APIs
//

package cloudguard

import (
	"strings"
)

// DetectorEnumEnum Enum with underlying type: string
type DetectorEnumEnum string

// Set of constants representing the allowable values for DetectorEnumEnum
const (
	DetectorEnumIaasActivityDetector      DetectorEnumEnum = "IAAS_ACTIVITY_DETECTOR"
	DetectorEnumIaasConfigurationDetector DetectorEnumEnum = "IAAS_CONFIGURATION_DETECTOR"
)

var mappingDetectorEnumEnum = map[string]DetectorEnumEnum{
	"IAAS_ACTIVITY_DETECTOR":      DetectorEnumIaasActivityDetector,
	"IAAS_CONFIGURATION_DETECTOR": DetectorEnumIaasConfigurationDetector,
}

// GetDetectorEnumEnumValues Enumerates the set of values for DetectorEnumEnum
func GetDetectorEnumEnumValues() []DetectorEnumEnum {
	values := make([]DetectorEnumEnum, 0)
	for _, v := range mappingDetectorEnumEnum {
		values = append(values, v)
	}
	return values
}

// GetDetectorEnumEnumStringValues Enumerates the set of values in String for DetectorEnumEnum
func GetDetectorEnumEnumStringValues() []string {
	return []string{
		"IAAS_ACTIVITY_DETECTOR",
		"IAAS_CONFIGURATION_DETECTOR",
	}
}

// GetMappingDetectorEnumEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDetectorEnumEnum(val string) (DetectorEnumEnum, bool) {
	mappingDetectorEnumEnumIgnoreCase := make(map[string]DetectorEnumEnum)
	for k, v := range mappingDetectorEnumEnum {
		mappingDetectorEnumEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingDetectorEnumEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
