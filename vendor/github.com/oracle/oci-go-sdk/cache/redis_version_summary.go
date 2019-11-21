// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// OraCache Public API
//
// API for the Data Caching Service. Use this service to manage Redis replicated caches.
//

package cache

import (
	"github.com/oracle/oci-go-sdk/common"
)

// RedisVersionSummary The Redis version number
type RedisVersionSummary struct {

	// Redis version
	Version *string `mandatory:"true" json:"version"`
}

func (m RedisVersionSummary) String() string {
	return common.PointerString(m)
}
