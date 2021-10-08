package api

type ApiOutput struct {
	ForWho   string `json:"for_who"`
	Output   string `json:"output"`
	Exited   bool   `json:"exited"`
}
