package hw02unpackstring

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "a4bc2d5e", expected: "aaaabccddddde"},
		{input: "abccd", expected: "abccd"},
		{input: "", expected: ""},
		{input: "aaa0b", expected: "aab"},
		// uncomment if task with asterisk completed
		// {input: `qwe\4\5`, expected: `qwe45`},
		// {input: `qwe\45`, expected: `qwe44444`},
		// {input: `qwe\\5`, expected: `qwe\\\\\`},
		// {input: `qwe\\\3`, expected: `qwe\3`},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			result, err := Unpack(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestUnpackNew(t *testing.T) {
	data := map[string]string{
		"a4bc2d5e": "aaaabccddddde",
		"abcd":     "abcd",
	}

	for s, e := range data {
		r, err := Unpack(s)
		if err != nil {
			t.Fatalf("bad unpack for %s: got error %v", s, err)
		}
		if r != e {
			t.Fatalf("bad unpack for %s: got %v expected %v", s, r, e)
		}
	}
}

func TestUnpackError(t *testing.T) {
	s := "45"
	r, err := Unpack(s)
	if r != "" {
		t.Fatalf("bad unpack for %s: expected empty string", s)
	}
	if err == nil {
		t.Fatalf("bad unpack for %s: expected error", s)
	}
}

func TestUnpackEscape(t *testing.T) {
	data := map[string]string{
		"qwe\\4\\5": "qwe45",
		"qwe\\45":   "qwe44444",
		"qwe\\\\5":  "qwe\\\\\\\\\\",
	}

	for s, e := range data {
		r, err := Unpack(s)
		if err != nil {
			t.Fatalf("bad unpack for %s: got error %v", s, err)
		}
		if r != e {
			t.Fatalf("bad unpack for %s: got %v expected %v", s, r, e)
		}
	}
}
