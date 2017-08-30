// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

// Region. A localized geographic area, such as Phoenix, AZ. Oracle Bare Metal Cloud Services is hosted in regions and Availability
// Domains. A region is composed of several Availability Domains. An Availability Domain is one or more data centers
// located within a region. For more information, see [Regions and Availability Domains](/Content/General/Concepts/regions.htm).
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access,
// see [Getting Started with Policies](/Content/Identity/Concepts/policygetstarted.htm).

type Region struct {

	// The key of the region.
	// Allowed values are:
	// - `PHX`
	// - `IAD`
	Key string `json:"key,omitempty"`

	// The name of the region.
	// Allowed values are:
	// - `us-phoenix-1`
	// - `us-ashburn-1`
	Name string `json:"name,omitempty"`
}
