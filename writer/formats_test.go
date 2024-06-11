package writer_test

import (
	"testing"

	"github.com/Method-Security/pkg/writer"
)

func TestIsUnknown(t *testing.T) {
	unknownFormat := writer.New_Format(writer.Format_Value("unknown"))

	if !unknownFormat.IsUnknown() {
		t.Errorf("Expected unknown format to be unknown, but it is not")
	}

	jsonFormat := writer.New_Format(writer.JSON)

	if jsonFormat.IsUnknown() {
		t.Errorf("Expected JSON format to not be unknown, but it is")
	}

	yamlFormat := writer.New_Format(writer.YAML)

	if yamlFormat.IsUnknown() {
		t.Errorf("Expected YAML format to not be unknown, but it is")
	}

	signalFormat := writer.New_Format(writer.SIGNAL)

	if signalFormat.IsUnknown() {
		t.Errorf("Expected SIGNAL format to not be unknown, but it is")
	}
}
func TestUnmarshalText(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected writer.Format
	}{
		{
			name:     "Valid JSON",
			input:    []byte("json"),
			expected: writer.New_Format(writer.JSON),
		},
		{
			name:     "Valid YAML",
			input:    []byte("yaml"),
			expected: writer.New_Format(writer.YAML),
		},
		{
			name:     "Valid SIGNAL",
			input:    []byte("signal"),
			expected: writer.New_Format(writer.SIGNAL),
		},
		{
			name:     "Unknown format",
			input:    []byte("unknown"),
			expected: writer.New_Format(writer.UNKNOWN),
		},
		{
			name:     "Empty input",
			input:    []byte(""),
			expected: writer.New_Format(writer.Format_Value("")),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var format writer.Format
			err := format.UnmarshalText(test.input)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if format != test.expected {
				t.Errorf("Expected format %v, but got %v", test.expected, format)
			}
			result, err := format.MarshalText()
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if string(result) != string(test.input) {
				t.Errorf("Expected marshaled text %s, but got %s", test.input, result)
			}
		})
	}
}
