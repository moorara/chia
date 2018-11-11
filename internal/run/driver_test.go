package run

import (
	"testing"

	"github.com/moorara/chia/internal/plan"
	"github.com/stretchr/testify/assert"
)

func TestNewTestDriver(t *testing.T) {
	tests := []struct {
		name               string
		file               string
		expectedError      string
		expectedTestDriver *testDriver
	}{
		{
			name:               "NotExist",
			file:               "./test/noplan.yaml",
			expectedError:      "no such file or directory",
			expectedTestDriver: nil,
		},
		{
			name:               "EmptyFile",
			file:               "./test/empty.yaml",
			expectedError:      "EOF",
			expectedTestDriver: nil,
		},
		{
			name:               "InvalidFile",
			file:               "./test/error.yaml",
			expectedError:      "cannot unmarshal",
			expectedTestDriver: nil,
		},
		{
			name:          "Simple",
			file:          "./test/simple.yaml",
			expectedError: "",
			expectedTestDriver: &testDriver{
				Plan: &plan.TestPlan{
					Name:         "Simple",
					RESTPlans:    []plan.RESTPlan{},
					GRPCPlans:    []plan.GRPCPlan{},
					NATSPlans:    []plan.NATSPlan{},
					GraphQLPlans: []plan.GraphQLPlan{},
				},
			},
		},
		{
			name:          "Full",
			file:          "./test/full.yaml",
			expectedError: "",
			expectedTestDriver: &testDriver{
				Plan: &plan.TestPlan{
					Name: "Full",
					RESTPlans: []plan.RESTPlan{
						plan.RESTPlan{
							Name:    "site-service",
							Address: "site-service:4010",
							Base:    "/v1",
							Headers: map[string]string{
								"Is-Test": "True",
							},
							Tests: []plan.RESTTest{
								plan.RESTTest{
									Name:     "CheckHealth",
									Method:   "GET",
									Endpoint: "/health",
									Headers: map[string]string{
										"Test-Name": "Check-Health",
									},
									Expect: plan.RESTExpect{
										StatusCode: 200,
									},
									Report: plan.RESTReport{},
								},
								plan.RESTTest{
									Name:     "GetSites",
									Method:   "GET",
									Endpoint: "/sites",
									Headers: map[string]string{
										"Test-Name": "Get-Sites",
									},
									Expect: plan.RESTExpect{
										StatusCode: 200,
										Headers:    map[string]string{},
										Body:       map[string]interface{}{},
									},
									Report: plan.RESTReport{},
								},
							},
						},
					},
					GRPCPlans: []plan.GRPCPlan{
						plan.GRPCPlan{
							Name:    "switch-service",
							Address: "switch-service:4030",
						},
					},
					NATSPlans: []plan.NATSPlan{
						plan.NATSPlan{
							Name:         "asset-service",
							NATSServers:  []string{"nats-0:4222", "nats-1:4222", "nats-2:4222"},
							NATSUser:     "chia",
							NATSPassword: "password",
						},
					},
					GraphQLPlans: []plan.GraphQLPlan{
						plan.GraphQLPlan{
							Name:    "graphql-server",
							Address: "graphql-server:5000",
						},
					},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			driver, err := NewTestDriver(tc.file)

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.Contains(t, err.Error(), tc.expectedError)
			}

			if tc.expectedTestDriver == nil {
				assert.Nil(t, driver)
			} else {
				testDriver, ok := driver.(*testDriver)
				assert.True(t, ok)
				assert.Equal(t, tc.expectedTestDriver, testDriver)
			}
		})
	}
}

func TestRun(t *testing.T) {
	tests := []struct {
		name          string
		driver        testDriver
		expectedError string
	}{
		{
			name:          "Empty",
			driver:        testDriver{},
			expectedError: "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.driver.Run()

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.Equal(t, tc.expectedError, err.Error())
			}
		})
	}
}
