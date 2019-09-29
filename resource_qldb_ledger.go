package main

import (
  "log"

  "github.com/hashicorp/terraform/helper/schema"

  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/qldb"
)

func resourceQldbLedger() *schema.Resource {
  return &schema.Resource{
    Create: qldbLedgerCreate,
    Read:   qldbLedgerRead,
    Update: qldbLedgerUpdate,
    Delete: qldbLedgerDelete,

    Schema: map[string]*schema.Schema{
            "name": &schema.Schema{
              Type:     schema.TypeString,
              Required: true,
              ForceNew: true,
            },
            "permissions_mode": &schema.Schema{
              Type:     schema.TypeString,
              Optional: true,
              InputDefault: "ALLOW_ALL",
            },
    },
  }
}

func qldbLedgerCreate(d *schema.ResourceData, m interface{}) error {
  sess := session.Must(session.NewSession())

  params := &qldb.CreateLedgerInput{
    Name: aws.String(d.Get("name").(string)),
    PermissionsMode: aws.String(d.Get("permissions_mode").(string)),
  }

  svc := qldb.New(sess)

  svc.CreateLedger(params)

  return qldbLedgerRead(d, m)
}

func qldbLedgerRead(d *schema.ResourceData, m interface{}) error {
  params := &qldb.DescribeLedgerInput{
    Name: aws.String(d.Get("name").(string)),
  }

  conn := qldb.New(session.New())

  getLedgerOutput, err := conn.DescribeLedger(params)

  log.Println("[DEBUG] getLedgerOutput: ", getLedgerOutput)

  if err != nil {
    if awsErr, ok := err.(awserr.Error); ok && awsErr.Code() == "ResourceNotFoundException" && !d.IsNewResource() {
      d.SetId("")
      return nil
    }
    return err
  }

  d.SetId(*getLedgerOutput.Arn)

  return nil
}

func qldbLedgerUpdate(d *schema.ResourceData, m interface{}) error {
  log.Println("[DEBUG] Ledger update")

  return qldbLedgerRead(d, m)
}

func qldbLedgerDelete(d *schema.ResourceData, m interface{}) error {
  d.SetId("")
  return nil
}

func test() {
  log.Println("[DEBUG] Testing")
  params := &qldb.DescribeLedgerInput{
    Name: aws.String("core-banking-event-store-wazaap"),
  }

  sess := session.Must(session.NewSession())

  conn := qldb.New(sess)

  getLedgerOutput, err := conn.DescribeLedger(params)

  if err != nil {
    awsErr, ok := err.(awserr.Error)
    if ok {
      log.Printf("Error: %s - %s\n", awsErr.Code(), awsErr.Message())
      return
    }
    log.Println("[DEBUG] ?")
    return
  }

  log.Printf("bleh %s\n", *getLedgerOutput.Arn)

  return
}
