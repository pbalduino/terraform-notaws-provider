package main

import (
  "log"

  "github.com/hashicorp/terraform/helper/schema"

  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/awserr"
  "github.com/aws/aws-sdk-go/aws/endpoints"

	// "github.com/aws/aws-sdk-go/aws/awsutil"
	// "github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/aws/request"
	// "github.com/aws/aws-sdk-go/private/protocol"
	// "github.com/aws/aws-sdk-go/private/protocol/restjson"
  "stash.int.klarna.net/plinio.balduino/terraform-qldb/qldb"
)

func QldbResolver(service, region string, optFns ...func(*endpoints.Options)) (endpoints.ResolvedEndpoint, error) {
    if service == "qldb" {
      return endpoints.ResolvedEndpoint{
        URL:           "http://qldb.us-east-1.amazonaws.com/",
        SigningRegion: "us-east-1",
      }, nil
    }

    return endpoints.DefaultResolver().EndpointFor(service, region, optFns...)
}

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
  sess := session.Must(session.NewSession(&aws.Config{
	   Region: aws.String("us-east-1"),
     EndpointResolver: endpoints.ResolverFunc(QldbResolver),
     // DisableSSL: &[]bool{true}[0],
   }))

  qldb.New(sess)

  return qldbLedgerRead(d, m)
}

func qldbLedgerRead(d *schema.ResourceData, m interface{}) error {
  params := &qldb.DescribeLedgerInput{
    // Name: aws.String(d.Get("name").(string)),
  }

  conn := qldb.New(session.New())

  getLedgerOutput, err := conn.DescribeLedger(params)

  if err != nil {
    if awsErr, ok := err.(awserr.Error); ok && awsErr.Code() == "ResourceNotFoundException" && !d.IsNewResource() {
      d.SetId("")
      return nil
    }
    return err
  }

  log.Printf("bleh %s\n", getLedgerOutput)

  return nil
}

func qldbLedgerUpdate(d *schema.ResourceData, m interface{}) error {
  log.Println("Ledger update")
  return qldbLedgerRead(d, m)
}

func qldbLedgerDelete(d *schema.ResourceData, m interface{}) error {
  return nil
}

func test() {
  log.Println("Testing")
  params := &qldb.DescribeLedgerInput{
    // Name: aws.String("core-banking-event-store-dev"),
  }

  sess := session.Must(session.NewSession(&aws.Config{
	   Region: aws.String("us-east-1"),
     EndpointResolver: endpoints.ResolverFunc(QldbResolver),
     // DisableSSL: &[]bool{true}[0],
   }))

  sess.Handlers.Send.PushFront(func(r *request.Request) {
    log.Printf("Request: %s | %s, Payload: %s",
		r.ClientInfo.ServiceName, r.Operation, r.Params)
  })

  conn := qldb.New(sess)

  getLedgerOutput, err := conn.DescribeLedger(params)

  if err != nil {
    awsErr, ok := err.(awserr.Error)
    if ok {
      log.Printf("Error: %s - %s\n", awsErr.Code(), awsErr.Message())
      return
    }
    log.Println("?")
    return
  }

  log.Printf("bleh %s\n", getLedgerOutput)

  return
}
