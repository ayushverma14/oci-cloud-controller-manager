/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"errors"
	"reflect"
	"testing"

	"github.com/oracle/oci-cloud-controller-manager/api/v1beta1"
	"github.com/oracle/oci-cloud-controller-manager/pkg/util"
	"github.com/oracle/oci-go-sdk/v50/core"
)

func TestComputeAveragesByReturnCode(t *testing.T) {
	testCases := []struct {
		name     string
		metrics  []ErrorMetric
		expected map[string]float64
	}{
		{
			name:     "base case",
			metrics:  nil,
			expected: map[string]float64{},
		},
		{
			name:     "base case 2",
			metrics:  endToEndLatencySlice{}.ErrorMetric(),
			expected: map[string]float64{},
		},
		{
			name: "base case e2e time",
			metrics: endToEndLatencySlice{
				endToEndLatency{timeTaken: 5.0},
				endToEndLatency{timeTaken: 10.0},
			}.ErrorMetric(),
			expected: map[string]float64{
				util.Success: 7.5,
			},
		},
		{
			name: "base case vnic attachment time",
			metrics: VnicAttachmentResponseSlice{
				VnicAttachmentResponse{timeTaken: 8.5},
				VnicAttachmentResponse{timeTaken: 6.5},
			}.ErrorMetric(),
			expected: map[string]float64{
				util.Success: 7.5,
			},
		},
		{
			name: "base case ip application time",
			metrics: IPAllocationSlice{
				IPAllocation{timeTaken: 5.0},
			}.ErrorMetric(),
			expected: map[string]float64{
				util.Success: 5.0,
			},
		},
		{
			name: "ip application failures",
			metrics: IPAllocationSlice{
				IPAllocation{timeTaken: 1.0, err: errors.New("http status code: 500")},
				IPAllocation{timeTaken: 2.0, err: errors.New("http status code: 500")},
				IPAllocation{timeTaken: 3.0, err: errors.New("http status code: 429")},
				IPAllocation{timeTaken: 4.0, err: errors.New("http status code: 401")},
				IPAllocation{timeTaken: 5.0},
				IPAllocation{timeTaken: 6.0},
			}.ErrorMetric(),
			expected: map[string]float64{
				util.Err5XX:  1.5,
				util.Err429:  3.0,
				util.Err4XX:  4.0,
				util.Success: 5.5,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			averages := computeAveragesByReturnCode(tc.metrics)
			if !reflect.DeepEqual(averages, tc.expected) {
				t.Errorf("expected metrics:\n%+v\nbut got:\n%+v", tc.expected, averages)
			}
		})
	}
}

var (
	trueVal      = true
	falseVal     = false
	testAddress1 = "1.1.1.1"
	testAddress2 = "2.2.2.2"
)

func TestFilterPrivateIp(t *testing.T) {
	testCases := []struct {
		name     string
		ips      []core.PrivateIp
		expected []core.PrivateIp
	}{
		{
			name:     "base case",
			ips:      []core.PrivateIp{},
			expected: []core.PrivateIp{},
		},
		{
			name: "only primary ip",
			ips: []core.PrivateIp{
				{IsPrimary: &trueVal},
			},
			expected: []core.PrivateIp{},
		},
		{
			name: "primary and secondary ip",
			ips: []core.PrivateIp{
				{IsPrimary: &trueVal},
				{IsPrimary: &falseVal, IpAddress: &testAddress1},
			},
			expected: []core.PrivateIp{
				{IsPrimary: &falseVal, IpAddress: &testAddress1},
			},
		},
		{
			name: "only secondary ip",
			ips: []core.PrivateIp{
				{IsPrimary: &falseVal, IpAddress: &testAddress1},
			},
			expected: []core.PrivateIp{
				{IsPrimary: &falseVal, IpAddress: &testAddress1},
			},
		},
		{
			name: "only termiated/termiatinging ips",
			ips: []core.PrivateIp{
				{IsPrimary: &falseVal, IpAddress: &testAddress1},
			},
			expected: []core.PrivateIp{
				{IsPrimary: &falseVal, IpAddress: &testAddress1},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			filtered := filterPrivateIp(tc.ips)
			if !reflect.DeepEqual(filtered, tc.expected) {
				t.Errorf("expected ips:\n%+v\nbut got:\n%+v", tc.expected, filtered)
			}
		})
	}
}

func TestTotalAllocatedSecondaryIpsForInstance(t *testing.T) {
	testCases := []struct {
		name     string
		ips      map[string][]core.PrivateIp
		expected int
	}{
		{
			name:     "base case",
			ips:      map[string][]core.PrivateIp{},
			expected: 0,
		},
		{
			name: "one vnic, two ips",
			ips: map[string][]core.PrivateIp{
				"one": {{IpAddress: &testAddress1}, {IpAddress: &testAddress2}},
			},
			expected: 2,
		},
		{
			name: "two vnic, 1/2 ips ",
			ips: map[string][]core.PrivateIp{
				"one": {{IpAddress: &testAddress1}, {IpAddress: &testAddress2}},
				"two": {{IpAddress: &testAddress2}},
			},
			expected: 3,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			allocated := totalAllocatedSecondaryIpsForInstance(tc.ips)
			if !reflect.DeepEqual(allocated, tc.expected) {
				t.Errorf("expected ip count:\n%+v\nbut got:\n%+v", tc.expected, allocated)
			}
		})
	}
}

func TestGetAdditionalSecondaryIPsNeededPerVNIC(t *testing.T) {
	testCases := []struct {
		name                   string
		existingIpsByVnic      map[string][]core.PrivateIp
		additionalSecondaryIps int
		expected               []VnicIPAllocations
		err                    error
	}{
		{
			name:                   "base case",
			existingIpsByVnic:      map[string][]core.PrivateIp{},
			additionalSecondaryIps: 0,
			expected:               []VnicIPAllocations{},
			err:                    nil,
		},
		{
			name: "one vnic with one additional IP required",
			existingIpsByVnic: map[string][]core.PrivateIp{
				"one": {{IpAddress: &testAddress1}},
			},
			additionalSecondaryIps: 1,
			expected:               []VnicIPAllocations{{"one", 1}},
			err:                    nil,
		},
		{
			name: "one vnic with space for required IPs",
			existingIpsByVnic: map[string][]core.PrivateIp{
				"one": {{IpAddress: &testAddress1}, {IpAddress: &testAddress1}, {IpAddress: &testAddress1}, {IpAddress: &testAddress1}, {IpAddress: &testAddress1}, {IpAddress: &testAddress1}, {IpAddress: &testAddress1},
					{IpAddress: &testAddress1}, {IpAddress: &testAddress1}, {IpAddress: &testAddress1}, {IpAddress: &testAddress1}, {IpAddress: &testAddress1}, {IpAddress: &testAddress1}, {IpAddress: &testAddress1}, {IpAddress: &testAddress1},
					{IpAddress: &testAddress1}, {IpAddress: &testAddress1}, {IpAddress: &testAddress1}},
			},
			additionalSecondaryIps: 13,
			expected:               []VnicIPAllocations{{"one", 13}},
			err:                    nil,
		},
		{
			name: "one vnic without space for required IPs",
			existingIpsByVnic: map[string][]core.PrivateIp{
				"one": {{IpAddress: &testAddress1}, {IpAddress: &testAddress1}, {IpAddress: &testAddress1}, {IpAddress: &testAddress1}, {IpAddress: &testAddress1}, {IpAddress: &testAddress1}, {IpAddress: &testAddress1},
					{IpAddress: &testAddress1}, {IpAddress: &testAddress1}, {IpAddress: &testAddress1}, {IpAddress: &testAddress1}, {IpAddress: &testAddress1}, {IpAddress: &testAddress1}, {IpAddress: &testAddress1}, {IpAddress: &testAddress1},
					{IpAddress: &testAddress1}, {IpAddress: &testAddress1}, {IpAddress: &testAddress1}},
			},
			additionalSecondaryIps: 31,
			expected:               nil,
			err:                    errors.New("failed to allocate the required number of IPs with existing VNICs"),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			allocation, err := getAdditionalSecondaryIPsNeededPerVNIC(tc.existingIpsByVnic, tc.additionalSecondaryIps)
			if (err == nil && tc.err != nil) || err != nil && tc.err == nil {
				t.Errorf("expected err:\n%+v\nbut got err:\n%+v", tc.err, err)
				t.FailNow()
			}
			if err != nil && err.Error() != tc.err.Error() {
				t.Errorf("expected err:\n%+v\nbut got err:\n%+v", tc.expected, allocation)
			}
			if !reflect.DeepEqual(allocation, tc.expected) {
				t.Errorf("expected ip allocation:\n%+v\nbut got:\n%+v", tc.expected, allocation)
			}
		})
	}
}

var (
	one         = "one"
	mac1        = "11.bb.cc.dd.ee.66"
	routerIP1   = "192.168.1.1"
	cidr1       = "10.0.0.0/64"
	subnetVnic1 = SubnetVnic{
		Vnic:   &core.Vnic{Id: &one, MacAddress: &mac1},
		Subnet: &core.Subnet{VirtualRouterIp: &routerIP1, CidrBlock: &cidr1},
	}
	npnVnic1 = v1beta1.VNICAddress{
		VNICID:     &one,
		MACAddress: &mac1,
		RouterIP:   &routerIP1,
		Addresses:  []*string{&testAddress1, &testAddress2},
		SubnetCidr: &cidr1,
	}
)

func TestConvertCoreVNICtoNPNStatus(t *testing.T) {
	testCases := []struct {
		name                   string
		existingSecondaryVNICs []SubnetVnic
		additionalSecondaryIps map[string][]core.PrivateIp
		expected               []v1beta1.VNICAddress
	}{
		{
			name:                   "base case",
			existingSecondaryVNICs: []SubnetVnic{},
			additionalSecondaryIps: map[string][]core.PrivateIp{},
			expected:               []v1beta1.VNICAddress{},
		},
		{
			name:                   "base case",
			existingSecondaryVNICs: []SubnetVnic{subnetVnic1},
			additionalSecondaryIps: map[string][]core.PrivateIp{
				one: {{IpAddress: &testAddress1}, {IpAddress: &testAddress2}},
			},
			expected: []v1beta1.VNICAddress{npnVnic1},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			vnics := convertCoreVNICtoNPNStatus(tc.existingSecondaryVNICs, tc.additionalSecondaryIps)
			if !reflect.DeepEqual(vnics, tc.expected) {
				t.Errorf("expected npnVNIC to be:\n%+v\nbut got:\n%+v", tc.expected, vnics)
			}
		})
	}
}
