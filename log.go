// Package log provides a simplified interface for interacting with
// logging systems.
package log

import "os"

// Logger is an interface this package uses to log to loggers that
// fulfull this interface.
type Logger interface {
	// Debugf logs very detailed information.
	Debugf(format string, args ...interface{})

	// Infof logs informative information.
	Infof(format string, args ...interface{})

	// Warnf logs information that suggests something my be wrong.
	Warnf(format string, args ...interface{})

	// Errorf logs information related to an error.
	Errorf(format string, args ...interface{})
}

var logger Logger = NewWriter(os.Stdout)

// Set the default logger for this package.
func Set(l Logger) {
	logger = l
}

// Debugf logs a Debugf message to the default logger.
func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

// Infof logs an Infof message to the default logger.
func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

// Warnf logs a Warnf message to the default logger.
func Warnf(format string, args ...interface{}) {
	logger.Warnf(format, args...)
}

// Errorf logs an Errorf message to the default logger.
func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}

// Multi creates a log which logs to all the given loggers.
func NewMulti(ll ...Logger) *Multi {
	return &Multi{ll: ll}
}

// Multi represents a Configurer that uses multiple configurations to
// provide values.
type Multi struct {
	ll []Logger
}

// Debugf calls Debugf for all the loggers.
func (m *Multi) Debugf(format string, args ...interface{}) {
	for _, l := range m.ll {
		l.Debugf(format, args...)
	}
}

// Infof calls Infos for all the loggers.
func (m *Multi) Infof(format string, args ...interface{}) {
	for _, l := range m.ll {
		l.Infof(format, args...)
	}
}

// Warnf calls Warnf for all the loggers.
func (m *Multi) Warnf(format string, args ...interface{}) {
	for _, l := range m.ll {
		l.Warnf(format, args...)
	}
}

// Errorf calls Errorf for all the loggers.
func (m *Multi) Errorf(format string, args ...interface{}) {
	for _, l := range m.ll {
		l.Errorf(format, args...)
	}
}
