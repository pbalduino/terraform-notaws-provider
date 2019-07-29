package main

import (
  "github.com/hashicorp/terraform/helper/schema"
)

func Provider() *schema.Provider {
  return &schema.Provider{
    Schema: map[string]*schema.Schema{
      "region": {
        Type:     schema.TypeString,
        Required: true,
        DefaultFunc: schema.MultiEnvDefaultFunc([]string{
          "AWS_REGION",
          "AWS_DEFAULT_REGION",
        }, nil),
        Description:  "The region where AWS operations will take place. Examples\nare us-east-1, us-west-2, etc.",
        InputDefault: "us-east-1",
      },
    },
    ResourcesMap: map[string]*schema.Resource{
      "notaws_qldb_ledger": resourceQldbLedger(),
    },
  }
}
