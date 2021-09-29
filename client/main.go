package main

import (
	"bufio"
	"log"
	"net/http"
)

const serverAddr = "localhost:8080"

func main() {

	// HTTP client
	req, err := http.NewRequest("GET", "http://"+serverAddr+"/commands", nil)
	if err != nil {
		log.Fatal("Error creating HTTP request: ", err.Error())
	}

	req.Header.Set("id", "cum")
	client := http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error making HTTP request: ", err.Error())
	}

	// Read the response header

	// Read the response body

	reader := bufio.NewReader(resp.Body)
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			break
		}

		log.Println(string(line))
	}

}
