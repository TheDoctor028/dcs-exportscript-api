package utils

import (
	"github.com/thedoctor028/dcsexportscriptapi/dcs"
	"regexp"
	"strings"
)

func ExtractUIDAndValue(str string, splitChar string) *dcs.DCSExportData {

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

func testStrLen(str string) bool {
	return len(str) == 0
}

func convertKeyValStrToMap(strSplitedByChar []string) (*dcs.DCSExportData, bool) {
	keyValueMap := dcs.NewDcsData()

	// Remove values without UID(key)
	for i := range strSplitedByChar {
		iStr := strSplitedByChar[i]
		matched, err := regexp.MatchString("^\\d+=[-]?\\d+[.]?\\d*$", iStr)
		if err != nil {
			return nil, false
		}

		if matched {
			keyValSplited := strings.Split(iStr, "=")
			keyValueMap.Data[keyValSplited[0]] = keyValSplited[1]
		}
	}
	return &keyValueMap, true
}
