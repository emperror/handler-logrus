package logrus_test

import (
	"github.com/sirupsen/logrus"

	logrushandler "emperror.dev/handler/logrus"
)

func ExampleNew() {
	logger := logrus.New()
	_ = logrushandler.New(logger)

	// Output:
}
