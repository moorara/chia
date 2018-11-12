package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	tests := []struct {
		name                    string
		expectedLogLevel        string
		expectedPushgatewayAddr string
		expectedJaegerAgentAddr string
		expectedJaegerLogSpans  bool
	}{
		{
			name:                    "Defauts",
			expectedLogLevel:        defaultLogLevel,
			expectedPushgatewayAddr: defaultPushgatewayAddr,
			expectedJaegerAgentAddr: defaultJaegerAgentAddr,
			expectedJaegerLogSpans:  defaultJaegerLogSpans,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedLogLevel, Config.LogLevel)
			assert.Equal(t, tc.expectedPushgatewayAddr, Config.PushgatewayAddr)
			assert.Equal(t, tc.expectedJaegerAgentAddr, Config.JaegerAgentAddr)
			assert.Equal(t, tc.expectedJaegerLogSpans, Config.JaegerLogSpans)
		})
	}
}
