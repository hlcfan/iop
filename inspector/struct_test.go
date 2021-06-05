package inspector_test

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/hlcfan/iop"
	"github.com/hlcfan/iop/inspector"
)

type person struct {
	ID    int
	Name  string
	Phone string
}

func TestInspectStruct(t *testing.T) {
	t.Run("It inspects struct", func(t *testing.T) {
		var output bytes.Buffer

		person := person{
			ID:    1,
			Name:  "alex",
			Phone: "12345678",
		}

		vType := reflect.TypeOf(person)
		vValue := reflect.ValueOf(person)

		ioP := iop.New()
		ioP.SetOutput(&output)
		structInspector := inspector.NewStructInspector()
		structInspector.Inspect(ioP, vType, vValue, 0)

		expected := "inspector_test.person{\n\tID:\t1,\n\tName:\talex,\n\tPhone:\t12345678,\n}\n"
		got := output.String()
		if got != expected {
			t.Errorf("Expect: %s, but got: %s", expected, got)
		}
	})
}
