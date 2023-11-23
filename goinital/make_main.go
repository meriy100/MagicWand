package goinital

import (
	"fmt"
	"github.com/cockroachdb/errors"
	"github.com/meriy100/magicwand/entities"
	"html/template"
	"os"
	"path/filepath"
)

const mainTemplate = `
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, world!")
}
`

func (i *Interactor) CreateMain(appName string, appType entities.AppType) error {
	fmt.Printf("Create cmd/%s/main.go\n", appName)
	fmt.Printf("%v\n", appType)
	if err := os.Mkdir("cmd", 0755); err != nil {
		return errors.Wrapf(err, "make directory failed")
	}

	if err := os.Mkdir(filepath.Join("cmd", appName), 0755); err != nil {
		return errors.Wrapf(err, "make directory failed")
	}

	tmpl, err := template.New("main.go").Parse(mainTemplate)
	if err != nil {
		return errors.Wrapf(err, "failed to parse template")
	}

	file, err := os.Create(filepath.Join("cmd", appName, "main.go"))
	if err != nil {
		return errors.Wrapf(err, "failed to create file")
	}
	if err := tmpl.Execute(file, nil); err != nil {
		return errors.Wrapf(err, "failed to execute template")
	}

	return nil
}
