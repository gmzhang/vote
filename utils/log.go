package utils

import (
	"github.com/sirupsen/logrus"
)

// Internal levels of library output that are initialised to not print
// anything but can be overridden by programmer
var (
	Logger *logrus.Logger
)

func init() {
	Logger = logrus.New()
	Logger.Formatter = &logrus.JSONFormatter{}
}
