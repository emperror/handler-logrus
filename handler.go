// Package logrus provides Logrus integration.
package logrus

import (
	"context"

	"emperror.dev/errors"
	"emperror.dev/errors/utils/keyval"
	"github.com/sirupsen/logrus"
)

// Handler logs errors using Logrus.
type Handler struct {
	logger logrus.FieldLogger
}

// New creates a new handler.
func New(logger logrus.FieldLogger) *Handler {
	return &Handler{
		logger: logger,
	}
}

// Handle logs an error.
func (h *Handler) Handle(err error) {
	logger := h.logger

	// Extract details from the error and attach it to the log
	if details := errors.GetDetails(err); len(details) > 0 {
		logger = h.logger.WithFields(keyval.ToMap(details))
	}

	type errorCollection interface {
		Errors() []error
	}

	if errs, ok := err.(errorCollection); ok {
		for _, e := range errs.Errors() {
			logger.WithField("parent", err.Error()).Error(e.Error())
		}
	} else {
		logger.Error(err.Error())
	}
}

// HandleContext logs an error.
func (h *Handler) HandleContext(_ context.Context, err error) {
	h.Handle(err)
}
