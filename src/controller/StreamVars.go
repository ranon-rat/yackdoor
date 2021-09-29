package controller

var (
	commands = make(map[string]chan string)
)

/*
const (
	S            = 600
	boundaryWord = "CUMMASTER"
	headerf      = "\r\n" +
		"--" + boundaryWord + "\r\n" +
		"Content-Type: text/plain\r\n" +
		"X-Timestamp: 0.000000\r\n" +
		"\r\n"
)*/
