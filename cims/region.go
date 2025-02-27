// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Support Management API
//
// Use the Support Management API to manage support requests. For more information, see Getting Help and Contacting Support (https://docs.cloud.oracle.com/iaas/Content/GSG/Tasks/contactingsupport.htm). **Note**: Before you can create service requests with this API, you need to have an Oracle Single Sign On (SSO) account, and you need to register your Customer Support Identifier (CSI) with My Oracle Support.
//

package cims

import (
	"strings"
)

// RegionEnum Enum with underlying type: string
type RegionEnum string

// Set of constants representing the allowable values for RegionEnum
const (
	RegionDev          RegionEnum = "DEV"
	RegionSea          RegionEnum = "SEA"
	RegionIntegNext    RegionEnum = "INTEG_NEXT"
	RegionIntegStable  RegionEnum = "INTEG_STABLE"
	RegionPhx          RegionEnum = "PHX"
	RegionIad          RegionEnum = "IAD"
	RegionFra          RegionEnum = "FRA"
	RegionEuFrankfurt1 RegionEnum = "EU_FRANKFURT_1"
	RegionLhr          RegionEnum = "LHR"
	RegionYyz          RegionEnum = "YYZ"
	RegionNrt          RegionEnum = "NRT"
	RegionIcn          RegionEnum = "ICN"
	RegionBom          RegionEnum = "BOM"
	RegionGru          RegionEnum = "GRU"
	RegionSyd          RegionEnum = "SYD"
	RegionZrh          RegionEnum = "ZRH"
	RegionJed          RegionEnum = "JED"
	RegionAms          RegionEnum = "AMS"
	RegionKix          RegionEnum = "KIX"
	RegionMel          RegionEnum = "MEL"
	RegionYul          RegionEnum = "YUL"
	RegionHyd          RegionEnum = "HYD"
	RegionYny          RegionEnum = "YNY"
)

var mappingRegionEnum = map[string]RegionEnum{
	"DEV":            RegionDev,
	"SEA":            RegionSea,
	"INTEG_NEXT":     RegionIntegNext,
	"INTEG_STABLE":   RegionIntegStable,
	"PHX":            RegionPhx,
	"IAD":            RegionIad,
	"FRA":            RegionFra,
	"EU_FRANKFURT_1": RegionEuFrankfurt1,
	"LHR":            RegionLhr,
	"YYZ":            RegionYyz,
	"NRT":            RegionNrt,
	"ICN":            RegionIcn,
	"BOM":            RegionBom,
	"GRU":            RegionGru,
	"SYD":            RegionSyd,
	"ZRH":            RegionZrh,
	"JED":            RegionJed,
	"AMS":            RegionAms,
	"KIX":            RegionKix,
	"MEL":            RegionMel,
	"YUL":            RegionYul,
	"HYD":            RegionHyd,
	"YNY":            RegionYny,
}

// GetRegionEnumValues Enumerates the set of values for RegionEnum
func GetRegionEnumValues() []RegionEnum {
	values := make([]RegionEnum, 0)
	for _, v := range mappingRegionEnum {
		values = append(values, v)
	}
	return values
}

// GetRegionEnumStringValues Enumerates the set of values in String for RegionEnum
func GetRegionEnumStringValues() []string {
	return []string{
		"DEV",
		"SEA",
		"INTEG_NEXT",
		"INTEG_STABLE",
		"PHX",
		"IAD",
		"FRA",
		"EU_FRANKFURT_1",
		"LHR",
		"YYZ",
		"NRT",
		"ICN",
		"BOM",
		"GRU",
		"SYD",
		"ZRH",
		"JED",
		"AMS",
		"KIX",
		"MEL",
		"YUL",
		"HYD",
		"YNY",
	}
}

// GetMappingRegionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRegionEnum(val string) (RegionEnum, bool) {
	mappingRegionEnumIgnoreCase := make(map[string]RegionEnum)
	for k, v := range mappingRegionEnum {
		mappingRegionEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingRegionEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
