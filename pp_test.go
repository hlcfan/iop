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
		person := person{
			ID:        1,
			Name:      "alex 😊",
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

		pp.Puts(person)

		m := map[string]string{"foo": "bar", "hello": "world"}
		pp.PutsWithLabel(m, "DEBUG")
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
		expected := "\x1b[38;2;135;255;215m\"im a string, Fran & Freddie's Diner\\t☺ \"\x1b[0m"

		if got != expected {
			t.Errorf("Expect: %#v, but got: %#v", expected, got)
		}
	})

	t.Run("print slice", func(t *testing.T) {
		var b bytes.Buffer
		p := pp.New(&b, defaultPrefix, defaultIndent)
		p.SetOutput(&b)

		list := []string{"abc", "123"}
		p.Inspect(reflect.ValueOf(list), 0)
		p.Flush()

		expected := "[\n \x1b[38;2;135;255;215m\"abc\"\x1b[0m\x1b[38;2;255;121;198m,\x1b[0m\n \x1b[38;2;135;255;215m\"123\"\x1b[0m\x1b[38;2;255;121;198m,\x1b[0m\n]"
		got := b.String()
		if got != expected {
			t.Errorf("Expect: %#v, but got: %#v", expected, got)
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

		expected := "\x1b[38;2;174;129;255mperson\x1b[0m {\n ID\x1b[38;2;255;121;198m: \x1b[0m\x1b[38;2;239;228;161m1\x1b[0m\x1b[38;2;255;121;198m,\x1b[0m\n Name\x1b[38;2;255;121;198m: \x1b[0m\x1b[38;2;135;255;215m\"alex\"\x1b[0m\x1b[38;2;255;121;198m,\x1b[0m\n Phone\x1b[38;2;255;121;198m: \x1b[0m\x1b[38;2;135;255;215m\"12345678\"\x1b[0m\x1b[38;2;255;121;198m,\x1b[0m\n Graduated\x1b[38;2;255;121;198m: \x1b[0m\x1b[38;2;239;228;161mtrue\x1b[0m\x1b[38;2;255;121;198m,\x1b[0m\n CreatedAt\x1b[38;2;255;121;198m: \x1b[0m\x1b[38;2;174;129;255mNullTime\x1b[0m {\n  Time\x1b[38;2;255;121;198m: \x1b[0m\x1b[38;2;174;129;255mTime\x1b[0m {\n   wall\x1b[38;2;255;121;198m: \x1b[0m\x1b[38;2;239;228;161m651387237\x1b[0m\x1b[38;2;255;121;198m,\x1b[0m\n   ext\x1b[38;2;255;121;198m: \x1b[0m\x1b[38;2;239;228;161m63394086898\x1b[0m\x1b[38;2;255;121;198m,\x1b[0m\n   loc\x1b[38;2;255;121;198m: \x1b[0mnil\x1b[38;2;255;121;198m,\x1b[0m\n  }\x1b[38;2;255;121;198m,\x1b[0m\n  Valid\x1b[38;2;255;121;198m: \x1b[0m\x1b[38;2;239;228;161mtrue\x1b[0m\x1b[38;2;255;121;198m,\x1b[0m\n }\x1b[38;2;255;121;198m,\x1b[0m\n Addresses\x1b[38;2;255;121;198m: \x1b[0m{\n  \x1b[38;2;239;228;161m1\x1b[0m\x1b[38;2;255;121;198m: \x1b[0m\x1b[38;2;174;129;255maddress\x1b[0m {\n   PostalCode\x1b[38;2;255;121;198m: \x1b[0m\x1b[38;2;239;228;161m123\x1b[0m\x1b[38;2;255;121;198m,\x1b[0m\n  }\x1b[38;2;255;121;198m,\x1b[0m\n }\x1b[38;2;255;121;198m,\x1b[0m\n vehicles\x1b[38;2;255;121;198m: \x1b[0m[\n  \x1b[38;2;174;129;255mvehicle\x1b[0m {\n   plate\x1b[38;2;255;121;198m: \x1b[0m\x1b[38;2;135;255;215m\"CA-1234\"\x1b[0m\x1b[38;2;255;121;198m,\x1b[0m\n  }\x1b[38;2;255;121;198m,\x1b[0m\n ]\x1b[38;2;255;121;198m,\x1b[0m\n NilField\x1b[38;2;255;121;198m: \x1b[0mnil\x1b[38;2;255;121;198m,\x1b[0m\n}"
		got := b.String()
		if got != expected {
			t.Errorf("Expect: %#v, but got: %#v", expected, got)
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
		expected1 := "{\n \x1b[38;2;135;255;215m\"hello\"\x1b[0m\x1b[38;2;255;121;198m: \x1b[0m\x1b[38;2;135;255;215m\"world\"\x1b[0m\x1b[38;2;255;121;198m,\x1b[0m\n \x1b[38;2;135;255;215m\"foo\"\x1b[0m\x1b[38;2;255;121;198m: \x1b[0m\x1b[38;2;135;255;215m\"bar\"\x1b[0m\x1b[38;2;255;121;198m,\x1b[0m\n}"
		expected2 := "{\n \x1b[38;2;135;255;215m\"foo\"\x1b[0m\x1b[38;2;255;121;198m: \x1b[0m\x1b[38;2;135;255;215m\"bar\"\x1b[0m\x1b[38;2;255;121;198m,\x1b[0m\n \x1b[38;2;135;255;215m\"hello\"\x1b[0m\x1b[38;2;255;121;198m: \x1b[0m\x1b[38;2;135;255;215m\"world\"\x1b[0m\x1b[38;2;255;121;198m,\x1b[0m\n}"
		allExpected := []string{expected1, expected2}

		var met bool
		for _, expected := range allExpected {
			if got == expected {
				met = true
			}
		}

		if !met {
			t.Errorf("Expect to be one of %#v, but got: %#v", allExpected, got)
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
		expected := "[\n \x1b[38;2;174;129;255mperson\x1b[0m {\n  ID\x1b[38;2;255;121;198m: \x1b[0m\x1b[38;2;239;228;161m1\x1b[0m\x1b[38;2;255;121;198m,\x1b[0m\n  Name\x1b[38;2;255;121;198m: \x1b[0m\x1b[38;2;135;255;215m\"alex\"\x1b[0m\x1b[38;2;255;121;198m,\x1b[0m\n  Phone\x1b[38;2;255;121;198m: \x1b[0m\x1b[38;2;135;255;215m\"12345678\"\x1b[0m\x1b[38;2;255;121;198m,\x1b[0m\n  Graduated\x1b[38;2;255;121;198m: \x1b[0m\x1b[38;2;239;228;161mtrue\x1b[0m\x1b[38;2;255;121;198m,\x1b[0m\n  CreatedAt\x1b[38;2;255;121;198m: \x1b[0m\x1b[38;2;174;129;255mNullTime\x1b[0m {\n   Time\x1b[38;2;255;121;198m: \x1b[0m\x1b[38;2;174;129;255mTime\x1b[0m {\n    wall\x1b[38;2;255;121;198m: \x1b[0m\x1b[38;2;239;228;161m651387237\x1b[0m\x1b[38;2;255;121;198m,\x1b[0m\n    ext\x1b[38;2;255;121;198m: \x1b[0m\x1b[38;2;239;228;161m63394086898\x1b[0m\x1b[38;2;255;121;198m,\x1b[0m\n    loc\x1b[38;2;255;121;198m: \x1b[0mnil\x1b[38;2;255;121;198m,\x1b[0m\n   }\x1b[38;2;255;121;198m,\x1b[0m\n   Valid\x1b[38;2;255;121;198m: \x1b[0m\x1b[38;2;239;228;161mtrue\x1b[0m\x1b[38;2;255;121;198m,\x1b[0m\n  }\x1b[38;2;255;121;198m,\x1b[0m\n  Addresses\x1b[38;2;255;121;198m: \x1b[0m{\n   \x1b[38;2;239;228;161m1\x1b[0m\x1b[38;2;255;121;198m: \x1b[0m\x1b[38;2;174;129;255maddress\x1b[0m {\n    PostalCode\x1b[38;2;255;121;198m: \x1b[0m\x1b[38;2;239;228;161m123\x1b[0m\x1b[38;2;255;121;198m,\x1b[0m\n   }\x1b[38;2;255;121;198m,\x1b[0m\n  }\x1b[38;2;255;121;198m,\x1b[0m\n  vehicles\x1b[38;2;255;121;198m: \x1b[0m[\n   \x1b[38;2;174;129;255mvehicle\x1b[0m {\n    plate\x1b[38;2;255;121;198m: \x1b[0m\x1b[38;2;135;255;215m\"CA-1234\"\x1b[0m\x1b[38;2;255;121;198m,\x1b[0m\n   }\x1b[38;2;255;121;198m,\x1b[0m\n  ]\x1b[38;2;255;121;198m,\x1b[0m\n  NilField\x1b[38;2;255;121;198m: \x1b[0mnil\x1b[38;2;255;121;198m,\x1b[0m\n }\x1b[38;2;255;121;198m,\x1b[0m\n \x1b[38;2;174;129;255mperson\x1b[0m {\n  ID\x1b[38;2;255;121;198m: \x1b[0m\x1b[38;2;239;228;161m2\x1b[0m\x1b[38;2;255;121;198m,\x1b[0m\n  Name\x1b[38;2;255;121;198m: \x1b[0m\x1b[38;2;135;255;215m\"bob\"\x1b[0m\x1b[38;2;255;121;198m,\x1b[0m\n  Phone\x1b[38;2;255;121;198m: \x1b[0m\x1b[38;2;135;255;215m\"87654321\"\x1b[0m\x1b[38;2;255;121;198m,\x1b[0m\n  Graduated\x1b[38;2;255;121;198m: \x1b[0m\x1b[38;2;239;228;161mfalse\x1b[0m\x1b[38;2;255;121;198m,\x1b[0m\n  CreatedAt\x1b[38;2;255;121;198m: \x1b[0m\x1b[38;2;174;129;255mNullTime\x1b[0m {\n   Time\x1b[38;2;255;121;198m: \x1b[0m\x1b[38;2;174;129;255mTime\x1b[0m {\n    wall\x1b[38;2;255;121;198m: \x1b[0m\x1b[38;2;239;228;161m651387237\x1b[0m\x1b[38;2;255;121;198m,\x1b[0m\n    ext\x1b[38;2;255;121;198m: \x1b[0m\x1b[38;2;239;228;161m63758493298\x1b[0m\x1b[38;2;255;121;198m,\x1b[0m\n    loc\x1b[38;2;255;121;198m: \x1b[0m\x1b[38;2;174;129;255mLocation\x1b[0m {\n     name\x1b[38;2;255;121;198m: \x1b[0m\x1b[38;2;135;255;215m\"Local\"\x1b[0m\x1b[38;2;255;121;198m,\x1b[0m\n     zone\x1b[38;2;255;121;198m: \x1b[0m[\n      \x1b[38;2;255;121;198m,\x1b[0m\n      \x1b[38;2;255;121;198m,\x1b[0m\n      \x1b[38;2;255;121;198m,\x1b[0m\n      \x1b[38;2;255;121;198m,\x1b[0m\n      \x1b[38;2;255;121;198m,\x1b[0m\n      \x1b[38;2;255;121;198m,\x1b[0m\n      \x1b[38;2;255;121;198m,\x1b[0m\n      \x1b[38;2;255;121;198m,\x1b[0m\n     ]\x1b[38;2;255;121;198m,\x1b[0m\n     tx\x1b[38;2;255;121;198m: \x1b[0m[\n      \x1b[38;2;255;121;198m,\x1b[0m\n      \x1b[38;2;255;121;198m,\x1b[0m\n      \x1b[38;2;255;121;198m,\x1b[0m\n      \x1b[38;2;255;121;198m,\x1b[0m\n      \x1b[38;2;255;121;198m,\x1b[0m\n      \x1b[38;2;255;121;198m,\x1b[0m\n      \x1b[38;2;255;121;198m,\x1b[0m\n      \x1b[38;2;255;121;198m,\x1b[0m\n     ]\x1b[38;2;255;121;198m,\x1b[0m\n     extend\x1b[38;2;255;121;198m: \x1b[0m\x1b[38;2;135;255;215m\"<+08>-8\"\x1b[0m\x1b[38;2;255;121;198m,\x1b[0m\n     cacheStart\x1b[38;2;255;121;198m: \x1b[0m\x1b[38;2;239;228;161m9223372036854775807\x1b[0m\x1b[38;2;255;121;198m,\x1b[0m\n     cacheEnd\x1b[38;2;255;121;198m: \x1b[0m\x1b[38;2;239;228;161m9223372036854775807\x1b[0m\x1b[38;2;255;121;198m,\x1b[0m\n     cacheZone\x1b[38;2;255;121;198m: \x1b[0m\x1b[38;2;174;129;255mzone\x1b[0m {\n      name\x1b[38;2;255;121;198m: \x1b[0m\x1b[38;2;255;121;198m,\x1b[0m\n      offset\x1b[38;2;255;121;198m: \x1b[0m\x1b[38;2;255;121;198m,\x1b[0m\n      isDST\x1b[38;2;255;121;198m: \x1b[0m\x1b[38;2;255;121;198m,\x1b[0m\n     }\x1b[38;2;255;121;198m,\x1b[0m\n    }\x1b[38;2;255;121;198m,\x1b[0m\n   }\x1b[38;2;255;121;198m,\x1b[0m\n   Valid\x1b[38;2;255;121;198m: \x1b[0m\x1b[38;2;239;228;161mtrue\x1b[0m\x1b[38;2;255;121;198m,\x1b[0m\n  }\x1b[38;2;255;121;198m,\x1b[0m\n  Addresses\x1b[38;2;255;121;198m: \x1b[0m{\n   \x1b[38;2;239;228;161m2\x1b[0m\x1b[38;2;255;121;198m: \x1b[0m\x1b[38;2;174;129;255maddress\x1b[0m {\n    PostalCode\x1b[38;2;255;121;198m: \x1b[0m\x1b[38;2;239;228;161m876\x1b[0m\x1b[38;2;255;121;198m,\x1b[0m\n   }\x1b[38;2;255;121;198m,\x1b[0m\n  }\x1b[38;2;255;121;198m,\x1b[0m\n  vehicles\x1b[38;2;255;121;198m: \x1b[0m[\n  ]\x1b[38;2;255;121;198m,\x1b[0m\n  NilField\x1b[38;2;255;121;198m: \x1b[0mnil\x1b[38;2;255;121;198m,\x1b[0m\n }\x1b[38;2;255;121;198m,\x1b[0m\n]"
		if got != expected {
			t.Errorf("Expect: %s, but got: %s", expected, got)
		}
	})
}
