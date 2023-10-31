package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type data struct {
	Url string `json:"url"`
}

func main() {
	NASAUrl := "https://api.nasa.gov/planetary/apod?api_key="
	apiKey := "xOZSHBa1Otmj8PMPAIfxMDmUpFc9JOm2eaFafTD0"
	count := 1

	var imageUrl []data

	requestUrl := fmt.Sprintf("%s%s&count=%d", NASAUrl, apiKey, count)
	response, err := http.Get(requestUrl)
	if err != nil {
		fmt.Printf("Couldn't get response")
		return
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Couldn't read response body")
		return
	}
	response.Body.Close()

	// fmt.Printf("Response: %s", body)

	err = json.Unmarshal([]byte(body), &imageUrl)
	if err != nil {
		fmt.Printf("Unable to Unmarshal body: %v", err)
		return
	}

	fmt.Printf("Body: %s", imageUrl[0].Url)

	err = downloadImage(imageUrl[0].Url)
	if err != nil {
		fmt.Printf("Unable to download image: %v", err)
	}
}

func downloadImage(imageUrl string) error {
	response, err := http.Get(imageUrl)
	if err != nil {
		fmt.Printf("Couldn't get response from url: %v", err)
		return err
	}

	imageName, err = getFileNameFromUrl(imageUrl)
	if err != nil {
		fmt.Printf("Unable to get image name from url: %v", err)
	}
	return nil
}

func getFileNameFromUrl(imageUrl string) (string, error) {
	fileUrl, err := url.Parse(imageUrl)
	if err != nil {
		fmt.Printf("Unagle to parse ImageUrl: %v", err)
		return "", err
	}

	path := fileUrl.Path
	segment := strings.Split(path, "/")

	imageName := segment[len(segment)-1]

	return fmt.Sprint(imageName), nil
}
