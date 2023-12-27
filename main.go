package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

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
