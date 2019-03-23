package runner

import (
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/FenixAra/online-code-run/dtos"
)

type GoRunner struct {
}

func NewGoRunner() *GoRunner {
	return &GoRunner{}
}

func (g *GoRunner) Run(req *dtos.APIReq) *dtos.APIRes {
	res := &dtos.APIRes{}
	err := os.MkdirAll("/tmp/ocr/"+req.ID, os.ModePerm)
	if err != nil {
		res.Error = err.Error()
		return res
	}

	p := "/tmp/ocr/" + req.ID + "/main.go"
	err = ioutil.WriteFile(p, []byte(req.Source), 0644)
	if err != nil {
		res.Error = err.Error()
		return res
	}

	cmd := exec.Command("go", "run", p)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		res.Error = err.Error()
		return res
	}

	io.WriteString(stdin, strings.Join(req.Inputs, "\n"))
	stdin.Close()

	t := time.Now()
	out, err := cmd.CombinedOutput()
	if err != nil {
		if _, ok := err.(*exec.ExitError); ok {
			errs := strings.Replace(string(out), p, req.Name, -1)
			errs = strings.Replace(errs, "# command-line-arguments\n", "", 1)

			res.Error = errs
			return res
		}

		res.Error = err.Error()
		return res
	}

	res.Status = true
	res.Output = string(out)
	res.TimeTaken = time.Since(t).Seconds()
	os.RemoveAll(p)
	return res
}

func (g *GoRunner) Compile(req *dtos.APIReq) *dtos.APIRes {
	return nil
}
