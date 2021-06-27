# pp

pp is a simple printer for Golang.

```
import "github.com/hlcfan/pp"

m := map[string]string{"foo": "bar", "hello": "world"}
pp.Print(m)

    map[string]string {
      foo: 	bar,
      hello: 	world,
    }


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

pp.Inspect(people)

    []pp_test.person {
      {
        ID:	1,
        Name:	alex,
        Phone:	12345678,
        Graduated:	true,
        CreatedAt:{
          Time:	2009-11-17 20:34:58.651387237 +0000 UTC,
          Valid:	true,
        },
        Addresses: map[int]pp_test.address {
          1: {
            PostalCode:	123,
          },
        },
        vehicles:[]pp_test.vehicle {
          {
            plate:	CA-1234,
          },
        },
      },
      {
        ID:	2,
        Name:	bob,
        Phone:	87654321,
        Graduated:	false,
        CreatedAt:{
          Time:	2021-06-05 20:34:58.651387237 +0800 +08,
          Valid:	true,
        },
        Addresses: map[int]pp_test.address {
          2: {
            PostalCode:	876,
          },
        },
        vehicles:[]pp_test.vehicle {
        },
      },
    }
```
