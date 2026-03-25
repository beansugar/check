package do

import (
	"testing"
)

func TestCheck(t *testing.T) {
	test := []struct {
		name  string
		input string
		max   int
		want  bool
		min   int
	}{
		{name: "test1", input: "12345qqe", want: true, min: 3, max: 7},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			result := NewCheck(tt.input).Min(tt.min).Max(tt.max)
			if result.Input != tt.input {
				t.Logf("have some bug %v", result)
			} else {
				t.Log("test successful")
			}
		})
	}
}
