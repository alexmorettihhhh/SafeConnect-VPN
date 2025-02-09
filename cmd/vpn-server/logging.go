package main

import (
	"github.com/sirupsen/logrus"
)

func logError(err error) {
	if err != nil {
		logrus.Error(err)
	}
}

func logInfo(message string) {
	logrus.Info(message)
}
