package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
)

var (
	region    = "ap-northeast-1"
	namespace = "Namespace"
	okState   = "OK"
	ngState   = "ALARM"
	ok        = 0
	ng        = 0
	metric    string
)

func main() {
	sess, err := session.NewSession()
	if err != nil {
		panic(err)
	}

	svc := cloudwatch.New(sess, &aws.Config{Region: aws.String(region)})
	resp, err := svc.DescribeAlarms(nil)
	if err != nil {
		panic(err)
	}

	metric = os.Args[1]
	for _, m := range resp.MetricAlarms {
		if *m.Namespace != "AWS/EC2" {
			continue
		}
		if *m.MetricName != metric {
			continue
		}

		if *m.StateValue == okState {
			ok++
		} else if *m.StateValue == ngState {
			ng++
		}
	}
	fmt.Printf("%d,%d\n", ok, ng)
}
