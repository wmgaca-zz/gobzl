package main

import (
	"context"
	"flag"
	"os"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/rds"
	"github.com/google/subcommands"
	"github.com/olekukonko/tablewriter"
)

type instancesCommand struct {
	region string
}

func (*instancesCommand) Name() string     { return "instances" }
func (*instancesCommand) Synopsis() string { return "List EC2 instances." }
func (*instancesCommand) Usage() string    { return "instances [-region <aws-region-name>]" }

func (c *instancesCommand) SetFlags(f *flag.FlagSet) {
	f.StringVar(&c.region, "region", "eu-west-1", "AWS region name")
}

func (c *instancesCommand) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	listInstances(c.region)

	return subcommands.ExitSuccess
}

type rdsInstancesCommand struct {
	region string
}

func (*rdsInstancesCommand) Name() string     { return "rds-instances" }
func (*rdsInstancesCommand) Synopsis() string { return "List RDS DB instances." }
func (*rdsInstancesCommand) Usage() string    { return "rds-instances [-region <aws-region-name>]" }

func (c *rdsInstancesCommand) SetFlags(f *flag.FlagSet) {
	f.StringVar(&c.region, "region", "eu-west-1", "AWS region name")
}

func (c *rdsInstancesCommand) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	listRdsInstances(c.region)

	return subcommands.ExitSuccess
}

func getName(tags []*ec2.Tag) string {
	for _, tag := range tags {
		if *tag.Key == "Name" {
			return *tag.Value
		}
	}

	return ""
}

func NewTable(headers []string) *tablewriter.Table {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetBorder(false)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetRowSeparator("")
	table.SetColumnSeparator("")
	table.SetHeader(headers)

	return table
}

func listInstances(region string) {
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

func listRdsInstances(region string) {
	conn := rds.New(session.New(), &aws.Config{Region: aws.String(region)})

	resp, err := conn.DescribeDBInstances(nil)
	if err != nil {
		panic(err)
	}

	table := NewTable([]string{
		"ID", "NAME", "ENGINE", "CLASS", "ENDPOINT", "PORT",
	})

	for _, dbInstance := range resp.DBInstances {
		table.Append([]string{
			stringFromPointer(dbInstance.DBInstanceIdentifier, ""),
			stringFromPointer(dbInstance.DBName, ""),
			stringFromPointer(dbInstance.Engine, ""),
			stringFromPointer(dbInstance.DBInstanceClass, ""),
			stringFromPointer(dbInstance.Endpoint.Address, ""),
			strconv.FormatInt(*dbInstance.Endpoint.Port, 10),
		})
	}

	table.Render()
}

func stringFromPointer(value *string, defaultValue string) string {
	if value == nil {
		return defaultValue
	}

	return *value
}

func main() {
	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")
	subcommands.Register(&instancesCommand{}, "")
	subcommands.Register(&rdsInstancesCommand{}, "")

	flag.Parse()
	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))
}
