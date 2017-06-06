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
			table.Append([]string{
				stringFromPointer(instance.InstanceId, ""),
				getName(instance.Tags),
				stringFromPointer(instance.PrivateIpAddress, ""),
				stringFromPointer(instance.PublicIpAddress, ""),
				stringFromPointer(instance.InstanceType, ""),
			})
		}
	}

	table.Render()
}

func stringFromPointer(value *string, defaultValue string) string {
	if value == nil {
		return defaultValue
	}

	return *value
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
