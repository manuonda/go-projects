package redisclient

import "encoding/json"

var persons string = `
[
  {Name: "Jhon", Age: 20},
  {Name: "Jane", Age: 21}
]
`

var employees string = `
[
	{Name: "Jhon", Position : "Manager"},
	{Name: "Jane", Position : "Developer"}
]
`

var db map[string]string = map[string]string{
	"persons":   persons,
	"employees": employees,
}

func Read[T any](key string) (T, error) {
	var data T

	result := db[key]
	error := json.Unmarshal([]byte(result), &data)
	if error != nil {
		return data, error
	}
	return data, nil
}
