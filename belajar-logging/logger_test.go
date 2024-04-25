package belajarlogging

import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
)

func Test(t *testing.T) {
	logger := logrus.New()
	logger.Println("Hello logger") //logger.Println itu == logger.info()

	fmt.Println("hallo fmt")

}
