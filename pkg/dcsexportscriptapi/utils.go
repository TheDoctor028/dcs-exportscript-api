package dcsexportscriptapi

import (
	"bytes"
	"encoding/ascii85"
	"encoding/gob"
	"regexp"
	"strings"
)

func DeserializeDataStream(dataStream *[]byte) (string, error) {
	buffer := bytes.NewBuffer(*dataStream)
	encoder := gob.NewEncoder(buffer)
	err := encoder.Encode(buffer)
	if err != nil {
		return "", err
	}

	var encodedDataStream []byte
	ascii85.Encode(encodedDataStream, *dataStream)

	var b strings.Builder

	_, err = b.Write(encodedDataStream)

	if err != nil {
		return "nil", err
	}

	return b.String(), nil
}

func testStrLen(str string) bool {
	return len(str) == 0
}

func ExtractUIDAndValue(str string, splitChar string) map[string]string {
	if testStrLen(str) {
		return nil
	}

	// Remove memory address from the start of the string
	strWOAddr := strings.Split(str, "*")[1]

	if testStrLen(strWOAddr) {
		return nil
	}

	// Remove EOL(s) and spaces from the string
	strSingleLine := strings.ReplaceAll(strWOAddr, "\n", "")
	strSingleLine = strings.ReplaceAll(strSingleLine, " ", "")

	if testStrLen(strSingleLine) {
		return nil
	}

	// Split the string to sub strings by split char
	strSplitedByChar := strings.Split(strSingleLine, splitChar)

	if len(strSplitedByChar) == 0 {
		return nil
	}

	keyValStrMap, done := convertKeyValStrToMap(strSplitedByChar)
	if !done {
		return nil
	}

	return keyValStrMap
}

func convertKeyValStrToMap(strSplitedByChar []string) (map[string]string, bool) {
	keyValueMap := make(map[string]string)

	// Remove values without UID(key)
	for i := range strSplitedByChar {
		iStr := strSplitedByChar[i]
		matched, err := regexp.MatchString("^\\d+=[-]?\\d+[.]?\\d*$", iStr)
		if err != nil {
			return nil, false
		}

		if matched {
			keyValSplited := strings.Split(iStr, "=")
			keyValueMap[keyValSplited[0]] = keyValSplited[1]
		}
	}
	return keyValueMap, true
}
