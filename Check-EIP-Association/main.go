package main

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

var (
	region        = "ap-northeast-1"
	associationId = "AssociationId"
	ok            = 0
	ng            = 0
)

func main() {
	sess, err := session.NewSession()
	if err != nil {
		panic(err)
	}

	svc := ec2.New(sess, &aws.Config{Region: aws.String(region)})
	resp, err := svc.DescribeAddresses(nil)
	if err != nil {
		panic(err)
	}

	for _, a := range resp.Addresses {
		var s string
		s = fmt.Sprint(*a)
		if strings.Contains(s, associationId) {
			ok++
		} else {
			ng++
		}
	}
	fmt.Printf("%d,%d\n", ok, ng)
}
