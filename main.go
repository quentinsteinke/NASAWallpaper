package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	apiKey := "xOZSHBa1Otmj8PMPAIfxMDmUpFc9JOm2eaFafTD0"
	url := "https://api.nasa.gov/planetary/apod?api_key="
	count := 1

	requestUrl := fmt.Sprintf("%s%s&count=%d", url, apiKey, count)
	response, err := http.Get(requestUrl)
	if err != nil {
		fmt.Printf("Couldn't get response")
	}
	response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Couldn't read response body")
	}

	fmt.Printf("Response: %s", body)
}
