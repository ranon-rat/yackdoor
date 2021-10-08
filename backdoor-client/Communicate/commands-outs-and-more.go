package Communicate

import (
	"bytes"
	"os/exec"
)

const (
	Id     = "cum"
	Origin = "http://localhost:8080/commands?id=" + Id
	Url    = "ws://localhost:8080/commands?id=" + Id
)

var (
	commandExecChan = make(chan string)
	secondCommand   = make(chan string)
	b               bytes.Buffer
	cmd             *exec.Cmd
	exited          = make(chan bool)
)
