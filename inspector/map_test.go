package inspector_test

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/hlcfan/iop/inspector"
)

func TestInspectMap(t *testing.T) {
	t.Run("It inspects slice", func(t *testing.T) {
		var output bytes.Buffer

		person := map[string]interface{}{}
		person["name"] = "alex"
		person["age"] = 20
		person["father"] = false

		vType := reflect.TypeOf(person)
		vValue := reflect.ValueOf(person)

		mapInspector := inspector.NewMapInspector()
		mapInspector.Inspect(&output, vType, vValue)

		expected := "map[string]interface {} {\n\t\t\tname:\talex,\n\t\t\tage:\t20,\n\t\t\tfather:\tfalse,\n}\n"
		got := output.String()
		if got != expected {
			t.Errorf("Expect: %s, but got: %s", expected, got)
		}
	})
}
