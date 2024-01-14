
terraform {
  backend "gcs" {
    bucket = "sample-app-terraform-state"
    prefix = "terraform/github_actions"
  }
}
