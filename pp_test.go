package pp_test

import (
	"bytes"
	"database/sql"
	"reflect"
	"testing"
	"time"

	"github.com/hlcfan/pp"
)

const (
	defaultPrefix = ""
	defaultIndent = " "
)

type person struct {
	ID        int
	Name      string
	Phone     string
	Graduated bool
	CreatedAt sql.NullTime
	Addresses map[int]address
	vehicles  []vehicle
	NilField  *int
}

type vehicle struct {
	plate string
}

type address struct {
	PostalCode int
}

func TestPuts(t *testing.T) {
	t.Run("print to stdout", func(t *testing.T) {
		s := "im a string, Fran & Freddie's Diner	☺ "
		pp.Puts(s)
	})
}

func TestInspect(t *testing.T) {
	t.Run("print string", func(t *testing.T) {
		var b bytes.Buffer
		p := pp.New(&b, defaultPrefix, defaultIndent)
		p.SetOutput(&b)

		s := "im a string, Fran & Freddie's Diner	☺ "
		p.Inspect(reflect.ValueOf(s), 0)

		p.Flush()
		got := b.String()
		expected := "\"im a string, Fran & Freddie's Diner\\t☺ \"\n"

		if got != expected {
			t.Errorf("Expect: %v, but got: %v", expected, got)
		}
	})

	t.Run("print slice", func(t *testing.T) {
		var b bytes.Buffer
		p := pp.New(&b, defaultPrefix, defaultIndent)
		p.SetOutput(&b)

		list := []string{"abc", "123"}
		p.Inspect(reflect.ValueOf(list), 0)
		p.Flush()

		expected := "[\n \"abc\",\n \"123\",\n]\n"
		got := b.String()
		if got != expected {
			t.Errorf("Expect: %v, but got: %v", expected, got)
		}
	})

	t.Run("print struct", func(t *testing.T) {
		var b bytes.Buffer
		p := pp.New(&b, defaultPrefix, defaultIndent)
		p.SetOutput(&b)

		person := person{
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
			},
			vehicles: []vehicle{
				{
					plate: "CA-1234",
				},
			},
		}

		p.Inspect(reflect.ValueOf(person), 0)
		p.Flush()

		expected := "person {\n ID: 1,\n Name: \"alex\",\n Phone: \"12345678\",\n Graduated: true,\n CreatedAt: NullTime {\n  Time: Time {\n   wall: 651387237,\n   ext: 63394086898,\n   loc: nil,\n  },\n  Valid: true,\n },\n Addresses: {\n  1: address {\n   PostalCode: 123,\n  },\n },\n vehicles: [\n  vehicle {\n   plate: \"CA-1234\",\n  },\n ],\n NilField: nil,\n}\n"
		got := b.String()
		if got != expected {
			t.Errorf("Expect: %v, but got: %v", expected, got)
		}
	})

	t.Run("print map", func(t *testing.T) {
		var b bytes.Buffer
		p := pp.New(&b, defaultPrefix, defaultIndent)
		p.SetOutput(&b)

		m := map[string]string{"foo": "bar", "hello": "world"}
		p.Inspect(reflect.ValueOf(m), 0)
		p.Flush()

		got := b.String()
		expected1 := "{\n hello: \"world\",\n foo: \"bar\",\n}\n"
		expected2 := "{\n foo: \"bar\",\n hello: \"world\",\n}\n"
		allExpected := []string{expected1, expected2}

		var met bool
		for _, expected := range allExpected {
			if got == expected {
				met = true
			}
		}

		if !met {
			t.Errorf("Expect to be one of %v, but got: %v", allExpected, got)
		}
	})

	t.Run("print slice of structs", func(t *testing.T) {
		var b bytes.Buffer
		p := pp.New(&b, defaultPrefix, defaultIndent)
		p.SetOutput(&b)

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
				},
				vehicles: []vehicle{
					{
						plate: "CA-1234",
					},
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
					2: {PostalCode: 876},
				},
			},
		}

		p.Inspect(reflect.ValueOf(people), 0)
		p.Flush()

		got := b.String()
		expected := "[\n person {\n  ID: 1,\n  Name: \"alex\",\n  Phone: \"12345678\",\n  Graduated: true,\n  CreatedAt: NullTime {\n   Time: Time {\n    wall: 651387237,\n    ext: 63394086898,\n    loc: nil,\n   },\n   Valid: true,\n  },\n  Addresses: {\n   1: address {\n    PostalCode: 123,\n   },\n  },\n  vehicles: [\n   vehicle {\n    plate: \"CA-1234\",\n   },\n  ],\n  NilField: nil,\n },\n person {\n  ID: 2,\n  Name: \"bob\",\n  Phone: \"87654321\",\n  Graduated: false,\n  CreatedAt: NullTime {\n   Time: Time {\n    wall: 651387237,\n    ext: 63758493298,\n    loc: Location {\n     name: \"Local\",\n     zone: [\n      ,\n      ,\n      ,\n      ,\n      ,\n      ,\n      ,\n      ,\n     ],\n     tx: [\n      ,\n      ,\n      ,\n      ,\n      ,\n      ,\n      ,\n      ,\n     ],\n     extend: \"<+08>-8\",\n     cacheStart: 9223372036854775807,\n     cacheEnd: 9223372036854775807,\n     cacheZone: zone {\n      name: ,\n      offset: ,\n      isDST: ,\n     },\n    },\n   },\n   Valid: true,\n  },\n  Addresses: {\n   2: address {\n    PostalCode: 876,\n   },\n  },\n  vehicles: [\n  ],\n  NilField: nil,\n },\n]\n"
		if got != expected {
			t.Errorf("Expect: %s, but got: %v", expected, got)
		}
	})
}
