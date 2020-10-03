package log

import (
	"bytes"
	"reflect"
	"testing"
	"time"

	"bou.ke/monkey"
)

func TestColor(t *testing.T) {
	// Patch time.Now to return a static time.
	monkey.Patch(time.Now, func() time.Time {
		return time.Time{}
	})
	defer monkey.UnpatchAll()

	buf := new(bytes.Buffer)
	w := NewColor(buf)
	w.Debugf("%s message", "debug")
	w.Infof("%s message", "info")
	w.Warnf("%s message", "warn")
	w.Errorf("%s message", "error")
	w.WithFields(map[string]interface{}{"foo": "bar"}).Infof("hello")

	exp := `when="0001-01-01 00:00:00 +0000 UTC" level="debug" message="debug message"
when="0001-01-01 00:00:00 +0000 UTC" level="info" message="info message"
when="0001-01-01 00:00:00 +0000 UTC" level="warn" message="warn message"
when="0001-01-01 00:00:00 +0000 UTC" level="error" message="error message"
when="0001-01-01 00:00:00 +0000 UTC" level="info" message="hello" foo="bar"
`
	if buf.String() != exp {
		t.Errorf("writer got:\n%s\nwant:\n%s", buf.String(), exp)
	}
}

func TestColor_WithFields(t *testing.T) {
	// We just need to test overwriting.
	w := NewColor(nil)
	l := w.WithFields(map[string]interface{}{"foo": "bar", "too": "boo"}).
		WithFields(map[string]interface{}{"foo": "baz", "bar": "foo"})

	exp := map[string]interface{}{"foo": "baz", "too": "boo", "bar": "foo"}

	if !reflect.DeepEqual(l.(*Color).f, exp) {
		t.Errorf("with fields:\n%v\nwant:\n%v", l.(*Color).f, exp)
	}

}
