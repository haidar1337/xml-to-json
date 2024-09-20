package main

import (
	"strings"
	"testing"
)

func TestToJSON(t *testing.T) {
	tests := map[string]struct{
		Input map[string]any
		Expected string
		ErrorContains string
	}{
		"Basic Conversion": {
			Input: map[string]any{
				"name": "John",
				"age": 25,
				"is_smart": true,
				"balance": 287.15, 
			},
			Expected: `{
				"name": "John",
				"age": 25,
				"is_smart": true,
				"balance": 287.15
			}`,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := ToJSON(tc.Input)
			if err != nil && tc.ErrorContains == "" {
				t.Errorf("Test **%s** FAIL: unexpected error: %v", name, err)
				return
			} else if err != nil && !strings.Contains(err.Error(), tc.ErrorContains) {
				t.Errorf("Test **%s** FAIL\nExpected Error to Contain: %s\nActual Error: %s", name, tc.ErrorContains, err.Error())
			} else if err == nil && tc.ErrorContains != "" {
				t.Errorf("Test **%s** FAIL: expected error, got nothing", name)
			}

			if actual != tc.Expected {
				t.Errorf("Test **%s** FAIL\nExpected: %v\nActual: %v", name, tc.Expected, actual)
			}
		})
	}
}