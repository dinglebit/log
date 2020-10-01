package log

import (
	"fmt"
	"io"
	"time"
)

// Write write messages to a writer in a k="v" format for all the
// fields on a single line.
type Writer struct {
	w io.Writer
	f map[string]interface{}
}

// NewWrite creates a write that sends messages to the given io.Writer.
func NewWriter(w io.Writer) *Writer {
	return &Writer{w: w}
}

func (w *Writer) log(level, msg string, args ...interface{}) {
	fmt.Fprintf(w.w, `when="%s" level="%s" message="%s"`, time.Now().String(), level,
		fmt.Sprintf(msg, args...))

	for k, v := range w.f {
		fmt.Fprintf(w.w, ` %s="%v"`, k, v)
	}
	fmt.Fprintf(w.w, "\n")
}

// Debugf logs a Debugf message.
func (w *Writer) Debugf(format string, args ...interface{}) {
	w.log("debug", format, args...)
}

// Infof logs an Infof message.
func (w *Writer) Infof(format string, args ...interface{}) {
	w.log("info", format, args...)
}

// Warnf logs a Warnf message.
func (w *Writer) Warnf(format string, args ...interface{}) {
	w.log("warn", format, args...)
}

// Errorf logs an Errorf message.
func (w *Writer) Errorf(format string, args ...interface{}) {
	w.log("error", format, args...)
}

// WithFields calls WithFields for the default logger.
func (w *Writer) WithFields(f map[string]interface{}) Logger {
	if w.f != nil {
		for k, v := range w.f {
			if _, ok := f[k]; !ok {
				f[k] = v
			}
		}
	}
	return &Writer{w: w.w, f: f}
}
