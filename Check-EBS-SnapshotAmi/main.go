package main

import (
	"fmt"
	"regexp"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

var (
	region  = "ap-northeast-1"
	ownerId = "Your AWS Account ID"
	ok      = 0
	ng      = 0
	amiList []string
)

/*
ex.) AMI Description
     Created by CreateImage(i-5d8851f8) for ami-fa2948fa from vol-ee2dd011
*/
var re = regexp.MustCompile(`^Created by CreateImage\(i-\w+\) for (ami-\w+) from vol-\w+`)

func main() {
	sess, err := session.NewSession()
	if err != nil {
		panic(err)
	}

	svc := ec2.New(sess, &aws.Config{Region: aws.String(region)})

	amiParams := &ec2.DescribeImagesInput{
		Owners: []*string{
			aws.String(ownerId),
		},
	}

	amiResp, err := svc.DescribeImages(amiParams)
	if err != nil {
		panic(err)
	}

	for _, i := range amiResp.Images {
		amiList = append(amiList, *i.ImageId)
	}

	ssParams := &ec2.DescribeSnapshotsInput{
		OwnerIds: []*string{
			aws.String(ownerId),
		},
	}

	ssResp, err := svc.DescribeSnapshots(ssParams)
	if err != nil {
		panic(err)
	}

	for _, s := range ssResp.Snapshots {
		desc := fmt.Sprint(*s.Description)
		result := re.FindAllStringSubmatch(desc, -1)
		if result == nil {
			continue
		}

		ssAmiId := fmt.Sprint(result[0][1])
		if contains(amiList, ssAmiId) {
			ok++
		} else {
			ng++
		}
	}

	fmt.Printf("%d,%d\n", ok, ng)
}

func contains(list []string, ssa string) bool {
	for _, l := range list {
		if ssa == l {
			return true
		}
	}
	return false
}
