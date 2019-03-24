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

type Python3Runner struct {
}

func NewPython3Runner() *Python3Runner {
	return &Python3Runner{}
}

func (py *Python3Runner) Run(req *dtos.APIReq) *dtos.APIRes {
	res := &dtos.APIRes{}
	p := "/tmp/ocr/" + req.ID
	err := os.MkdirAll(p, os.ModePerm)
	if err != nil {
		res.Error = err.Error()
		return res
	}

	err = ioutil.WriteFile(p+"/"+req.Name, []byte(req.Source), 0644)
	if err != nil {
		res.Error = err.Error()
		return res
	}

	cmd := exec.Command("python3", p+"/"+req.Name)
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
			errs := strings.Replace(string(out), p+"/"+req.Name, req.Name, -1)
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
	err = os.RemoveAll(p)
	if err != nil {
		res.Error = err.Error()
		return res
	}

	return res
}

func (py *Python3Runner) Compile(req *dtos.APIReq) *dtos.APIRes {
	return nil
}
