package main

import (
	"fmt"
	"testing"
)

func TestInt64FromPointer_WhenNotNil(t *testing.T) {
	value := int64(5)
	if 5 != int64FromPointer(&value, 0) {
		t.Error("Expected 5")
	}
}

func TestInt64FromPointer_WhenNil(t *testing.T) {
	var value int64
	if 0 != int64FromPointer(&value, 0) {
		t.Error("Expected 0")
	}
}

func TestStringromPointer_WhenNotNil(t *testing.T) {
	value := "foo"
	defaultValue := "bar"
	expectedValue := "foo"

	if retValue := stringFromPointer(&value, defaultValue); retValue != expectedValue {
		t.Error(fmt.Sprintf("Expected %s, got %s", expectedValue, retValue))
	}
}

func TestStringFromPointer_WhenNil(t *testing.T) {
	var value *string
	defaultValue := "bar"
	expectedValue := "bar"

	if retValue := stringFromPointer(value, defaultValue); retValue != expectedValue {
		t.Error(fmt.Sprintf("Expected %s, got %s", expectedValue, retValue))
	}
}
