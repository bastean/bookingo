package logger

import (
	"github.com/bastean/bookingo/pkg/context/shared/infrastructure/loggers"
)

var Logger = new(loggers.Logger)

var Debug = Logger.Debug

var Error = Logger.Error

var Fatal = Logger.Fatal

var Info = Logger.Info
