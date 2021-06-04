package inspector_test

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/hlcfan/iop/inspector"
)

func TestInspectInt(t *testing.T) {
	t.Run("It inspects slice", func(t *testing.T) {
		var output bytes.Buffer

		integer := 1
		vType := reflect.TypeOf(integer)
		vValue := reflect.ValueOf(integer)

		sliceInspector := inspector.NewIntegerInspector()
		sliceInspector.Inspect(&output, vType, vValue)

		expected := "1\n"
		got := output.String()
		if got != expected {
			t.Errorf("Expect: %s, but got: %s", expected, got)
		}
	})
}
