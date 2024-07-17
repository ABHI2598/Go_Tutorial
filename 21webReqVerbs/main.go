package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Printf("Welcome to HTTP REQUESTS \n")

	const apiUrl = "https://{key}.free.beeceptor.com/api/users/"
	PerformGetRequest(apiUrl)
	//PerformPostRequest(apiUrl)

	//PerformFormMethodPost(apiUrl)

}

func PerformGetRequest(url string) {

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	var responseString strings.Builder

	dataBytes, err := io.ReadAll(response.Body)

	responseString.Write(dataBytes)

	if err != nil {
		panic(err)
	}

	fmt.Println(responseString.String())
}

func PerformPostRequest(url string) {

	requestBody := strings.NewReader(`
	    {"data":"Hello Abhinav"}
	
	`)

	response, err := http.Post(url, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	var responseString strings.Builder

	dataBytes, err := io.ReadAll(response.Body)

	responseString.Write(dataBytes)

	if err != nil {
		panic(err)
	}

	fmt.Println(responseString.String())
}

func PerformFormMethodPost(apiUrl string) {

	formData := url.Values{}

	formData.Add("name", "Abhinav is Awesomeee")

	response, err := http.PostForm(apiUrl, formData)

	if err != nil {
		panic(err)
	}

	var responseString strings.Builder

	dataBytes, err := io.ReadAll(response.Body)

	responseString.Write(dataBytes)

	if err != nil {
		panic(err)
	}

	fmt.Println(responseString.String())
}
