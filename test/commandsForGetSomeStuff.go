package main

import (
	"bytes"
	"fmt"
)

/*
im testing this for
receive a json from the server and send a
response to the server with a json with some stuff and execute the command and get the output

just a simple stdin test for the os/exec
*/
func main() {
	stdin := new(bytes.Buffer)
	stdin.Write([]byte("hello world"))
	fmt.Println(stdin.String())
}
