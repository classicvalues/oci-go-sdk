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
	"github.com/oracle/oci-go-sdk/v37/common"
)

// SqlPlanInsightAggregation SQL execution plan Performance statistics.
type SqlPlanInsightAggregation struct {

	// Plan hash value for the SQL Execution Plan
	PlanHash *int64 `mandatory:"true" json:"planHash"`

	// IO Time in seconds
	IoTimeInSec *float64 `mandatory:"true" json:"ioTimeInSec"`

	// CPU Time in seconds
	CpuTimeInSec *float64 `mandatory:"true" json:"cpuTimeInSec"`

	// Inefficient Wait Time in seconds
	InefficientWaitTimeInSec *float64 `mandatory:"true" json:"inefficientWaitTimeInSec"`

	// Total number of executions
	ExecutionsCount *int64 `mandatory:"true" json:"executionsCount"`
}

func (m SqlPlanInsightAggregation) String() string {
	return common.PointerString(m)
}
