package inspector_test

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/hlcfan/iop"
	"github.com/hlcfan/iop/inspector"
)

func TestInspectSlice(t *testing.T) {
	t.Run("It inspects slice", func(t *testing.T) {
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

		ioP := iop.New()
		ioP.SetOutput(&output)
		sliceInspector := inspector.NewSliceInspector()
		sliceInspector.Inspect(ioP, vType, vValue, 0)

		expected := "[]inspector_test.person {\n\t{\n\t\tID:\t1,\n\t\tName:\talex,\n\t\tPhone:\t12345678,\n\t},\n}\n"
		got := output.String()
		// fmt.Printf("===: %v\n", got)
		if got != expected {
			t.Errorf("Expect: %s, but got: %s", expected, got)
		}
	})
}
