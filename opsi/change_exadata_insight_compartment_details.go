// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Operations Insights API
//
// Use the Operations Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Operations Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"github.com/oracle/oci-go-sdk/v53/common"
)

// ChangeExadataInsightCompartmentDetails The information to be updated.
type ChangeExadataInsightCompartmentDetails struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment into which the resource should be moved.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`
}

func (m ChangeExadataInsightCompartmentDetails) String() string {
	return common.PointerString(m)
}
