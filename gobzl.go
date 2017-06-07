package main

import (
	"context"
	"flag"
	"os"

	"github.com/google/subcommands"
)

type instancesCommand struct {
	region string
}

func (*instancesCommand) Name() string {
	return "instances"
}

func (*instancesCommand) Synopsis() string {
	return "List EC2 instances."
}

func (*instancesCommand) Usage() string {
	return "instances [-region <aws-region-name>]"
}

func (c *instancesCommand) SetFlags(f *flag.FlagSet) {
	f.StringVar(&c.region, "region", "eu-west-1", "AWS region name")
}

func (c *instancesCommand) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	ListEc2Instances(c.region)

	return subcommands.ExitSuccess
}

type rdsInstancesCommand struct {
	region string
}

func (*rdsInstancesCommand) Name() string {
	return "rds-instances"
}

func (*rdsInstancesCommand) Synopsis() string {
	return "List RDS DB instances."
}

func (*rdsInstancesCommand) Usage() string {
	return "rds-instances [-region <aws-region-name>]"
}

func (c *rdsInstancesCommand) SetFlags(f *flag.FlagSet) {
	f.StringVar(&c.region, "region", "eu-west-1", "AWS region name")
}

func (c *rdsInstancesCommand) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	ListRdsInstances(c.region)

	return subcommands.ExitSuccess
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
