package main

import (
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

	// fmt.Println(string(res))
	f, err := os.Create("data.txt")
	if err != nil {
		fmt.Println("failed fetching", err)
	}
	defer f.Close()

	_, err = f.WriteString(string(res))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("successfuly written string")
}
