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
		iop.SetOutput(&output)

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

		expected := "[]iop_test.person{\n\t{\n\t\tID:\t1,\n\t\tName:\talex,\n\t\tPhone:\t12345678,\n\t\tGraduated:\ttrue,\n\t\tCreatedAt:\t{\n\t\t\tTime:\t2009-11-17 20:34:58.651387237 +0000 UTC,\n\t\t\tValid:\ttrue,\n\t\t},\n\t\tAddresses: map[int]iop_test.address {\n\t\t\t1:\t{\n\t\t\t\tPostalCode:\t123,\n\t\t\t},\n\t\t\t2:\t{\n\t\t\t\tPostalCode:\t456,\n\t\t\t},\n\t\t}\n\t},\n\t{\n\t\tID:\t2,\n\t\tName:\tbob,\n\t\tPhone:\t87654321,\n\t\tGraduated:\tfalse,\n\t\tCreatedAt:\t{\n\t\t\tTime:\t2021-06-05 20:34:58.651387237 +0800 +08,\n\t\t\tValid:\ttrue,\n\t\t},\n\t\tAddresses: map[int]iop_test.address {\n\t\t\t1:\t{\n\t\t\t\tPostalCode:\t876,\n\t\t\t},\n\t\t\t2:\t{\n\t\t\t\tPostalCode:\t654,\n\t\t\t},\n\t\t}\n\t},\n}\n"
		got := output.String()
		// fmt.Printf("=%#v\n", got)
		if got != expected {
			t.Errorf("Expect: %s, but got: %s", expected, got)
		}
	})
}
