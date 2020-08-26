package log

import (
	"fmt"
	"io"
	"time"
)

type Writer struct {
	w io.Writer
}

func NewWriter(w io.Writer) *Writer {
	return &Writer{w: w}
}

func (w *Writer) log(level, msg string, args ...interface{}) {
	fmt.Fprintf(w.w, "[%s] %s %s\n", level, time.Now().String(),
		fmt.Sprintf(msg, args...))
}

// Debugf logs a Debugf message.
func (w *Writer) Debugf(format string, args ...interface{}) {
	w.log("D", format, args...)
}

// Infof logs an Infof message.
func (w *Writer) Infof(format string, args ...interface{}) {
	w.log("I", format, args...)
}

// Warnf logs a Warnf message.
func (w *Writer) Warnf(format string, args ...interface{}) {
	w.log("W", format, args...)
}

// Errorf logs an Errorf message.
func (w *Writer) Errorf(format string, args ...interface{}) {
	w.log("E", format, args...)
}
