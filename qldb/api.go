package qldb

import (
	// "encoding/json"
  "log"

	"github.com/aws/aws-sdk-go/aws/request"
)

type DescribeLedgerInput struct {
	_ struct{} `type:"structure"`

	Name *string `locationName:"Name"`
}

type GetLedgerOutput struct {
	_ struct{} `type:"structure"`

  Arn *string
  CreationDateTime *uint64
  Name *string
  State *string
}

const opsDescribeLedger = "DescribeLedger"

func (c *Qldb) DescribeLedgerRequest(input *DescribeLedgerInput) (req *request.Request, output *GetLedgerOutput) {
	op := &request.Operation{
		Name:       opsDescribeLedger,
		HTTPMethod: "GET",
		HTTPPath:   "/ledgers/core-banking-event-store-dev",
	}

	if input == nil {
		input = &DescribeLedgerInput{}
	}

	log.Printf("input: %s\n", input)

	output = &GetLedgerOutput{}
	req = c.newRequest(op, input, output)

	return
}

func (c *Qldb) DescribeLedger(input *DescribeLedgerInput) (*GetLedgerOutput, error) {
	req, out := c.DescribeLedgerRequest(input)
	log.Printf("DescribeLedger req: %s\n", req)
	log.Printf("DescribeLedger out: %s\n", out)
	s := req.Send()
	log.Printf("DescribeLedger s: %s\n", s)
	return out, s
}
