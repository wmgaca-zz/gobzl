package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func getName(tags []*ec2.Tag) string {
	for _, tag := range tags {
		if *tag.Key == "Name" {
			return *tag.Value
		}
	}

	return ""
}

func ListEc2Instances(region string) {
	conn := ec2.New(session.New(), &aws.Config{Region: aws.String(region)})

	resp, err := conn.DescribeInstances(nil)
	if err != nil {
		panic(err)
	}

	table := NewTable([]string{
		"ID", "NAME", "PRIVATE IP", "PUBLIC IP", "TYPE", "STATE",
	})

	for _, reservation := range resp.Reservations {
		for _, instance := range reservation.Instances {
			table.Append([]string{
				stringFromPointer(instance.InstanceId, ""),
				getName(instance.Tags),
				stringFromPointer(instance.PrivateIpAddress, ""),
				stringFromPointer(instance.PublicIpAddress, ""),
				stringFromPointer(instance.InstanceType, ""),
				stringFromPointer(instance.State.Name, ""),
			})
		}
	}

	table.Render()
}
