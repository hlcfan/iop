package inspector_test

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/hlcfan/pp"
	"github.com/hlcfan/pp/inspector"
)

func TestInspectSlice(t *testing.T) {
	t.Run("It inspects struct slice", func(t *testing.T) {
		var output bytes.Buffer

		type person struct {
			ID    int
			Name  string
			Phone string
		}

		people := []person{
			{
				ID:    1,
				Name:  "alex",
				Phone: "12345678",
			},
		}

		vType := reflect.TypeOf(people)
		vValue := reflect.ValueOf(people)

		ioP := pp.New()
		ioP.SetOutput(&output)
		sliceInspector := inspector.NewSliceInspector()
		sliceInspector.Inspect(ioP, vType, vValue, 0)

		expected := "[]inspector_test.person {\n    {\n        ID:    1,\n        Name:  alex,\n        Phone: 12345678,\n    },\n}\n"
		got := output.String()
		// fmt.Printf("===: %#v\n", got)
		if got != expected {
			t.Errorf("Expect: %s, but got: %s", expected, got)
		}
	})

	t.Run("It inspects string slice", func(t *testing.T) {
		var output bytes.Buffer

		list := []string{"string1", "string2", "string3"}

		vType := reflect.TypeOf(list)
		vValue := reflect.ValueOf(list)

		ioP := pp.New()
		ioP.SetOutput(&output)
		sliceInspector := inspector.NewSliceInspector()
		sliceInspector.Inspect(ioP, vType, vValue, 0)

		expected := "[]string {\n    string1,\n    string2,\n    string3,\n}\n"
		got := output.String()
		// fmt.Printf("===: %#v\n", got)
		if got != expected {
			t.Errorf("Expect: %s, but got: %s", expected, got)
		}
	})
}
