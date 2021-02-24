package main

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	Name    string      `json:"name"`
	Data    interface{} `json:"data"`
	Version int         `json:"version"`
}

func main() {
	message := Message{
		Name:    "Test1",
		Data:    "some random string",
		Version: 1,
	}

	var m Message

	msgJSON, err := json.Marshal(message)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("msgJSON:", string(msgJSON))
	json.Unmarshal(msgJSON, &m)
	fmt.Println("the unmarshaled message is:", m)
	fmt.Println("name is", m.Name)
	fmt.Println("data is", m.Data)
	fmt.Println("version is", m.Version)

	message = Message{
		Name:    "Test2",
		Data:    123,
		Version: 1,
	}

	msgJSON, err = json.Marshal(&message) // it doesn't make a difference whether you pass in a pointer or not
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("msgJSON:", string(msgJSON))
	json.Unmarshal(msgJSON, &m)
	fmt.Println("the unmarshaled message is:", m)
	fmt.Println("name is", m.Name)
	fmt.Println("data is", m.Data)
	fmt.Println("version is", m.Version)
}
