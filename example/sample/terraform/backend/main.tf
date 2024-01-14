
provider "google" {
  project = "sample"
  region  = "asia-northeast1"
  zone    = "asia-northeast1-a"
}

resource "google_storage_bucket" "backend" {
  name          = "sample-terraform-state"
  location      = "asia-northeast1"
  force_destroy = false
  storage_class = "NEARLINE"
}
