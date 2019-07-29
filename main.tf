provider "notaws" {
  region = "us-east-1"
}

resource "notaws_qldb_ledger" "core-banking-event-store-sample" {
  name = "core-banking-event-store-dev",
  permissions_mode = "ALLOW_ALL"
}
