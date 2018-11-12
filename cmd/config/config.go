package config

import (
	"github.com/moorara/goto/config"
)

const (
	defaultLogLevel        = "info"
	defaultPushgatewayAddr = "localhost:9091"
	defaultJaegerAgentAddr = "localhost:6831"
	defaultJaegerLogSpans  = false
)

// Config defines the configuration values
var Config = struct {
	LogLevel        string
	PushgatewayAddr string
	JaegerAgentAddr string
	JaegerLogSpans  bool
}{
	LogLevel:        defaultLogLevel,
	PushgatewayAddr: defaultPushgatewayAddr,
	JaegerAgentAddr: defaultJaegerAgentAddr,
	JaegerLogSpans:  defaultJaegerLogSpans,
}

func init() {
	config.Pick(&Config)
}
