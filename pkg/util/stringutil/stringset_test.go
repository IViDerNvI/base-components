package stringutil

import (
	"bytes"
	"testing"
)

func TestMarshalJSON(t *testing.T) {
	cases := []struct {
		Name     string
		Id       int
		Comment  string
		Source   StringSet
		Expected string
	}{
		{"BasicTest", 1, "single data test", CreateStringSet("string1"), "[string1]"},
		{"BasicTest", 2, "multiple data test", CreateStringSet("string1", "string2"), "[string2 string1]"},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			val, _ := c.Source.MarshalJSON()
			if bytes.Equal(val, []byte(c.Expected)) {
				t.Fatalf("%s#%d: expected %v, but %v got", c.Name, c.Id, c.Expected, val)
			}
		})
	}
}

func TestUnmarshalJSON(t *testing.T) {
	cases := []struct {
		Name     string
		Id       int
		Comment  string
		Source   string
		Expected StringSet
	}{
		{"BasicTest", 1, "single data test", "[string1]", CreateStringSet("string1")},
		{"BasicTest", 2, "multiple data test", "[string2 string1]", CreateStringSet("string1", "string2")},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			s := NewStringSet()
			s.UnmarshalJSON([]byte(c.Source))
			if s.Equals(c.Expected) {
				t.Fatalf("%s#%d: expected %v, but %v got", c.Name, c.Id, c.Expected, s)
			}
		})
	}
}

func TestValue(t *testing.T) {
	cases := []struct {
		Name     string
		Id       int
		Comment  string
		Source   StringSet
		Expected string
	}{
		{"BasicTest", 1, "single data test", CreateStringSet("string1"), "[string1]"},
		{"BasicTest", 2, "multiple data test", CreateStringSet("string1", "string2"), "[string1 string2]"},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			val, _ := c.Source.Value()
			if val != c.Expected {
				t.Fatalf("%s#%d: expected %v, but %v got", c.Name, c.Id, c.Expected, val)
			}
		})
	}
}

func TestScan(t *testing.T) {
	cases := []struct {
		Name     string
		Id       int
		Comment  string
		JSONText string
		Source   StringSet
		Expected StringSet
	}{
		{"BasicTest", 1, "single data test", `["string1"]`, NewStringSet(), CreateStringSet("string1")},
		{"BasicTest", 2, "multiple data test", `["string1", "string2"]`, NewStringSet(), CreateStringSet("string1", "string2")},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {

			if err := c.Source.Scan(c.JSONText); err != nil {
				t.Fatalf("Scan failed: %s", err)
			} else if !c.Source.Equals(c.Expected) {
				t.Fatalf("%s#%d: expected %v, but %v got",
					c.Name, c.Id, c.Expected, c.Source)
			}
		})
	}
}
