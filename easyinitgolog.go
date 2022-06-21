package easyinitgolog

import (
	logging "github.com/op/go-logging"
	"os"
)

/**
 * logger can be retrieved by calling go-logging/GetLogger
 * @param level can be one of (CRITICAL, ERROR, WARNING, NOTICE, INFO, DEBUG)
 **/
func InitLogger(logPath *string, level *string, moduleName string) *logging.Logger {

	logFile := os.Stderr
	defLevel := logging.ERROR

	if logPath != nil {
		var file, err = os.OpenFile(*logPath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0640)
		if err != nil {
			panic("Unable to init logger with following file: " + *logPath + "\n" + err.Error())
		}
		logFile = file
	}

	if level != nil {
		var l, err = logging.LogLevel(*level)

		if err != nil {
			panic("Unable to set level: " + *level + "\n" + err.Error())
		}
		defLevel = l
	}

	// https://github.com/op/go-logging/blob/master/examples/example.go
	logger := logging.MustGetLogger(moduleName)
	backend1 := logging.NewLogBackend(logFile, "", 0)
	backend2 := logging.NewLogBackend(logFile, "", 0) // for logging error messages as pure text

	format := logging.MustStringFormatter(
		`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
	)

	formatFile := logging.MustStringFormatter(
		`%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x} %{message}`,
	)

	var backend2Formatter logging.Backend
	if logPath == nil {
		backend2Formatter = logging.NewBackendFormatter(backend2, format)
	} else {
		backend2Formatter = logging.NewBackendFormatter(backend2, formatFile)
	}

	backend1Leveled := logging.AddModuleLevel(backend1)

	backend1Leveled.SetLevel(logging.ERROR, "") // level ERROR for error messages

	logging.SetLevel(defLevel, moduleName)
	logger.SetBackend(logging.SetBackend(backend1Leveled, backend2Formatter))

	return logger
}
