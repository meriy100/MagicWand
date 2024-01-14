package terraforms

import (
	"github.com/cockroachdb/errors"
	"github.com/meriy100/magicwand/entities"
	"os"
)

type Backend struct {
}

func NewBackend() *Backend {
	return &Backend{}
}

func (b *Backend) Create(gcpConfig entities.GCPConfig) error {
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

	if err := os.Mkdir("terraform/backend", 0755); err != nil {
		return errors.Wrapf(err, "make directory failed")
	}

	return createFile("backend", "main.tf", mainTmpl, gcpConfig)
}
