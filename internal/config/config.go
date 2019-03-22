package config

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

const (
	Java    = "Java"
	Golang  = "Go"
	C       = "C"
	CPP     = "C++"
	Python  = "Python"
	Python3 = "Python3"
	Nodejs  = "Node.js"
	Ruby    = "Ruby"
)

var LanguageVersion map[string]string
var Port string

func init() {
	Port = os.Getenv("Port")
	if Port == "" {
		Port = "3000"
	}

	LanguageVersion = make(map[string]string)
	// Get golang version
	cmd := exec.Command("go", "version")

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	strs := strings.Split(string(out), " ")
	if len(strs) > 2 {
		LanguageVersion[Golang] = strs[2]
	}

	// Get Java version
	cmd = exec.Command("java", "--version")

	out, err = cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	strs = strings.Split(string(out), " ")
	if len(strs) > 1 {
		LanguageVersion[Java] = strs[1]
	}

	// C version
	cmd = exec.Command("gcc", "--version")

	out, err = cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	strs = strings.Split(string(out), " ")
	if len(strs) > 3 {
		s := strings.Split(strs[3], "\n")
		LanguageVersion[C] = s[0]
	}

	// C++ version
	cmd = exec.Command("g++", "--version")

	out, err = cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	strs = strings.Split(string(out), " ")
	if len(strs) > 3 {
		s := strings.Split(strs[3], "\n")
		LanguageVersion[CPP] = s[0]
	}

	// Get Python version
	cmd = exec.Command("python", "--version")

	out, err = cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	strs = strings.Split(string(out), " ")
	if len(strs) > 1 {
		s := strings.Split(strs[1], "\n")
		LanguageVersion[Python] = s[0]
	}

	// Get Python3 version
	cmd = exec.Command("python3", "--version")

	out, err = cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	strs = strings.Split(string(out), " ")
	if len(strs) > 1 {
		s := strings.Split(strs[1], "\n")
		LanguageVersion[Python3] = s[0]
	}

	// Get Nodejs version
	cmd = exec.Command("nodejs", "--version")

	out, err = cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	strs = strings.Split(string(out), " ")
	if len(strs) > 0 {
		s := strings.Split(strs[0], "\n")
		LanguageVersion[Nodejs] = s[0]
	}

	// Get Ruby version
	cmd = exec.Command("ruby", "-v")

	out, err = cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	strs = strings.Split(string(out), " ")
	if len(strs) > 1 {
		LanguageVersion[Ruby] = strs[1]
	}

}
