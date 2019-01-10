package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

var (
	region = "ap-northeast-1"
	ok     = 0
	ng     = 0
	tagKey string
)

func main() {
	sess, err := session.NewSession()
	if err != nil {
		panic(err)
	}

	svc := ec2.New(sess, &aws.Config{Region: aws.String(region)})
	resp, err := svc.DescribeInstances(nil)
	if err != nil {
		panic(err)
	}

	tagKey = os.Args[1]
	for _, r := range resp.Reservations {
		for _, i := range r.Instances {
			var tag string
			for _, t := range i.Tags {
				if *t.Key == tagKey {
					tag = *t.Value
				}
			}

			if tag == "" {
				ng++
			} else {
				ok++
			}
		}
	}

	fmt.Printf("%d,%d\n", ok, ng)
}
