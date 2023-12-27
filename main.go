package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type ytType struct {
	Title            string `json:"title"`
	Author_name      string `json:"author_name"`
	Author_url       string `json:"author_url"`
	Type             string `json:"type"`
	Height           string `json:"height"`
	Width            string `json:"width"`
	Version          string `json:"version"`
	Provider_name    string `json:"provider_name"`
	Provider_url     string `json:"provider_url"`
	Thumbnail_height string `json:"thumbnail_height"`
	Thumbnail_width  string `json:"thumbnail_width"`
	Thumbnail_url    string `json:"thumbnail_url"`
	Html             string `json:"html"`
}

func main() {
	data, err := http.Get("https://www.youtube.com/oembed?url=" + "https://www.youtube.com/watch?v=uvb-1wjAtk4" + "&format=json")
	if err != nil {
		fmt.Println("failed fetching", err)
		return
	}

	res, err := io.ReadAll(data.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = saveJson(string(res))
	if err != nil {
		fmt.Println("save file problem")
		return
	}

	fmt.Println("successfuly written string")
}

func saveJson(data string) error {
	saveFile, err := os.Create("data.json")
	if err != nil {
		fmt.Println("error creating file: ", err)
		return err
	}

	defer saveFile.Close()

	encoder := json.NewEncoder(saveFile)
	err = encoder.Encode(data)
	if err != nil {
		fmt.Println("error encoding data: ", err)
		return err
	}

	return nil

}
