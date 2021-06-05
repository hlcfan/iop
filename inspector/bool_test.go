package inspector_test

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/hlcfan/iop"
	"github.com/hlcfan/iop/inspector"
)

func TestInspectBool(t *testing.T) {
	t.Run("It inspects bool", func(t *testing.T) {
		var output bytes.Buffer

		trueValue := true
		vType := reflect.TypeOf(trueValue)
		vValue := reflect.ValueOf(trueValue)

		ioP := iop.New()
		ioP.SetOutput(&output)

		sliceInspector := inspector.NewBoolInspector()
		sliceInspector.Inspect(ioP, vType, vValue, 0)

		expected := "true,\n"
		got := output.String()
		if got != expected {
			t.Errorf("Expect: %s, but got: %s", expected, got)
		}
	})
}
