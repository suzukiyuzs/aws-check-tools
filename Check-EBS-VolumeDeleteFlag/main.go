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

	var deleteFlag bool
	for _, v := range resp.Volumes {
		for _, a := range v.Attachments {
			deleteFlag = *a.DeleteOnTermination
			if deleteFlag {
				ok++
			} else {
				ng++
			}
		}
	}

	fmt.Printf("%d,%d\n", ok, ng)
}
