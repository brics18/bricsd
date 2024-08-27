package prefixmanager

import (
	"github.com/brics18/bricsd/infrastructure/logger"
	"github.com/brics18/bricsd/util/panics"
)

var log = logger.RegisterSubSystem("PRFX")
var spawn = panics.GoroutineWrapperFunc(log)
