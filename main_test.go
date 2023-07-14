package main

import "testing"

func TestTableCalcDirSize(t *testing.T) {
	type Inputs struct {
		inputDir   string
		outputName string
		expected   int64
	}

	var tests = []Inputs{
		{"./test/1", "./test/output_test_1", 91050},
		{"./test/2", "./test/output_test_2", 6730635},
	}

	for _, test := range tests {
		output, err := calcDirSize(test.inputDir, test.outputName)

		if err != nil {
			t.Errorf("Error: %s", err)
		}

		if output != test.expected {
			t.Errorf("Expected %d, got %d", test.expected, output)
		}
	}
}

func BenchmarkCalcDirSize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		calcDirSize("./test/1", "output_test")
	}
}
