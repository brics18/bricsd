package main

import (
	"github.com/brics18/bricsd/infrastructure/logger"
	"github.com/brics18/bricsd/util/panics"
)

var (
	backendLog = logger.NewBackend()
	log        = backendLog.Logger("ORPH")
	spawn      = panics.GoroutineWrapperFunc(log)
)
