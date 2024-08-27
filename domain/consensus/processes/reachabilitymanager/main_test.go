package reachabilitymanager_test

import (
	"os"
	"testing"

	"github.com/brics18/bricsd/infrastructure/logger"
)

const logLevel = logger.LevelWarn

func TestMain(m *testing.M) {
	logger.SetLogLevels(logLevel)
	logger.InitLogStdout(logLevel)
	os.Exit(m.Run())
}
