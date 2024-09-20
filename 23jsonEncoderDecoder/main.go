package main

import (
	"encoding/json"
	"fmt"
)

type Course struct {
	Name     string   `json:"courseName"`
	Price    int      `json:"price"`
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {

	fmt.Println("Welcome to JSON ENCODER DECODER")
	//Encoder()
	Decoder()
}

func Encoder() {
	lcoCourses := []Course{
		{"MERN Stack", 299, "abhi.com", "abc@1234", []string{"React", "Nodejs", "js", "mongo"}},
		{"ReactJs Bootcamp", 199, "lco.com", "ash@123", []string{"react"}},
		{"Golang Bootcamp", 499, "golang.org", "go@123", nil},
	}

	//js, err := json.Marshal(lcoCourses)

	js, err := json.MarshalIndent(lcoCourses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("Json Data is: %s\n", js)
}

func Decoder() {
	jsonFromWeb := []byte(`
	   {
            "courseName": "MERN Stack",
            "price": 299,
            "website": "abhi.com",
            "tags": ["React","Nodejs","js","mongo"]
        }
	`)

	var decodedJson Course

	isValidJson := json.Valid(jsonFromWeb)

	if isValidJson {
		json.Unmarshal(jsonFromWeb, &decodedJson)
		fmt.Printf("Decoded json: %#v\n", decodedJson)
	} else {
		fmt.Println("JSON WAS INVALID")
	}

	var myonlineData map[string]interface{}

	json.Unmarshal(jsonFromWeb, &myonlineData)

	for k, v := range myonlineData {
		fmt.Printf("Key is %v and value is %v\n", k, v)
	}
}
