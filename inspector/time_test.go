package inspector_test

import (
	"bytes"
	"reflect"
	"testing"
	"time"

	"github.com/hlcfan/pp"
	"github.com/hlcfan/pp/inspector"
)

func TestInspectTime(t *testing.T) {
	t.Run("It inspects struct", func(t *testing.T) {
		var output bytes.Buffer

		now := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)

		vType := reflect.TypeOf(now)
		vValue := reflect.ValueOf(now)

		ioP := pp.New()
		ioP.SetOutput(&output)
		structInspector := inspector.NewTimeInspector()
		structInspector.Inspect(ioP, vType, vValue, 0)

		expected := "\t2009-11-17 20:34:58.651387237 +0000 UTC\n"
		got := output.String()
		if got != expected {
			t.Errorf("Expect: %v, but got: %v", expected, got)
		}
	})
}
