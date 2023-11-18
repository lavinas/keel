package main

import (
	"encoding/json"
	"fmt"
)

const (
	entry = `{
		"name": "John",
		"age": 30,
		"cars": [
			{ "name": "Ford", "models": ["Fiesta", "Focus", "Mustang"] },
			{ "name": "BMW", "models": ["320", "X3", "X5"] },
			{ "name": "Fiat", "models": ["500", "Panda"] }
		]
	}`
)

type Car struct {
	Name   string   `json:"name"`
	Models []string `json:"models"`
}

type Entry struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Cars []Car  `json:"cars"`
}

func main() {
	var model Entry
	err := json.Unmarshal([]byte(entry), &model)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", model)
}
