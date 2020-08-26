package log

import (
	"bytes"
	"testing"
	"time"

	"bou.ke/monkey"
)

func TestWriter(t *testing.T) {
	// Patch time.Now to return a static time.
	monkey.Patch(time.Now, func() time.Time {
		return time.Time{}
	})
	defer monkey.UnpatchAll()

	buf := new(bytes.Buffer)
	w := NewWriter(buf)
	w.Debugf("%s message", "debug")
	w.Infof("%s message", "info")
	w.Warnf("%s message", "warn")
	w.Errorf("%s message", "error")

	exp := `[D] 0001-01-01 00:00:00 +0000 UTC debug message
[I] 0001-01-01 00:00:00 +0000 UTC info message
[W] 0001-01-01 00:00:00 +0000 UTC warn message
[E] 0001-01-01 00:00:00 +0000 UTC error message
`
	if buf.String() != exp {
		t.Errorf("writer got:\n%s\nwant:\n%s", buf.String(), exp)
	}
}
