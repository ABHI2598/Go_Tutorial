package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {

	response, err := http.Get("https://dummyjson.com/products")

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	var responseString strings.Builder

	databytes, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	totalBytes, _ := responseString.Write(databytes)
	fmt.Println("Total Bytes: ", totalBytes)
	fmt.Println("Content is: ", responseString.String())

	content := string(databytes)

	fmt.Printf("Response for the given url is....: %v\n", content)
	fmt.Printf("Header of the request is.....: %v\n", response.Header)
	fmt.Printf("Status of the URL is.....: %v\n", response.Status)
	fmt.Printf("Status Code of the URL is.....: %v\n", response.StatusCode)
	fmt.Printf("Proto of the response....: %v\n", response.Proto)
	fmt.Printf("Content Length of the response.....: %d\n", response.ContentLength)

}
