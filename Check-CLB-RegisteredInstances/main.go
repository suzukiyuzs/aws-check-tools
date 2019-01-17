package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elb"
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

	svc := elb.New(sess, &aws.Config{Region: aws.String(region)})
	resp, err := svc.DescribeLoadBalancers(nil)
	if err != nil {
		panic(err)
	}

	for _, l := range resp.LoadBalancerDescriptions {
		regNum := fmt.Sprint(len(l.Instances))

		if regNum != "0" {
			ok++
		} else {
			ng++
		}
	}
	fmt.Printf("%d,%d\n", ok, ng)
}
