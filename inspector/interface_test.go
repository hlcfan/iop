package inspector_test

import (
	"bytes"
	"reflect"
	"strconv"
	"testing"

	"github.com/hlcfan/pp"
	"github.com/hlcfan/pp/inspector"
)

func TestInspectInterface(t *testing.T) {
	t.Run("It inspects bool", func(t *testing.T) {
		tcs := []struct {
			id       int
			field    interface{}
			expected string
		}{
			{
				id:       1,
				field:    "alex",
				expected: "alex,\n",
			},
			{
				id:       2,
				field:    123,
				expected: "123,\n",
			},
			{
				id:       3,
				field:    []int{1, 2, 3},
				expected: "[1 2 3],\n",
			},
		}

		var output bytes.Buffer
		ioP := pp.New()
		ioP.SetOutput(&output)

		sliceInspector := inspector.NewInterfaceInspector()

		for _, tc := range tcs {
			t.Run(strconv.Itoa(tc.id), func(t *testing.T) {
				vType := reflect.TypeOf(tc.field)
				vValue := reflect.ValueOf(tc.field)

				sliceInspector.Inspect(ioP, vType, vValue, 0)

				got := output.String()
				// fmt.Printf("===Got: %#v\n", got)
				if got != tc.expected {
					t.Errorf("Case: %d, expect: %s, but got: %s", tc.id, tc.expected, got)
				}

				output.Reset()
			})
		}
	})
}
