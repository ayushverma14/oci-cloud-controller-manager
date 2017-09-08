// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package baremetal

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
)

type LoadbalancerTestSuite struct {
	suite.Suite
	requestor *mockRequestor
	nilHeader http.Header
}

func (s *LoadbalancerTestSuite) SetupTest() {
	s.requestor = new(mockRequestor)
	s.requestor.Client = createClientForTest()
	s.requestor.loadBalancerApi = s.requestor
}

func TestRunLoadbalancerTests(t *testing.T) {
	suite.Run(t, new(LoadbalancerTestSuite))
}
