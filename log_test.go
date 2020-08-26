package log

import (
	"bytes"
	"os"
	"testing"
	"time"

	"bou.ke/monkey"
)

func TestSet(t *testing.T) {
	w := NewWriter(os.Stdout)
	type args struct {
		l Logger
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "set",
			args: args{l: w},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Set(tt.args.l)
			if logger != w {
				t.Errorf("set failed")
			}
		})
	}
}

func TestDefaultLogger(t *testing.T) {
	// Patch time.Now to return a static time.
	monkey.Patch(time.Now, func() time.Time {
		return time.Time{}
	})
	defer monkey.UnpatchAll()

	buf := new(bytes.Buffer)
	w := NewWriter(buf)
	Set(w)
	Debugf("%s message", "debug")
	Infof("%s message", "info")
	Warnf("%s message", "warn")
	Errorf("%s message", "error")

	exp := `[D] 0001-01-01 00:00:00 +0000 UTC debug message
[I] 0001-01-01 00:00:00 +0000 UTC info message
[W] 0001-01-01 00:00:00 +0000 UTC warn message
[E] 0001-01-01 00:00:00 +0000 UTC error message
`
	if buf.String() != exp {
		t.Errorf("writer got:\n%s\nwant:\n%s", buf.String(), exp)
	}
}

func TestNewMulti(t *testing.T) {
	// Patch time.Now to return a static time.
	monkey.Patch(time.Now, func() time.Time {
		return time.Time{}
	})
	defer monkey.UnpatchAll()

	buf1 := new(bytes.Buffer)
	w1 := NewWriter(buf1)

	buf2 := new(bytes.Buffer)
	w2 := NewWriter(buf2)

	m := NewMulti(w1, w2)

	m.Debugf("%s message", "debug")
	m.Infof("%s message", "info")
	m.Warnf("%s message", "warn")
	m.Errorf("%s message", "error")

	exp := `[D] 0001-01-01 00:00:00 +0000 UTC debug message
[I] 0001-01-01 00:00:00 +0000 UTC info message
[W] 0001-01-01 00:00:00 +0000 UTC warn message
[E] 0001-01-01 00:00:00 +0000 UTC error message
`
	if buf1.String() != exp {
		t.Errorf("buf1 got:\n%s\nwant:\n%s", buf1.String(), exp)
	}

	if buf2.String() != exp {
		t.Errorf("buf2 got:\n%s\nwant:\n%s", buf2.String(), exp)
	}

}
