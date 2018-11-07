package main

import (
	"fmt"
	"github.com/qeesung/asciiplayer/cmd"
	"os"
)

func main() {
	builder := cmd.CommandBuilder{}
	cli := builder.Build()
	if err := cli.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		if exitErr, ok := err.(cmd.ExitError); ok {
			os.Exit(exitErr.Code)
		} else {
			os.Exit(-1)
		}
	}
}
