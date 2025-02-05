// Utilities to deal with loading test cases for various platforms
package utils

import (
	"encoding/json"
	"io"
	"log"
)

// we can have test cases in various formats
// Columnar, Row based, one per file, one column per file, and so on.
// We will deal with each here

// A single pack that contains all test cases including inputs and expected values
type OneFile struct {
	TestCases       []any
	ExpectedResults []any
}

// This the simplest case.  Here a single file contains all the inputs and expected values for all test cases.
// This is for simple examples where cases are small and simple.  Eg 2-sum etc
func UnmarshalOneFileForAllCases(reader io.Reader, output any) error {
	return nil
}

type CommandTestCase struct {
	Commands []string `json:"commands"`
	Args     []any    `json:"args"`
	Expected []any    `json:"expected"`
}

func LoadCases[T any](contents []byte) (out []*T) {
	err := json.Unmarshal(contents, &out)
	if err != nil {
		panic(err)
	}
	return
}

func LoadLargeCase[T any](contents map[string][]byte) (out T) {
	// unmarshal each bytelist and marshall the whole thing
	interim := make(map[string]any)
	for k, v := range contents {
		if len(v) != 0 {
			var data any
			if err := json.Unmarshal(v, &data); err != nil {
				log.Println("Invalid contents for key: ", k, string(v))
				panic(err)
			}
			interim[k] = data
		}
	}
	final, err := json.Marshal(interim)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(final, &out)
	if err != nil {
		panic(err)
	}
	return
}
