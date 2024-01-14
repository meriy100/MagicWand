
resource "google_project_service" "iamcredentials" {
  service            = "iamcredentials.googleapis.com"
  disable_on_destroy = false
}

resource "google_service_account" "github_action" {
  account_id   = "github-action"
  display_name = "Github Action CI/CD"
}

resource "google_project_iam_binding" "project" {
  project = "sample-app}}"
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
