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
        fmt.Println("run        runs the project (only works with src)")
        fmt.Println("build      builds the project (only works with src)")
    }

    if flag.Arg(0) == "init" {
        if len(flag.Args()) < 2 {
            fmt.Println("You need to specify a project name!")
        } else {
            var projectName = flag.Arg(1)

            runCmds(projectName)
        }
    }

	if flag.Arg(0) == "run" {
		run()
	}

	if flag.Arg(0) == "build" {
		build()
	}
}

func runCmds(projectName string) {
    os.Mkdir(projectName, 0755)
    os.Mkdir(projectName + "/src", 0755)
	mainGo, err := os.Create(projectName  + "/src/" + "main.go")

	mainGo.WriteString("package main\n\nimport (\n	" + `"` + "fmt" + `"` + "\n)\n\nfunc main() {\n	fmt.Println(" + `"` + "Hello, World!" + `"` + ")\n}")

    cmd := exec.Command("go", "mod", "init", "github.com/Zynith0" + projectName)
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

	f, err := os.Create(projectName + "/.gitignore")
	if err != nil {
		fmt.Println("skill issue", err)
	}
	f.WriteString("main")
}

func run() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("skill issue")
	}

	cmd := exec.Command("go", "run", "src/main.go")
	if err != nil {
		fmt.Println("skill issue", err)
	}
	cmd.Dir = dir

	output, err := cmd.Output()
	if err != nil {
		fmt.Println("skill issue", err)
	}
	fmt.Println(string(output))
}

func build() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("skill issue")
	}

	cmd := exec.Command("go", "build", "src/main.go")
	if err != nil {
		fmt.Println("skill issue", err)
	}
	cmd.Dir = dir

	output, err := cmd.Output()
	if err != nil {
		fmt.Println("skill issue", err)
	}
	fmt.Println(string(output))
}
