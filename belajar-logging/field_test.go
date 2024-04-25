package belajarlogging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestField(t *testing.T) {
	logger := logrus.New()

	logger.WithField("username", "asep").Info("Hello info")
	logger.WithField("username", "mbut").WithField("name", "jem").Info("Hello info")
}

func TestFields(t *testing.T) {
	logger := logrus.New()

	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithFields(logrus.Fields{"username": "asep", "name": "kayu"}).Info("Hello info")

}
