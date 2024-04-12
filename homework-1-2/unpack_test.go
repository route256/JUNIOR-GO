package unpack

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnpackRequired(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "", expected: ""},
		{input: "a", expected: "a"},
		{input: "a4bc2d5e", expected: "aaaabccddddde"},
		{input: "a4bc2d5e2", expected: "aaaabccdddddee"},
		{input: "abccd", expected: "abccd"},
		{input: "aaa0b", expected: "aab"},
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

func TestUnpackAsterisk(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: `qwe\4\5`, expected: `qwe45`},
		{input: `qwe\40`, expected: `qwe`},
		{input: `qwe\45`, expected: `qwe44444`},
		{input: `qwe\\5`, expected: `qwe\\\\\`},
		{input: `qwe\\\3`, expected: `qwe\3`},
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

func TestUnpackInvalidString(t *testing.T) {
	invalidStrings := []string{"3abc", "45", "aaa10b", `qw\ne`}
	for _, tc := range invalidStrings {
		tc := tc
		t.Run(tc, func(t *testing.T) {
			_, err := Unpack(tc)
			require.Truef(t, errors.Is(err, ErrInvalidString), "actual error %q", err)
		})
	}
}
