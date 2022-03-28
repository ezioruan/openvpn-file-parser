package parser

import (
	"testing"
)

type TestCase struct {
	FilePath string
}

func TestParseFromFile(t *testing.T) {

	testCases := []TestCase{
		{"../test-data/test.ovpn"},
	}

	for _, testCase := range testCases {
		config, err := ParseFromFile(testCase.FilePath)
		if err != nil {
			t.Errorf("ParseFromFile error %v", err)
		}
		if config.CA == "" {
			t.Errorf("ParseFromFile get CA error %s", config.CA)
		}

	}

}
