package dcsexportscriptapi

import (
	"bytes"
	"encoding/ascii85"
	"encoding/gob"
	"expvar"
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

func extractUIAndValue(str string) []expvar.KeyValue {
	if len(str) == 0 {
		return nil
	}

	//var kv []expvar.KeyValue

	// Remove memory address from the start of the string
	strWOAddr := strings.Split(str, "*")[1]

	if len(strWOAddr) == 0 {
		return nil
	}

	strSingleLine := strings.ReplaceAll(strWOAddr, "\n", "")

	if len(strSingleLine) == 0 {
		return nil
	}

	return nil
}
