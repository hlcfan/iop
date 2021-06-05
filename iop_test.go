package iop_test

import (
	"bytes"
	"database/sql"
	"testing"
	"time"

	"github.com/hlcfan/iop"
)

type person struct {
	ID        int
	Name      string
	Phone     string
	Graduated bool
	CreatedAt sql.NullTime
	Addresses map[int]address
}

type address struct {
	PostalCode int
}

func TestInspect(t *testing.T) {
	t.Run("It inspects slice", func(t *testing.T) {
		var output bytes.Buffer
		// iop.SetOutput(&output)

		people := []person{
			{
				ID:        1,
				Name:      "alex",
				Phone:     "12345678",
				Graduated: true,
				CreatedAt: sql.NullTime{
					Valid: true,
					Time: time.Date(
						2009, 11, 17, 20, 34, 58, 651387237, time.UTC),
				},
				Addresses: map[int]address{
					1: {PostalCode: 123},
					2: {PostalCode: 456},
				},
			},
			{
				ID:        2,
				Name:      "bob",
				Phone:     "87654321",
				Graduated: false,
				CreatedAt: sql.NullTime{
					Valid: true,
					Time: time.Date(
						2021, 06, 5, 20, 34, 58, 651387237, time.Local),
				},
				Addresses: map[int]address{
					1: {PostalCode: 876},
					2: {PostalCode: 654},
				},
			},
		}

		iop.Inspect(people)

		expected := "[]iop_test.person {\n\t\t{\n\t\t\tID:\t1,\n\t\t\tName:\talex,\n\t\t\tPhone:\t12345678,\n\t\t},\n}\n"
		got := output.String()
		if got != expected {
			t.Errorf("Expect: %s, but got: %s", expected, got)
		}
	})
}
