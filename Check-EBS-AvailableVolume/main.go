package main

/*
EBS Volume State:
  Creating
  Available
  In Use
  Deleting
  Deleted
  Error
*/

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

var (
	region = "ap-northeast-1"
	ok     = 0
	ng     = 0
	okStr  = "in-use"
	ngStr  = "available"
)

func main() {
	sess, err := session.NewSession()
	if err != nil {
		panic(err)
	}

	svc := ec2.New(sess, &aws.Config{Region: aws.String(region)})
	resp, err := svc.DescribeVolumes(nil)
	if err != nil {
		panic(err)
	}

	for _, v := range resp.Volumes {
		state := fmt.Sprint(*v.State)
		if state == okStr {
			ok++
		} else if state == ngStr {
			ng++
		}
	}

	fmt.Printf("%d,%d\n", ok, ng)
}
