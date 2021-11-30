// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Web Application Firewall (WAF) API
//
// API for the Web Application Firewall service.
// Use this API to manage regional Web App Firewalls and corresponding policies for protecting HTTP services.
//

package waf

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v53/common"
)

// CreateWebAppFirewallLoadBalancerDetails The information about new WebAppFirewallLoadBalancer.
type CreateWebAppFirewallLoadBalancerDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of WebAppFirewallPolicy, which is attached to the resource.
	WebAppFirewallPolicyId *string `mandatory:"true" json:"webAppFirewallPolicyId"`

	// LoadBalancer OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) to which the WebAppFirewallPolicy is attached to.
	LoadBalancerId *string `mandatory:"true" json:"loadBalancerId"`

	// WebAppFirewall display name, can be renamed.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

//GetDisplayName returns DisplayName
func (m CreateWebAppFirewallLoadBalancerDetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetCompartmentId returns CompartmentId
func (m CreateWebAppFirewallLoadBalancerDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetWebAppFirewallPolicyId returns WebAppFirewallPolicyId
func (m CreateWebAppFirewallLoadBalancerDetails) GetWebAppFirewallPolicyId() *string {
	return m.WebAppFirewallPolicyId
}

//GetFreeformTags returns FreeformTags
func (m CreateWebAppFirewallLoadBalancerDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m CreateWebAppFirewallLoadBalancerDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

//GetSystemTags returns SystemTags
func (m CreateWebAppFirewallLoadBalancerDetails) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m CreateWebAppFirewallLoadBalancerDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreateWebAppFirewallLoadBalancerDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateWebAppFirewallLoadBalancerDetails CreateWebAppFirewallLoadBalancerDetails
	s := struct {
		DiscriminatorParam string `json:"backendType"`
		MarshalTypeCreateWebAppFirewallLoadBalancerDetails
	}{
		"LOAD_BALANCER",
		(MarshalTypeCreateWebAppFirewallLoadBalancerDetails)(m),
	}

	return json.Marshal(&s)
}
