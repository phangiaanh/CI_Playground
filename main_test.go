package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test the greet handler
func TestGreet(t *testing.T) {
	// override greeting flag for testing
	*greeting = "HelloTest"

	tests := []struct {
		name     string
		path     string
		expected string
	}{
		{
			name:     "root path",
			path:     "/",
			expected: "HelloTest, Gopher!",
		},
		{
			name:     "with name",
			path:     "/Alice",
			expected: "HelloTest, Alice!",
		},
		// {
		// 	name:     "html escape",
		// 	path:     "/<script>",
		// 	expected: "HelloTest, &lt;script&gt;!",
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, tt.path, nil)
			rr := httptest.NewRecorder()

			greet(rr, req)

			body := rr.Body.String()
			assert.Contains(t, body, tt.expected)
			assert.Equal(t, http.StatusOK, rr.Code)
		})
	}
}

// Test the version handler
func TestVersion(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/version", nil)
	rr := httptest.NewRecorder()

	version(rr, req)

	body := rr.Body.String()

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Contains(t, body, "mod") // because debug.ReadBuildInfo should include module data
}
