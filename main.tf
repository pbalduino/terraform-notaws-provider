provider "notaws" {
  region = "us-east-1"
}

resource "notaws_qldb_ledger" "my-awesome-ledger" {
  name = "my-awesome-ledger",
  permissions_mode = "ALLOW_ALL"
}
