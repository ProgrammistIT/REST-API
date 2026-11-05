package validator

import "testing"

func TestCharValidate(t *testing.T) {
	tests := []struct {
		name     string
		md5      string
		expected bool
	}{
		{
			name:     "Valid MD5",
			md5:      "d41d8cd98f00b204e9800998ecf8427e",
			expected: true,
		},
		{
			name:     "Uppercase MD5",
			md5:      "D41D8CD98F00B204E9800998ECF8427E",
			expected: true,
		},
		{
			name:     "Invalid characters",
			md5:      "z41d8cd98f00b204e9800998ecf8427e",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CharValidate(tt.md5)

			if result != tt.expected {
				t.Errorf(
					"CharValidate(%q) = %v, expected %v",
					tt.md5,
					result,
					tt.expected,
				)
			}
		})
	}
}
