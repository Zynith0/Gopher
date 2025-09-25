package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	var init string
	flag.StringVar(&init, "Project Name", "", "The name of the project to be made")

	flag.Parse()

	// Print a help message if the user doesn't input any flags
	if len(flag.Args()) < 1 {
		fmt.Println("The current commands are:")
		fmt.Println("init       initializes a go project")
		fmt.Println("init-src   initializes a go project using src instead of cmd")
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
	} else if flag.Arg(0) == "init-src" {
		var projectName = flag.Arg(1)

		initProjectSrc(projectName)
	} else if flag.Arg(0) == "run-src" {
		runSrc()
	} else if flag.Arg(0) == "build-src" {
		buildSrc()
	} else if flag.Arg(0) == "init-web" {
		if len(flag.Args()) < 2 {
			fmt.Println("You need to specify a project name!")
		} else {
			var projectName = flag.Arg(1)

			initWebProject(projectName)
		}
	}
}

func initProject(projectName string) {
	os.Mkdir(projectName, 0755)
	os.Mkdir(projectName + "/cmd", 0755)
	mainGo, err := os.Create(projectName  + "/cmd/" + "main.go")

	mainGo.WriteString("package main\n\nimport (\n	" + `"` + "fmt" + `"` + "\n)\n\nfunc main() {\n	fmt.Println(" + `"` + "Hello, World!" + `"` + ")\n}")

	userName, err := exec.Command("git", "config", "--global", "user.name").Output()
	if err != nil {
		fmt.Println("skill issue", err)
	}

	userNameOut := string(userName)

	cmd := exec.Command("go", "mod", "init", "github.com/" + strings.Trim(userNameOut, "\n") + "/" +  projectName)
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

func initProjectSrc(projectName string) {
	os.Mkdir(projectName, 0755)
	os.Mkdir(projectName + "/src", 0755)
	mainGo, err := os.Create(projectName  + "/src/" + "main.go")

	mainGo.WriteString("package main\n\nimport (\n	" + `"` + "fmt" + `"` + "\n)\n\nfunc main() {\n	fmt.Println(" + `"` + "Hello, World!" + `"` + ")\n}")

	userName, err := exec.Command("git", "config", "--global", "user.name").Output()
	if err != nil {
		fmt.Println("skill issue", err)
	}

	userNameOut := string(userName)

	cmd := exec.Command("go", "mod", "init", "github.com/" + strings.Trim(userNameOut, "\n") + "/" +  projectName)
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

func runSrc() {
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

func buildSrc() {
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

func initWebProject(projectName string) {
	os.Mkdir(projectName, 0755)
	os.Mkdir(projectName + "/cmd", 0755)
	os.Mkdir(projectName + "/internal", 0755)
	os.Mkdir(projectName + "/internal/handler", 0755)
	os.Mkdir(projectName + "/view", 0755)
	mainGo, err := os.Create(projectName  + "/cmd/" + "main.go")
	handlerGo, err := os.Create(projectName  + "/internal/handler/" + "handler.go")
	_, err = os.Create(projectName  + "/view/" + "index.html")

	userName, err := exec.Command("git", "config", "--global", "user.name").Output()
	if err != nil {
		fmt.Println("skill issue", err)
	}

	userNameOut := string(userName)

	mainGo.WriteString("package main\n\nimport (\n	" + `"` + "net/http" + `"` + "\n	" + `"` + "github.com/" + strings.Trim(userNameOut, "\n") + "/" + projectName + "/internal/handler" + `"` + "\n)\n\nfunc main() {\n	mux := http.NewServeMux()\n	mux.HandleFunc(" + `"` + "/" + `"` + ", handler.HandleRoot)\n\n	http.ListenAndServe(" + `"` + ":8080" + `"` + ", mux)\n}")

	handlerGo.WriteString("package handler\n\nimport (\n	" + `"` + "net/http" + `"` + "\n	" + `"` + "html/template" + `"` + "\n	" + `"` + "fmt" + `"` + "\n)\n\nfunc HandleRoot(w http.ResponseWriter, r *http.Request) {\n	renderTemplate(" + `"` + "view/index.html" + `"` + ", w" + ")\n}\n\nfunc renderTemplate(tmpl string, w http.ResponseWriter) {\n	t, err := template.ParseFiles(tmpl)\n	if err != nil {\n		fmt.Println(" + `"` + "skill issue" + `"` + ", err)\n	}\n	t.Execute(w, " + `""` + ")\n}")


	cmd := exec.Command("go", "mod", "init", "github.com/" + strings.Trim(userNameOut, "\n") + "/" +  projectName)
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
