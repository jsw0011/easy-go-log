package main

import (
	easyinitgolog "github.com/jsw0011/easy-init-go-log"
)

func initLoggers() {
	level := "DEBUG"
	out1 := "./log1.log"
	out2 := "./log2.log"
	easyinitgolog.InitLogger(&out1, &level, "module1")
	easyinitgolog.InitLogger(&out2, &level, "module2")

}

func logSomething() {
	l1 := easyinitgolog.GetLoggerByName("module1")
	l1.Debug("Debug1 before l2")
	l2 := easyinitgolog.GetLoggerByName("module2")
	l1.Debug("Debug1 after l2")
	l2.Debug("Debug2")
}

func main() {
	initLoggers()
	logSomething()
}
