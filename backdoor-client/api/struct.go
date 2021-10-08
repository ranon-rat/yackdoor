package api

type ApiOutput struct {
	ForWho string `json:"for_who"`
	Output string `json:"output"`

	Exited bool `json:"exited"`
}
type ApiCommand struct {
	From    string `json:"from"`
	Command string `json:"command"`
}
