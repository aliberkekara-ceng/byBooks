package services

import (
	"backend/models"
	"testing"
)

func TestProcessURL(t *testing.T) {
	service := NewURLService()

	tests := []struct {
		name      string
		request   models.URLRequest
		expected  string
		expectErr bool
	}{
		{
			name: "All operation - example 1",
			request: models.URLRequest{
				URL:       "https://BYFOOD.com/food-EXPeriences?query=abc/",
				Operation: "all",
			},
			expected:  "https://www.byfood.com/food-experiences",
			expectErr: false,
		},
		{
			name: "Canonical operation - example 2",
			request: models.URLRequest{
				URL:       "https://BYFOOD.com/food-EXPeriences?query=abc/",
				Operation: "canonical",
			},
			expected:  "https://BYFOOD.com/food-EXPeriences",
			expectErr: false,
		},
		{
			name: "Redirection operation",
			request: models.URLRequest{
				URL:       "https://example.com/some-path/",
				Operation: "redirection",
			},
			expected:  "https://www.byfood.com/some-path/",
			expectErr: false,
		},
		{
			name: "Invalid URL format",
			request: models.URLRequest{
				URL:       "invalid-url%%%",
				Operation: "canonical",
			},
			expected:  "",
			expectErr: true,
		},
		{
			name: "Invalid operation type",
			request: models.URLRequest{
				URL:       "https://byfood.com",
				Operation: "invalid_op",
			},
			expected:  "",
			expectErr: true,
		},
		{
			name: "Missing scheme - prepends https",
			request: models.URLRequest{
				URL:       "BYFOOD.com/food-EXPeriences",
				Operation: "redirection",
			},
			expected:  "https://www.byfood.com/food-experiences",
			expectErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := service.ProcessURL(tt.request)
			if (err != nil) != tt.expectErr {
				t.Errorf("expected error: %v, got error: %v", tt.expectErr, err)
			}
			if result != tt.expected {
				t.Errorf("expected: %s, got: %s", tt.expected, result)
			}
		})
	}
}
