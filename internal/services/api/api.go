package api

import (
	"github.com/FenixAra/online-code-run/dtos"
	"github.com/FenixAra/online-code-run/internal/config"
)

func CompileCode() {

}

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
