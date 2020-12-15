// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard APIs
//
// A description of the Cloud Guard APIs
//

package cloudguard

import (
	"github.com/oracle/oci-go-sdk/v31/common"
)

// UpdateResponderRuleDetails Details of UpdateResponderRuleDetails.
type UpdateResponderRuleDetails struct {

	// Identifies state for ResponderRule
	IsEnabled *bool `mandatory:"true" json:"isEnabled"`
}

func (m UpdateResponderRuleDetails) String() string {
	return common.PointerString(m)
}