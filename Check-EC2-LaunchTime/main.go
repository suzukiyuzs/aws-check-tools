package main

import (
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

const (
	layout = "2006-01-02T15:04:05"
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
	resp, err := svc.DescribeInstances(nil)
	if err != nil {
		panic(err)
	}

	compareTime, _ := time.Parse(layout, os.Args[1])
	for _, r := range resp.Reservations {
		for _, i := range r.Instances {
			launchTime := *i.LaunchTime

			if launchTime.After(compareTime) {
				ok++
			} else {
				ng++
			}
		}
	}

	fmt.Printf("%d,%d\n", ok, ng)
}
