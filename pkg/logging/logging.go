package logging

import (
	"github.com/indiependente/pkg/logger"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	proverbKey logger.LogKey = "proverb"
)

// ProverbLogger defines the behavior of the logger.
// Exposes a function for each loggable field which maps to a LogKey.
// The functions invocations can be chained and terminated by one of the levelled function calls (Fatal, Error, Warn, Info).
type ProverbLogger interface {
	logger.Logger
	Proverb(string) ProverbLogger
}

// PLogger implements the ProverbLogger interface and relies on http://github.com/rs/zerolog
type PLogger struct {
	logger.Logger
	lggr zerolog.Logger
}

// NewPLogger returns a new PLogger.
func NewPLogger(service string, level logger.LogLevel) PLogger {
	return PLogger{
		logger.GetLogger(service, level),
		log.With().Logger(),
	}
}

// Proverb instructs the logger to log the proverb.
func (l PLogger) Proverb(p string) ProverbLogger {
	lcopy := l
	lcopy.lggr = l.lggr.With().Str(proverbKey.String(), p).Logger()
	return lcopy
}
