// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

import (
	"github.com/oracle/oci-go-sdk/v31/common"
)

// BulkEditTagsResourceTypeCollection The list of resource types that support bulk editing of tags.
type BulkEditTagsResourceTypeCollection struct {

	// The collection of resource types that support bulk editing of tags.
	Items []BulkEditTagsResourceType `mandatory:"true" json:"items"`
}

func (m BulkEditTagsResourceTypeCollection) String() string {
	return common.PointerString(m)
}