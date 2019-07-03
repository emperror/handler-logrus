package logrus_test

import (
	"github.com/sirupsen/logrus"

	logrushandler "handler.emperror.dev/logrus"
)

func ExampleNew() {
	logger := logrus.New()
	_ = logrushandler.New(logger)

	// Output:
}
