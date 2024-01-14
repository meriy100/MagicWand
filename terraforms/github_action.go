package terraforms

import (
	"github.com/cockroachdb/errors"
	"github.com/meriy100/magicwand/entities"
	"os"
)

type GithubAction struct {
}

func NewGithubAction() *GithubAction {
	return &GithubAction{}
}

func (g *GithubAction) Create(projectConfig entities.ConfigSet, gcpConfig entities.GCPConfig) error {
	if err := os.Mkdir("terraform/github_action", 0755); err != nil {
		return errors.Wrapf(err, "make directory failed")
	}

	if err := g.createMain(gcpConfig); err != nil {
		return errors.Wrapf(err, "failed to create main.tf")
	}

	if err := g.createBackend(gcpConfig); err != nil {
		return errors.Wrapf(err, "failed to create backend.tf")
	}

	if err := g.createVariables(projectConfig); err != nil {
		return errors.Wrapf(err, "failed to create variables.tf")
	}

	if err := g.createProvider(gcpConfig); err != nil {
		return errors.Wrapf(err, "failed to create provider.tf")
	}

	return nil
}

func (g *GithubAction) createMain(gcpConfig entities.GCPConfig) error {
	const mainTmpl = `
resource "google_project_service" "iamcredentials" {
  service            = "iamcredentials.googleapis.com"
  disable_on_destroy = false
}

resource "google_service_account" "github_action" {
  account_id   = "github-action"
  display_name = "Github Action CI/CD"
}

resource "google_project_iam_binding" "project" {
  project = "{{.ProjectID}}}}"
  role    = "roles/artifactregistry.writer"

  members = [
    "serviceAccount:${google_service_account.github_action.email}",
  ]
}

resource "google_iam_workload_identity_pool" "github_action" {
  workload_identity_pool_id = "github-action"
  display_name              = "Github Action CI/CD"
  description               = "Github Action CI/CD"
}

resource "google_iam_workload_identity_pool_provider" "github" {
  workload_identity_pool_id          = google_iam_workload_identity_pool.github_action.workload_identity_pool_id
  workload_identity_pool_provider_id = "github"
  display_name                       = "Github"
  oidc {
    issuer_uri = "https://token.actions.githubusercontent.com"
  }

  attribute_mapping = {
    "google.subject"       = "assertion.sub"
    "attribute.repository" = "assertion.repository"
    "attribute.actor"      = "assertion.actor"
  }
}

resource "google_service_account_iam_member" "github_action" {
  service_account_id = google_service_account.github_action.id
  role               = "roles/iam.workloadIdentityUser"
  member             = "principalSet://iam.googleapis.com/${google_iam_workload_identity_pool.github_action.name}/attribute.repository/${var.github_repository}"
}
`
	return createFile("github_action", "main.tf", mainTmpl, gcpConfig)
}

func (g *GithubAction) createBackend(gcpConfig entities.GCPConfig) error {
	const backendTmpl = `
terraform {
  backend "gcs" {
    bucket = "{{.ProjectID}}-terraform-state"
    prefix = "terraform/github_actions"
  }
}
`
	return createFile("github_action", "backend.tf", backendTmpl, gcpConfig)
}

func (g *GithubAction) createVariables(projectConfig entities.ConfigSet) error {
	const variablesTmpl = `
variable "github_repository" {
  default = "{{.RepositoryOwner}}/{{.AppName}}"
  type    = string
}
`
	return createFile("github_action", "variables.tf", variablesTmpl, projectConfig)
}

func (g *GithubAction) createProvider(gcpConfig entities.GCPConfig) error {
	const providerTmpl = `
provider "google" {
  project = "{{.ProjectID}}"
  region  = "asia-northeast1"
  zone    = "asia-northeast1-a"
}
`
	return createFile("github_action", "provider.tf", providerTmpl, gcpConfig)
}
