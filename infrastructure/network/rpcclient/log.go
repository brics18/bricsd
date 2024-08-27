package rpcclient

import (
	"github.com/brics18/bricsd/infrastructure/logger"
	"github.com/brics18/bricsd/util/panics"
)

var log = logger.RegisterSubSystem("RPCC")
var spawn = panics.GoroutineWrapperFunc(log)
