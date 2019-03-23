package api

import (
	"github.com/FenixAra/online-code-run/dtos"
	"github.com/FenixAra/online-code-run/internal/config"
	"github.com/FenixAra/online-code-run/internal/runner"
	uuid "github.com/satori/go.uuid"
)

func CompileAndRunCode(req *dtos.APIReq) *dtos.APIRes {
	r := runner.New(req.Language)
	id, _ := uuid.NewV4()
	req.ID = id.String()

	return r.Run(req)
}

// API to get all the languages supported
func GetLanguages() []dtos.Language {
	var l []dtos.Language
	for k, v := range config.LanguageVersion {
		l = append(l, dtos.Language{
			Name:    k,
			Version: v,
		})
	}

	return l
}
