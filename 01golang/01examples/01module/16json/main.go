package main

import (
	"encoding/json"
	"fmt"
)

type Human struct {
	Name string
	Age  int
}

func main() {
	h := Human{Name: "cjx"}
	jsonStr := marshal2JsonString(h)
	fmt.Println(jsonStr)

	h2 := unmarshal2Struct(jsonStr)
	fmt.Println(h2)
}

func marshal2JsonString(h Human) string {
	h.Age = 30
	updatedBytes, err := json.Marshal(&h)
	if err != nil {
		fmt.Println(err)
	}
	return string(updatedBytes)
}

func unmarshal2Struct(HumanStr string) Human {
	h := Human{}
	err := json.Unmarshal([]byte(HumanStr), &h)
	if err != nil {
		fmt.Println(err)
	}
	return h
}
