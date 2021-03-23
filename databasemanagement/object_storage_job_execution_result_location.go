// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to perform tasks such as obtaining performance and resource usage metrics
// for a fleet of Managed Databases or a specific Managed Database, creating Managed Database Groups, and
// running a SQL job on a Managed Database or Managed Database Group.
//

package databasemanagement

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v37/common"
)

// ObjectStorageJobExecutionResultLocation The details about Object Storage job execution result location type.
type ObjectStorageJobExecutionResultLocation struct {

	// The Object Storage namespace used for job execution result storage.
	NamespaceName *string `mandatory:"false" json:"namespaceName"`

	// The name of the bucket used for job execution result storage.
	BucketName *string `mandatory:"false" json:"bucketName"`
}

func (m ObjectStorageJobExecutionResultLocation) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m ObjectStorageJobExecutionResultLocation) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeObjectStorageJobExecutionResultLocation ObjectStorageJobExecutionResultLocation
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeObjectStorageJobExecutionResultLocation
	}{
		"OBJECT_STORAGE",
		(MarshalTypeObjectStorageJobExecutionResultLocation)(m),
	}

	return json.Marshal(&s)
}
