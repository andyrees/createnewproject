/*
Copyright 2014 Andrew Rees.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
either express or implied. See the License for the specific
language governing permissions and limitations under the
License.
*/

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

type Project struct {
	namespace     string
	projectName   string
	mainOrPackage int
}

func checkerror(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (p *Project) getProjectDetails() {
	fmt.Println("Namespace: ")
	reader := bufio.NewReader(os.Stdin)
	input, _, err := reader.ReadLine()
	checkerror(err)
	p.namespace = fmt.Sprintf("%s", input)

	fmt.Println("Project Name: ")
	reader2 := bufio.NewReader(os.Stdin)
	input2, _, err2 := reader2.ReadLine()
	checkerror(err2)
	p.projectName = fmt.Sprintf("%s", input2)

	fmt.Println("Main / Package / Both ( M/P/B )")
	reader3 := bufio.NewReader(os.Stdin)
	input3, _, err3 := reader3.ReadLine()
	checkerror(err3)

	if strings.Contains(strings.ToLower(string(input3)), "m") {
		p.mainOrPackage = 1
	} else if strings.Contains(strings.ToLower(string(input3)), "p") {
		p.mainOrPackage = 2
	} else {
		p.mainOrPackage = 3
	}
}

func getGoPath() string {
	p := ""
	for _, value := range os.Environ() {
		if strings.Contains(strings.ToUpper(value), "GOPATH") {
			p = strings.Split(value, "=")[1]
			return p
		}
	}
	return p

}

func (p *Project) createProjectStructure() {
	gopath := getGoPath()
	if len(gopath) < 1 {
		fmt.Println("GOPATH NOT FOUND")
		os.Exit(1)
	}
	// Create Namespace & Project Folder
	projectPath := path.Join(gopath, "src", p.namespace, p.projectName)
	err := os.MkdirAll(projectPath, 0755)
	checkerror(err)
	// create stub file

	stubMain := `package main

import "fmt"

func main() {
	fmt.Println("Hello, golang")
}`

	stubPackage := fmt.Sprintf(`package %s

import (
	"fmt"
)
`, p.projectName)

	switch p.mainOrPackage {
	case 1:
		filepath := path.Join(projectPath, "main.go")
		err2 := ioutil.WriteFile(filepath, []byte(stubMain), 0644)
		checkerror(err2)
	case 2:
		filepath2 := path.Join(projectPath, fmt.Sprintf("%s.go", p.projectName))
		err3 := ioutil.WriteFile(filepath2, []byte(stubPackage), 0644)
		checkerror(err3)
	case 3:
		filepath := path.Join(projectPath, "main.go")
		err2 := ioutil.WriteFile(filepath, []byte(stubMain), 0644)
		checkerror(err2)

		filepath2 := path.Join(projectPath, fmt.Sprintf("%s.go", p.projectName))
		err3 := ioutil.WriteFile(filepath2, []byte(stubPackage), 0644)
		checkerror(err3)
	}
}

func main() {
	p := new(Project)
	p.getProjectDetails()
	p.createProjectStructure()
	fmt.Println("The project Stub has been created")
}
