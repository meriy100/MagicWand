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
	"io"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello, world!")
	handler1 := func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s\n", r.Method, r.URL.Path)
		io.WriteString(w, "Hello-1\n")
	}
	handler2 := func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s\n", r.Method, r.URL.Path)
		io.WriteString(w, "Hello-2\n")
	}

	http.HandleFunc("/foo", handler1)
	http.HandleFunc("/bar", handler2)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
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
