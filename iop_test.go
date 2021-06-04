package iop_test

import (
	"bytes"
	"testing"

	"github.com/hlcfan/iop"
)

type Person struct {
	ID    int
	Name  string
	Phone string
}

func TestInspect(t *testing.T) {
	t.Run("It inspects slice", func(t *testing.T) {
		var output bytes.Buffer
		iop.SetOutput(&output)

		people := []Person{
			{
				ID:    1,
				Name:  "alex",
				Phone: "12345678",
			},
		}

		iop.Inspect(people)

		expected := "[]iop_test.Person {\n\t\t{\n\t\t\tID:\t1,\n\t\t\tName:\talex,\n\t\t\tPhone:\t12345678,\n\t\t},\n}\n"
		got := output.String()
		if got != expected {
			t.Errorf("Expect: %s, but got: %s", expected, got)
		}
	})
}
