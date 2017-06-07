package main

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/service/ec2"
)

var (
	nameString  = "Name"
	fooString   = "foo"
	barString   = "bar"
	emptyString = ""
)

func TestGetName_ShouldReturnAnEmptyString_WhenArgumentIsNil(t *testing.T) {
	if ret := getName(nil); ret != emptyString {
		t.Error(fmt.Sprintf("Expected '%s', got '%s'", emptyString, ret))
	}
}

func TestGetName_ShouldReturnNameTagValue_WhenNameTagPresent(t *testing.T) {
	tags := []*ec2.Tag{&ec2.Tag{Key: &nameString, Value: &fooString}}

	if ret := getName(tags); ret != fooString {
		t.Error(fmt.Sprintf("Expected '%s', got '%s'", fooString, ret))
	}
}

func TestGetName_ShouldReturnEmptyString_WhenNameTagNotPresent(t *testing.T) {
	tags := []*ec2.Tag{}

	if ret := getName(tags); ret != emptyString {
		t.Error(fmt.Sprintf("Expected '%s', got '%s'", emptyString, ret))
	}
}
