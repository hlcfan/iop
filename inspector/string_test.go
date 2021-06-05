package inspector_test

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/hlcfan/iop"
	"github.com/hlcfan/iop/inspector"
)

func TestInspectString(t *testing.T) {
	t.Run("It inspects string", func(t *testing.T) {
		var output bytes.Buffer

		text := "Howdy"
		vType := reflect.TypeOf(text)
		vValue := reflect.ValueOf(text)

		ioP := iop.New()
		ioP.SetOutput(&output)

		sliceInspector := inspector.NewStringInspector()
		sliceInspector.Inspect(ioP, vType, vValue, 0)

		expected := "\tHowdy,\n"
		got := output.String()
		if got != expected {
			t.Errorf("Expect: %s, but got: %s", expected, got)
		}
	})
}
