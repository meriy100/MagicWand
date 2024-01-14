package terraforms

import (
	"github.com/cockroachdb/errors"
	"os"
	"path/filepath"
	"text/template"
)

func createFile(moduleDir, fileName string, tempStr string, data any) error {
	tmpl, err := template.New(fileName).Parse(tempStr)
	if err != nil {
		return errors.Wrapf(err, "failed to parse template")
	}

	file, err := os.Create(filepath.Join("terraform", moduleDir, fileName))
	if err != nil {
		return errors.Wrapf(err, "failed to create file")
	}
	if err := tmpl.Execute(file, data); err != nil {
		return errors.Wrapf(err, "failed to execute template")
	}
	return nil
}
