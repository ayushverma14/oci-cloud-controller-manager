// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// DataCatalog API
//
// A description of the DataCatalog API
//

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/common"
)

// JobExecution A Job Execution is a unit of work being executed on behalf of a Job.
type JobExecution struct {

	// Unique key of the Job Execution resource.
	Key *string `mandatory:"true" json:"key"`

	// The unique key of the parent Job.
	JobKey *string `mandatory:"false" json:"jobKey"`

	// Type of the Job Execution.
	JobType JobTypeEnum `mandatory:"false" json:"jobType,omitempty"`

	// Sub-Type of this job execution.
	SubType *string `mandatory:"false" json:"subType"`

	// The unique key of the parent execution or null if this Job Execution has no parent.
	ParentKey *string `mandatory:"false" json:"parentKey"`

	// The unique key of the triggering external scheduler resource or null if this Job Execution is not externally triggered.
	ScheduleInstanceKey *string `mandatory:"false" json:"scheduleInstanceKey"`

	// Status of the Job Execution. For eg: Running, Paused, Completed etc
	LifecycleState JobExecutionStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The date and time the JobExecution was created, in the format defined by RFC3339.
	// Example: `2019-03-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Time that Job Execution started. An RFC3339 formatted datetime string.
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// Time that the Job Execution ended or null if it hasn't yet completed.
	// An RFC3339 formatted datetime string.
	TimeEnded *common.SDKTime `mandatory:"false" json:"timeEnded"`

	// Error code returned from the Job Execution ot null if job is still running or didn't return an error.
	ErrorCode *string `mandatory:"false" json:"errorCode"`

	// Error message returned from the Job Execution ot null if job is still running or didn't return an error.
	ErrorMessage *string `mandatory:"false" json:"errorMessage"`

	// Process identifier related to the Job Execution if the job is an external job.
	ProcessKey *string `mandatory:"false" json:"processKey"`

	// If the job is an external process, then a URL of the job for accessing this resource and its status.
	ExternalUrl *string `mandatory:"false" json:"externalUrl"`

	// This is related to the eventId used by OJDL. Such an event is created when OJDL receives an Http request,
	// and is then passed across the microservice layers for correlating log messages etc. May not be relevant
	// until OCI have such a facility. It is only used for correlating log messages across the layers.
	EventKey *string `mandatory:"false" json:"eventKey"`

	// The key of the associated Data Entity resource.
	DataEntityKey *string `mandatory:"false" json:"dataEntityKey"`

	// Id (OCID) of the user who created the Job Execution.
	CreatedById *string `mandatory:"false" json:"createdById"`

	// Id (OCID) of the user who updated the Job Execution.
	UpdatedBy *string `mandatory:"false" json:"updatedBy"`

	// URI to the Job Execution instance in the API.
	Uri *string `mandatory:"false" json:"uri"`

	// A map of maps which contains the execution context properties which are specific to a job execution. Each job
	// execution may define it's set of required and optional properties. The map keys are category names and the
	// values are maps of property name to property value. Every property is contained inside of a category. Most
	// job executions have required properties within the "default" category.
	// Example: `{"properties": { "default": { "host": "host1", "port": "1521", "database": "orcl"}}}`
	Properties map[string]map[string]string `mandatory:"false" json:"properties"`
}

func (m JobExecution) String() string {
	return common.PointerString(m)
}
