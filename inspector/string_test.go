package inspector_test

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/hlcfan/pp"
	"github.com/hlcfan/pp/inspector"
)

func TestInspectString(t *testing.T) {
	t.Run("It inspects string", func(t *testing.T) {
		tcs := []struct {
			actual   string
			level    int
			expected string
		}{
			{
				actual:   "Howdy",
				level:    1,
				expected: "\tHowdy,\n",
			},
			{
				actual:   "Howdy",
				level:    0,
				expected: "Howdy\n",
			},
		}

		for _, tc := range tcs {
			var output bytes.Buffer

			vType := reflect.TypeOf(tc.actual)
			vValue := reflect.ValueOf(tc.actual)

			ioP := pp.New()
			ioP.SetOutput(&output)

			sliceInspector := inspector.NewStringInspector()
			sliceInspector.Inspect(ioP, vType, vValue, tc.level)

			expected := "Howdy,\n"
			got := output.String()
			if got != expected {
				t.Errorf("Expect: %s, but got: %s", expected, got)
			}
		}
	})
}
