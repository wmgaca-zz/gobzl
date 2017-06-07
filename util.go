package main

import (
	"github.com/olekukonko/tablewriter"
	"os"
)

func stringFromPointer(value *string, defaultValue string) string {
	if value == nil {
		return defaultValue
	}

	return *value
}

func int64FromPointer(value *int64, defaultValue int64) int64 {
	if value == nil {
		return defaultValue
	}

	return *value
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
