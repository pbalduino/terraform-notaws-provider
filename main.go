package main

import (
  "github.com/hashicorp/terraform/plugin"
  "github.com/hashicorp/terraform/terraform"
)

func main() {
  // test()

  plugin.Serve(&plugin.ServeOpts{
    ProviderFunc: func() terraform.ResourceProvider {
      return Provider()
    },
  })
}
