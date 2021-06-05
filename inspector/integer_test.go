package inspector_test

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/hlcfan/iop"
	"github.com/hlcfan/iop/inspector"
)

func TestInspectInteger(t *testing.T) {
	t.Run("It inspects integer", func(t *testing.T) {
		var output bytes.Buffer

		integer := 1
		vType := reflect.TypeOf(integer)
		vValue := reflect.ValueOf(integer)

		ioP := iop.New()
		ioP.SetOutput(&output)

		sliceInspector := inspector.NewIntegerInspector()
		sliceInspector.Inspect(ioP, vType, vValue, 0)

		expected := "\t1,\n"
		got := output.String()
		if got != expected {
			t.Errorf("Expect: %s, but got: %s", expected, got)
		}
	})
}
