package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type Person struct {
		Name string `json:"name"`
		Id   int    `json:"id"`
	}
	person := Person{"Bapan", 21}
	fmt.Println(person)
	data, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error Occured:")
		return
	}
	fmt.Println(string(data))
	mydata := json.RawMessage(string(data))
	newdata, err := mydata.MarshalJSON()
	if err != nil {
		fmt.Println("Error occured:")
		return
	}
	p := Person{}
	err = json.Unmarshal(newdata, &p)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(p)
	m := make(map[string]interface{})
	err = json.Unmarshal(newdata, &m)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(m["name"].(string))
	fmt.Println(m["id"].(float64))
}
