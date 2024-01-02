package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
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
		log.Fatalf("Couldn't get response: %v", err)
		return
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Couldn't read response body: %v", err)
		return
	}
	response.Body.Close()

	// fmt.Printf("Response: %s", body)

	err = json.Unmarshal([]byte(body), &imageUrl)
	if err != nil {
		log.Fatalf("Unable to Unmarshal body: %v", err)
		return
	}

	fmt.Printf("Body: %s", imageUrl[0].Url)

	err = downloadImage(imageUrl[0].Url)
	if err != nil {
		log.Fatalf("Unable to download image: %v", err)
	}
}

func downloadImage(imageUrl string) error {
	response, err := http.Get(imageUrl)
	if err != nil {
		log.Fatalf("Couldn't get response from url: %v", err)
		return err
	}
	defer response.Body.Close()

	imageName, err := getFileNameFromUrl(imageUrl)
	if err != nil {
		log.Fatalf("Unable to get image name from url: %v", err)
	}

	file, err := os.Create(fmt.Sprintf("images/%s", imageName))
	if err != nil {
		log.Fatalf("Couldn't create file: %v", err)
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		log.Fatalf("Couldn't write file: %v", err)
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

// func setWallpaper(imagePath) error {
// 	syscall.LoadDLL()
// }
