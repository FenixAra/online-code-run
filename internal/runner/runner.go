package runner

import (
	"github.com/FenixAra/online-code-run/dtos"
	"github.com/FenixAra/online-code-run/internal/config"
)

type Runner interface {
	Run(req *dtos.APIReq) *dtos.APIRes
	Compile(req *dtos.APIReq) *dtos.APIRes
}

func New(lang string) Runner {
	switch lang {
	case config.Golang:
		return NewGoRunner()
	case config.Java:
		return NewJavaRunner()
	case config.C:
		return NewCRunner()
	case config.CPP:
		return NewCPPRunner()
	case config.Python:
		return NewPythonRunner()
	case config.Python3:
		return NewPython3Runner()
	case config.Nodejs:
		return NewNodejsRunner()
	case config.Ruby:
		return NewRubyRunner()
	default:
		return nil
	}
}
