package log

import (
	"fmt"
	"io"
	"time"

	"github.com/gookit/color"
)

// Color write messages to a writer in a k="v" format for all the
// fields on a single line *with* color!
type Color struct {
	w io.Writer
	f map[string]interface{}
}

// NewColor creates a color writer that sends messages to the given io.Writer.
func NewColor(w io.Writer) *Color {
	return &Color{w: w}
}

func (w *Color) log(level, msg string, args ...interface{}) {
	fmt.Fprintf(w.w, `when="%s" level="%s" message="%s"`, time.Now().String(), level,
		fmt.Sprintf(msg, args...))

	for k, v := range w.f {
		fmt.Fprintf(w.w, ` %s="%v"`, k, v)
	}
	fmt.Fprintf(w.w, "\n")
}

// Debugf logs a Debugf message.
func (w *Color) Debugf(format string, args ...interface{}) {
	color.Set(color.FgGray)
	w.log("debug", format, args...)
	color.Reset()
}

// Infof logs an Infof message.
func (w *Color) Infof(format string, args ...interface{}) {
	color.Set(color.FgWhite)
	w.log("info", format, args...)
	color.Reset()
}

// Warnf logs a Warnf message.
func (w *Color) Warnf(format string, args ...interface{}) {
	color.Set(color.FgMagenta)
	w.log("warn", format, args...)
	color.Reset()
}

// Errorf logs an Errorf message.
func (w *Color) Errorf(format string, args ...interface{}) {
	color.Set(color.FgRed)
	w.log("error", format, args...)
	color.Reset()
}

// WithFields calls WithFields for the default logger.
func (w *Color) WithFields(f map[string]interface{}) Logger {
	if w.f != nil {
		for k, v := range w.f {
			if _, ok := f[k]; !ok {
				f[k] = v
			}
		}
	}
	return &Color{w: w.w, f: f}
}
