package ncmd

import (
	"fmt"
	"log"
	"os/exec"
)

type NCmd struct {
	Out string
}

var std = New("")

func New(out string) *NCmd {
	return &NCmd{Out: out}
}

//run ssh command
func Execute(cmd string, args ...string) error {

	log.Println("log:", cmd, args)

	std.Out = ""

	out, err := exec.Command(cmd, args...).Output()

	if err != nil {
		log.Println("Error running:", err.Error())
		return err
	}

	std.Out = fmt.Sprintf("%s", out)

	return nil
}
