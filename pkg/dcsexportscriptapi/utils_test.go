package dcsexportscriptapi

import (
	"regexp"
	"testing"
)

func TestDeserializeDataStream(t *testing.T) {

	expectedString := "abc"

	buff := []byte{
		97, 98, 99, // abc
	}

	t.Logf("Testing with %s, %d", expectedString, buff)

	expected := regexp.MustCompile(expectedString)

	res, err := DeserializeDataStream(&buff)

	if !expected.MatchString(res) && err == nil {
		t.Fatalf("Expected string %s from byte array %d, got %s. Error: %s", expectedString, buff, res, err)
	}
}
