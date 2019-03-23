package dtos

type APIReq struct {
	ID        string
	Language  string   `json:"language"`
	Source    string   `json:"source"`
	Name      string   `json:"name"`
	ShouldRun bool     `json:"should_run"`
	Inputs    []string `json:"inputs"`
}

type APIRes struct {
	Output    string  `json:"output"`
	Error     string  `json:"error"`
	Status    bool    `json:"status"`
	TimeTaken float64 `json:"time_taken"`
}
