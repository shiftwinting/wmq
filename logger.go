package main

import (
	"time"

	"github.com/Gurpartap/logrus-stack"
	"github.com/Sirupsen/logrus"
	"github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
)

//initLog
func initLog() {
	//log.Formatter = new(logrus.JSONFormatter)
	log.Level = logrus.DebugLevel

	callerLevels := logrus.AllLevels
	stackLevels := []logrus.Level{logrus.PanicLevel, logrus.FatalLevel, logrus.WarnLevel, logrus.ErrorLevel}
	logrus.AddHook(logrus_stack.NewHook(callerLevels, stackLevels))

	infoWriter, _ := rotatelogs.New(
		"log/info.%Y%m%d.log",
		rotatelogs.WithLinkName("log/info.log"),
		rotatelogs.WithMaxAge(time.Hour*24*7),
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	errorWriter, _ := rotatelogs.New(
		"log/error.%Y%m%d.log",
		rotatelogs.WithLinkName("log/error.log"),
		rotatelogs.WithMaxAge(time.Hour*24*7),
		rotatelogs.WithRotationTime(time.Hour),
	)

	log.Hooks.Add(lfshook.NewHook(lfshook.WriterMap{
		logrus.InfoLevel:  infoWriter,
		logrus.ErrorLevel: errorWriter,
		logrus.WarnLevel:  errorWriter,
	}))

}
