package dcsexportscriptapi

import (
	"reflect"
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

func TestExtractUIDAndValue(t *testing.T) {
	text := "62b74d61*146=0.0017:141=0.0074 \n:128=-0.0682:146=0.0008:180=0.4000:141=0.0018\n=0:184=0:200=0"

	expected := getExtractUIDAndValueExpectedValues()

	t.Logf("Testing with %s", text)

	res := ExtractUIDAndValue(text, ":")

	if !reflect.DeepEqual(res, expected) {
		t.Fatalf("Expected: %s \n Got: %s", expected, res)
	}
}

func getExtractUIDAndValueExpectedValues() map[string]string {
	keyValueMap := make(map[string]string)
	keyValueMap["200"] = "0"
	keyValueMap["184"] = "0"
	keyValueMap["141"] = "0.0074"
	keyValueMap["180"] = "0.4000"
	keyValueMap["146"] = "0.0008"
	keyValueMap["128"] = "-0.0682"

	return keyValueMap
}
