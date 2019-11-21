// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Big Data Service API
//
// API for the Big Data Service. Use this API to build, deploy, and manage fully elastic Big Data Service
// build on Hadoop, Spark and Data Science distribution, which can be fully integrated with existing enterprise
// data in Oracle Database and Oracle Applications..
//

package bds

// OperationTypesEnum Enum with underlying type: string
type OperationTypesEnum string

// Set of constants representing the allowable values for OperationTypesEnum
const (
	OperationTypesCreateBds               OperationTypesEnum = "CREATE_BDS"
	OperationTypesUpdateBds               OperationTypesEnum = "UPDATE_BDS"
	OperationTypesDeleteBds               OperationTypesEnum = "DELETE_BDS"
	OperationTypesAddWorkerNodes          OperationTypesEnum = "ADD_WORKER_NODES"
	OperationTypesAddCloudSql             OperationTypesEnum = "ADD_CLOUD_SQL"
	OperationTypesRemoveCloudSql          OperationTypesEnum = "REMOVE_CLOUD_SQL"
	OperationTypesChangeCompartmentForBds OperationTypesEnum = "CHANGE_COMPARTMENT_FOR_BDS"
)

var mappingOperationTypes = map[string]OperationTypesEnum{
	"CREATE_BDS":                 OperationTypesCreateBds,
	"UPDATE_BDS":                 OperationTypesUpdateBds,
	"DELETE_BDS":                 OperationTypesDeleteBds,
	"ADD_WORKER_NODES":           OperationTypesAddWorkerNodes,
	"ADD_CLOUD_SQL":              OperationTypesAddCloudSql,
	"REMOVE_CLOUD_SQL":           OperationTypesRemoveCloudSql,
	"CHANGE_COMPARTMENT_FOR_BDS": OperationTypesChangeCompartmentForBds,
}

// GetOperationTypesEnumValues Enumerates the set of values for OperationTypesEnum
func GetOperationTypesEnumValues() []OperationTypesEnum {
	values := make([]OperationTypesEnum, 0)
	for _, v := range mappingOperationTypes {
		values = append(values, v)
	}
	return values
}
