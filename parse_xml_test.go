package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestSomething(t *testing.T) {
	tests := map[string]struct {
		Input         string
		Expected      map[string]any
		ErrorContains string
	}{
		"Basic XML": {
			Input: `<name>haidar</name>
<age>18</age>`,
			Expected: map[string]any{
				"name": "haidar",
				"age":  18,
			},
		},
		"Bad XML Input": {
			Input:         `<nam>haidar</name>`,
			Expected:      nil,
			ErrorContains: "bad xml input",
		},
		"Nested XML": {
			Input: `<person>
	<name>haidar</name>
	<age>18</age>
	<favoriteColor>blue</favoriteColor>
</person>`,
			Expected: map[string]any{
				"person": map[string]any{
					"name":          "haidar",
					"age":           18,
					"favoriteColor": "blue",
				},
			},
		},
		"Bad Nested XML": {
			Input: `<person>
	<name>haidar</name>
	<age>18
	<favoriteColor>blue</favoriteColor>
</person>`,
			Expected:      nil,
			ErrorContains: "bad xml input",
		},
		"Complex XML": {
			Input: `<student>
    <name>haidar</name>
    <age>18</age>
    <courses>
        <code>CALC429</code>
        <courseName>Calculus III</courseName>
        <instructor>Dr. John Smith</instructor>
        <unisex>false</unisex>
        <references>
            <book>Book X</book>
            <slides>Slides Y</slides>
        </references>
    </courses>
    <courses>
        <code>STAT492</code>
        <courseName>Statistics</courseName>
        <instructor>Dr. John Doe</instructor>
        <unisex>true</unisex>
        <references>
            <book>Book Z</book>
            <slides>Slides W</slides>
        </references>
    </courses>
</student>
`,
			Expected: map[string]any{
				"student": map[string]any{
					"name": "haidar",
					"age":  18,
					"courses": []map[string]any{
						{
							"code":       "CALC429",
							"courseName": "Calculus III",
							"instructor": "Dr. John Smith",
							"unisex":     false,
							"references": map[string]any{
								"book":   "Book X",
								"slides": "Slides Y",
							},
						},
						{
							"code":       "STAT492",
							"courseName": "Statistics",
							"instructor": "Dr. John Doe",
							"unisex":     true,
							"references": map[string]any{
								"book":   "Book Z",
								"slides": "Slides W",
							},
						},
					},
				},
			},
		},
		"XML with Attributes": {
			Input: `<student>
    <name>haidar</name>
    <age>18</age>
    <courses>
        <code>CALC429</code>
        <courseName>Calculus III</courseName>
        <instructor id="1">Dr. John Smith</instructor>
        <unisex>false</unisex>
        <references refs="https://www.google.com">
            <book>Book X</book>
            <slides>Slides Y</slides>
        </references>
    </courses>
    <courses>
        <code>STAT492</code>
        <courseName>Statistics</courseName>
        <instructor id="2">Dr. John Doe</instructor>
        <unisex>true</unisex>
        <references refs="https://www.youtube.com">
            <book>Book Z</book>
            <slides>Slides W</slides>
        </references>
    </courses>
</student>
`,
			Expected: map[string]any{
				"student": map[string]any{
					"name": "haidar",
					"age":  18,
					"courses": []map[string]any{
						{
							"code":       "CALC429",
							"courseName": "Calculus III",
							"instructor": map[string]any{
								"_id":    1,
								"__text": "Dr. John Smith",
							},
							"unisex": false,
							"references": map[string]any{
								"book":   "Book X",
								"slides": "Slides Y",
								"_refs":  "https://www.google.com",
							},
						},
						{
							"code":       "STAT492",
							"courseName": "Statistics",
							"instructor": map[string]any{
								"_id":    2,
								"__text": "Dr. John Doe",
							},
							"unisex": true,
							"references": map[string]any{
								"book":   "Book Z",
								"slides": "Slides W",
								"_refs":  "https://www.youtube.com",
							},
						},
					},
				},
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := ParseXML(tc.Input)
			if err != nil && tc.ErrorContains == "" {
				t.Errorf("Test **%s** FAIL: unexpected error: %v", name, err)
				return
			} else if err != nil && !strings.Contains(err.Error(), tc.ErrorContains) {
				t.Errorf("Test **%s** FAIL\nExpected Error to Contain: %s\nActual Error: %s", name, tc.ErrorContains, err.Error())
			} else if err == nil && tc.ErrorContains != "" {
				t.Errorf("Test **%s** FAIL: expected error, got nothing", name)
			}

			if !reflect.DeepEqual(actual, tc.Expected) {
				t.Errorf("Test **%s** FAIL\nExpected: %v\nActual: %v", name, tc.Expected, actual)
			}
		})
	}

}
