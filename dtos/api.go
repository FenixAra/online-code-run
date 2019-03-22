package dtos

type APIReq struct {
	Language  string `json:"language"`
	Source    string `json:"source"`
	ShouldRun bool   `should_run`
}
