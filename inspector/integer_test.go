package inspector_test

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/hlcfan/pp"
	"github.com/hlcfan/pp/inspector"
)

func TestInspectInteger(t *testing.T) {
	t.Run("It inspects integer", func(t *testing.T) {
		var output bytes.Buffer

		integer := 1
		vType := reflect.TypeOf(integer)
		vValue := reflect.ValueOf(integer)

		ioP := pp.New()
		ioP.SetOutput(&output)

		sliceInspector := inspector.NewIntegerInspector()
		sliceInspector.Inspect(ioP, vType, vValue, 0)

		expected := "1,\n"
		got := output.String()
		if got != expected {
			t.Errorf("Expect: %s, but got: %s", expected, got)
		}
	})
}

func TestApplicable(t *testing.T) {
	t.Run("it returns true for integers", func(t *testing.T) {
		intInspector := inspector.NewIntegerInspector()
		integer := 1
		vType := reflect.TypeOf(integer)
		vValue := reflect.ValueOf(integer)
		applicable := intInspector.Applicable(vType, vValue)
		expected := true
		if applicable != expected {
			t.Errorf("Expect: %t, but got: %t", expected, applicable)
		}
	})
}
