package fuzzprocessor

import "github.com/thoas/picfit/tests"

func Fuzz(data []byte) int {
	op := string(data)
	processor := tests.NewDummyProcessor()
	_, err := processor.NewEngineOperationFromQuery(op)
	if err != nil { return 1 }
	return 0
}
