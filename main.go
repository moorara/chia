package main

import (
	"github.com/moorara/chia/cmd/config"
	"github.com/moorara/chia/cmd/version"
	"github.com/moorara/goto/log"
)

func main() {
	// Create logger
	logger := log.NewJSONLogger(config.Config.Name, config.Config.LogLevel)
	logger = logger.SyncLogger()
	logger = logger.With(
		config.Config.Name, map[string]string{
			"version":   version.Version,
			"revision":  version.Revision,
			"branch":    version.Branch,
			"goVersion": version.GoVersion,
			"buildTool": version.BuildTool,
			"buildTime": version.BuildTime,
		},
	)

	logger.Info("message", "Chia started.")
}
