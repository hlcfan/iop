package inspector_test

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/hlcfan/iop"
	"github.com/hlcfan/iop/inspector"
)

type Person struct {
	ID    int
	Name  string
	Phone string
}

func TestInspectSlice(t *testing.T) {
	t.Run("It inspects slice", func(t *testing.T) {
		var output bytes.Buffer

		people := []Person{
			{
				ID:    1,
				Name:  "alex",
				Phone: "12345678",
			},
		}

		vType := reflect.TypeOf(people)
		vValue := reflect.ValueOf(people)

		ioP := iop.New()
		// ioP.SetOutput(&output)
		sliceInspector := inspector.NewSliceInspector()
		sliceInspector.Inspect(ioP, vType, vValue, 0)

		expected := "[]inspector_test.Person {\n\t\t{\n\t\t\tID:\t1,\n\t\t\tName:\talex,\n\t\t\tPhone:\t12345678,\n\t\t},\n}\n"
		got := output.String()
		if got != expected {
			t.Errorf("Expect: %s, but got: %s", expected, got)
		}
	})
}
