package belajarlogging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLevel(t *testing.T) {
	logger := logrus.New()

	logger.Trace("Ini menggunakan Trace")
	logger.Debug("Ini menggunakan Debug")
	logger.Info("Ini menggunakan Info")
	logger.Warn("Ini menggunakan Warn")
	logger.Error("Ini menggunakan Error")
}
func TestLoggingLevel(t *testing.T) {
	logger := logrus.New()
	logger.SetLevel(logrus.TraceLevel)

	logger.Trace("Ini menggunakan Trace")
	logger.Debug("Ini menggunakan Debug")
	logger.Info("Ini menggunakan Info")
	logger.Warn("Ini menggunakan Warn")
	logger.Error("Ini menggunakan Error")
}
