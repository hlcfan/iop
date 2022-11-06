# pp

pp is a simple printer for Golang.

### Examples

#### Map

```
import "github.com/hlcfan/pp"

m := map[string]string{"foo": "bar", "hello": "world"}
pp.Puts(m)
```

*Output*
![screenshot](./screenshot2.png)

#### Complex data

```
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

pp.Puts(people)
```

*Output*
![screenshot](./screenshot1.png)
