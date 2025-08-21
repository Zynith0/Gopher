package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	var init string
	flag.StringVar(&init, "Project Name", "", "The name of the project to be made")

	flag.Parse()

	if len(flag.Args()) < 1 {
		os.Exit(1)
	}

	if flag.Arg(0) == "init" {
		var projectName = flag.Arg(1)

		cmd, err := exec.Command("/bin/sh", "/home/zynith/Scripts/GoInit.sh", projectName).Output()
		if err != nil {
			fmt.Println("skill issue", err)
		}
		fmt.Println(string(cmd))
	}
}
