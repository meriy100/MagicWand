package terraforms

import (
	"github.com/cockroachdb/errors"
	"github.com/meriy100/magicwand/entities"
	"os"
	"path/filepath"
	"text/template"
)

type Backend struct {
}

const mainTmpl = `
provider "google" {
  project = "{{.ProjectID}}"
  region  = "asia-northeast1"
  zone    = "asia-northeast1-a"
}

resource "google_storage_bucket" "backend" {
  name          = "{{.ProjectID}}-terraform-state"
  location      = "asia-northeast1"
  force_destroy = false
  storage_class = "NEARLINE"
}
`

func NewBackend() *Backend {
	return &Backend{}
}

func (b *Backend) Create(gcpConfig entities.GCPConfig) error {
	if err := os.Mkdir("terraform/backend", 0755); err != nil {
		return errors.Wrapf(err, "make directory failed")
	}

	tmpl, err := template.New("main.tf").Parse(mainTmpl)
	if err != nil {
		return errors.Wrapf(err, "failed to parse template")
	}

	file, err := os.Create(filepath.Join("terraform", "backend", "main.tf"))
	if err != nil {
		return errors.Wrapf(err, "failed to create file")
	}
	if err := tmpl.Execute(file, gcpConfig); err != nil {
		return errors.Wrapf(err, "failed to execute template")
	}
	return nil
}
