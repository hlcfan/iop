package inspector_test

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"

	"github.com/hlcfan/pp"
	"github.com/hlcfan/pp/inspector"
)

func TestInspectBool(t *testing.T) {
	t.Run("It inspects bool", func(t *testing.T) {
		tcs := []struct {
			actual   bool
			level    int
			expected string
		}{
			{
				actual:   true,
				level:    0,
				expected: "true\n",
			},
			{
				actual:   true,
				level:    1,
				expected: "true,\n",
			},
			{
				actual:   false,
				level:    0,
				expected: "false\n",
			},
			{
				actual:   false,
				level:    1,
				expected: "false,\n",
			},
		}

		for _, tc := range tcs {
			var output bytes.Buffer

			vType := reflect.TypeOf(tc.actual)
			vValue := reflect.ValueOf(tc.actual)

			ioP := pp.New()
			ioP.SetOutput(&output)

			sliceInspector := inspector.NewBoolInspector()
			sliceInspector.Inspect(ioP, vType, vValue, tc.level)

			got := output.String()
			fmt.Printf("===Got: %#v\n", got)
			fmt.Printf("===Exp: %#v\n", tc.expected)
			if got != tc.expected {
				t.Errorf("Expect: %s, but got: %s", tc.expected, got)
			}
		}
	})
}
