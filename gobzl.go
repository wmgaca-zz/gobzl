package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/olekukonko/tablewriter"
)

func getName(tags []*ec2.Tag) string {
	for _, tag := range tags {
		if *tag.Key == "Name" {
			return *tag.Value
		}
	}

	return ""
}

func listInstances() {
	conn := ec2.New(session.New(), &aws.Config{Region: aws.String("eu-west-1")})

	resp, err := conn.DescribeInstances(nil)
	if err != nil {
		panic(err)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetBorder(false)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetRowSeparator("")
	table.SetColumnSeparator("")
	table.SetHeader([]string{
		"ID", "NAME", "PRIVATE IP", "PUBLIC IP", "TYPE",
	})

	for _, reservation := range resp.Reservations {
		for _, instance := range reservation.Instances {
			pub := ""
			if instance.PublicIpAddress != nil {
				pub = *instance.PublicIpAddress
			}

			table.Append([]string{
				*instance.InstanceId,
				getName(instance.Tags),
				*instance.PrivateIpAddress,
				pub, //*instance.PublicIpAddress,
				*instance.InstanceType,
			})
		}
	}

	table.Render()
}

func printUsage() {
	fmt.Println("usage: gobzl instances")
}

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "instances":
		listInstances()
	default:
		printUsage()
		os.Exit(1)
	}
}
