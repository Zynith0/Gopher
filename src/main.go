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
		fmt.Println("The current commands are:")
		fmt.Println("init       initializes a go project")
	}

	if flag.Arg(0) == "init" {
		if len(flag.Args()) < 2 {
			fmt.Println("You need to specify a project name!")
		} else {
			var projectName = flag.Arg(1)

			runCmds(projectName)
		}
	}
}

func runCmds(projectName string) {
	os.Mkdir(projectName, 0755)
	os.Mkdir(projectName + "/src", 0755)
	os.Create(projectName  + "/src/" + "main.go")

	cmd := exec.Command("go", "mod", "init", "github.com/" + projectName)
	cmd.Dir = "./" + projectName
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("skill issue", err)
	}
	fmt.Println(string(output))

	gitInit := exec.Command("git", "init")
	gitInit.Dir = "./" + projectName
	gitInitOutput, err := gitInit.Output()
	if err != nil {
		fmt.Println("skill issue", err)
	}
	fmt.Println(string(gitInitOutput))
}
