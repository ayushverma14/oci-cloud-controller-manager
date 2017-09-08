// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package baremetal

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
)

type DatabaseTestSuite struct {
	suite.Suite
	requestor *mockRequestor
	nilHeader http.Header
	nilQuery  []interface{}
}

func (s *DatabaseTestSuite) SetupTest() {
	s.requestor = new(mockRequestor)
	s.requestor.Client = createClientForTest()
	s.requestor.databaseApi = s.requestor
}

func (s *DatabaseTestSuite) testDeleteResource(name resourceName, id string, funcUnderTest func(string, *IfMatchOptions) error) {
	option := &IfMatchOptions{IfMatch: "abcd"}

	details := &requestDetails{
		ids:      urlParts{id},
		name:     name,
		optional: option,
	}
	s.requestor.On("deleteRequest", details).Return(nil)

	e := funcUnderTest(id, option)
	s.requestor.AssertExpectations(s.T())
	s.Nil(e)
}

func TestRunDatabaseTests(t *testing.T) {
	suite.Run(t, new(DatabaseTestSuite))
}
