package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type File struct {
	fileName string
	content  string
	open     bool
}

func main() {
	if !rootCheck() {
		log.Fatal("Run from root")
	}

	day := getQuest()
	fileList := []File{
		{fileName: "example_I.txt", content: "", open: true},
		{fileName: "example_II.txt", content: "", open: false},
		{fileName: "example_III.txt", content: "", open: false},
		{fileName: "input_I.txt", content: "", open: true},
		{fileName: "input_II.txt", content: "", open: false},
		{fileName: "input_III.txt", content: "", open: false},
		{fileName: fmt.Sprintf("%s_suite_test.go", day), content: loadTemplate("cmd/generate/testSuite.tmpl", day), open: false},
		{fileName: fmt.Sprintf("%s_test.go", day), content: loadTemplate("cmd/generate/test.tmpl", day), open: true},
		{fileName: fmt.Sprintf("%s_logic.go", day), content: loadTemplate("cmd/generate/logic.tmpl", day), open: true},
		{fileName: fmt.Sprintf("%s.go", day), content: loadTemplate("cmd/generate/solver.tmpl", day), open: true},
	}

	createFolder(day)
	createFiles(fileList)
}
func getQuest() string {
	day := flag.String("quest", "", "Which quest are you creating?")
	flag.Parse()

	if *day == "" {
		log.Fatal("error: -quest flag is required")
	}
	return *day
}
func createFolder(day string) {
	//Make the folder
	err := os.Mkdir(day, 0755)
	if err != nil {
		log.Fatalf("Failed to create quest folder: %v", err)
	}

	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get current working directory: %v", err)
	} else {
		fmt.Printf("%s folder created at %s\n", day, dir)
	}

	//Move into folder
	err = os.Chdir(day)
	if err != nil {
		log.Fatalf("Failed to change directory: %v", err)
	}

	//Create data folder
	err = os.Mkdir("inputs", 0755)
	if err != nil {
		log.Fatalf("Failed to create data folder: %v", err)
	}

	dir, err = os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get current working directory: %v", err)
	} else {
		fmt.Printf("%s folder created at %s\n", "inputs", dir)
	}

}
func createFiles(files []File) {
	for _, context := range files {
		filename := context.fileName

		if strings.HasSuffix(filename, ".txt") {
			filename = filepath.Join("inputs", filename)
		}

		//Create file
		file, err := os.Create(filename)
		if err != nil {
			log.Fatalf("Failed to create file: %v", err)
		}
		//Write contents
		_, err = file.WriteString(context.content)
		if err != nil {
			log.Fatalf("Failed to write to file: %v", err)
		}
		fmt.Println(filename, "created")

		if context.open {
			openFile(filename)
		}

		file.Close()
	}
}
func loadTemplate(template, day string) string {
	data, err := os.ReadFile(template)
	if err != nil {
		log.Fatalf("Failed to read template: %v", err)
	}
	return strings.ReplaceAll(string(data), "$", day)
}
func rootCheck() bool {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get current working directory: %v", err)
	}

	if files, err := os.ReadDir(dir); err != nil {
		log.Fatalf("Failed to read filed in %s: %v", dir, err)
	} else {
		for _, file := range files {
			if file.Name() == "go.mod" {
				return true
			}
		}
	}
	return false
}
func openFile(file string) {
	cmd := exec.Command("code", file)
	err := cmd.Start()
	if err != nil {
		log.Printf("Failed to open %s in VS Code: %v", file, err)
	}

}
