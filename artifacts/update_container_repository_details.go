// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Container Images API
//
// API covering the Registry (https://docs.cloud.oracle.com/iaas/Content/Registry/Concepts/registryoverview.htm) services.
// Use this API to manage resources such as container images and repositories.
//

package artifacts

import (
	"github.com/oracle/oci-go-sdk/v53/common"
)

// UpdateContainerRepositoryDetails Update container repository request details.
type UpdateContainerRepositoryDetails struct {

	// Whether the repository is immutable. Images cannot be overwritten in an immutable repository.
	IsImmutable *bool `mandatory:"false" json:"isImmutable"`

	// Whether the repository is public. A public repository allows unauthenticated access.
	IsPublic *bool `mandatory:"false" json:"isPublic"`

	Readme *ContainerRepositoryReadme `mandatory:"false" json:"readme"`
}

func (m UpdateContainerRepositoryDetails) String() string {
	return common.PointerString(m)
}
