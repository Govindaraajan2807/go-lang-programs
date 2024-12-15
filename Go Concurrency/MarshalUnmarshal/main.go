package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {

	//  MarshallToJson
	p := Person{Name: "Govind", Age: 26}
	value, err := marshallToJson(p)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Println(value)
	}

	person, err := unmarshallFromJson(value)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Println(person)
	}

}

func marshallToJson(p Person) (string, error) {

	pBytes, err := json.Marshal(p)
	return string(pBytes), err
}

func unmarshallFromJson(jsonString string) (Person, error) {

	var person Person

	err := json.Unmarshal([]byte(jsonString), &person)
	if err != nil {
		return Person{}, fmt.Errorf("Failed to unmarshall %v", err)
	}
	return person, err
}
