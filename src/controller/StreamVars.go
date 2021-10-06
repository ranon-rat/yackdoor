package controller

var (
	commands = make(map[string]chan string)
	outputs  = make(map[string]chan string)
)
