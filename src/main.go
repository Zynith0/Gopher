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

	// Print a help message if the user doesn't input any flags
    if len(flag.Args()) < 1 {
        fmt.Println("The current commands are:")
        fmt.Println("init       initializes a go project")
		// TODO: actually make init-src a thing
        // fmt.Println("init-src   initializes a go project using src instead of cmd")
        fmt.Println("run        runs the project")
        fmt.Println("build      builds the project")
    }

    if flag.Arg(0) == "init" {
        if len(flag.Args()) < 2 {
            fmt.Println("You need to specify a project name!")
        } else {
            var projectName = flag.Arg(1)

            initProject(projectName)
        }
    } else if flag.Arg(0) == "run" {
		run()
	} else if flag.Arg(0) == "build" {
		build()
	}
}

func initProject(projectName string) {
    os.Mkdir(projectName, 0755)
    os.Mkdir(projectName + "/cmd", 0755)
	mainGo, err := os.Create(projectName  + "/cmd/" + "main.go")

	mainGo.WriteString("package main\n\nimport (\n	" + `"` + "fmt" + `"` + "\n)\n\nfunc main() {\n	fmt.Println(" + `"` + "Hello, World!" + `"` + ")\n}")

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

	f, err := os.Create(projectName + "/.gitignore")
	if err != nil {
		fmt.Println("skill issue", err)
	}
	f.WriteString("main\n*.exe")
}

func run() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("skill issue")
	}

	cmd := exec.Command("go", "run", "cmd/main.go")
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

	cmd := exec.Command("go", "build", "cmd/main.go")
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
