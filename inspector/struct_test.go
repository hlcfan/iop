package inspector_test

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/hlcfan/pp"
	"github.com/hlcfan/pp/inspector"
)

type person struct {
	ID      int
	Name    string
	Phone   string
	address string
}

func TestInspectStruct(t *testing.T) {
	t.Run("It inspects struct", func(t *testing.T) {
		var output bytes.Buffer

		person := person{
			ID:      1,
			Name:    "alex",
			Phone:   "12345678",
			address: "Singapore",
		}

		vType := reflect.TypeOf(person)
		vValue := reflect.ValueOf(person)

		ioP := pp.New()
		ioP.SetOutput(&output)
		structInspector := inspector.NewStructInspector()
		structInspector.Inspect(ioP, vType, vValue, 0)

		expected := "inspector_test.person{\n    ID:      1,\n    Name:    alex,\n    Phone:   12345678,\n    address: Singapore,\n}\n"
		got := output.String()
		// fmt.Printf("===: %#v\n", got)
		if got != expected {
			t.Errorf("Expect: %s, but got: %s", expected, got)
		}
	})
}
