package inspector_test

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/hlcfan/pp"
	"github.com/hlcfan/pp/inspector"
)

type Person struct {
	ID    int
	Name  string
	Phone string
}

func TestInspectMap(t *testing.T) {
	t.Run("It inspects map", func(t *testing.T) {
		var output bytes.Buffer

		maps := map[string]interface{}{}
		maps["name"] = "alex"
		maps["age"] = 20
		maps["graduated"] = false

		// maps := map[string]Person{
		// 	"alex": {
		// 		ID:    1,
		// 		Name:  "alex",
		// 		Phone: "12345678",
		// 	},
		// }

		vType := reflect.TypeOf(maps)
		vValue := reflect.ValueOf(maps)

		ioP := pp.New()
		ioP.SetOutput(&output)
		mapInspector := inspector.NewMapInspector()
		mapInspector.Inspect(ioP, vType, vValue, 0)

		expectedCases := []string{
			"map[string]interface {} {\n    name:      alex,\n    age:       20,\n    graduated: false,\n}\n",
			"map[string]interface {} {\n    graduated: false,\n    name:      alex,\n    age:       20,\n}\n",
			"map[string]interface {} {\n    age:       20,\n    graduated: false,\n    name:      alex,\n}\n",
		}

		got := output.String()
		pass := false
		for _, expected := range expectedCases {
			if got == expected {
				pass = true
				break
			}
		}

		if !pass {
			t.Errorf("Expect: \n%v, but got: \n%v", expectedCases, got)
		}
	})
}
