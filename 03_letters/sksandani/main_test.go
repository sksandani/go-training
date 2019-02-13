package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var tests = []struct {
	name string
	word string
	want map[rune]int
}{
	{"empty input string", "abcdad", map[rune]int{98: 1, 99: 1, 100: 2, 97: 2}}}

func TestLetters(t *testing.T) {
	r := require.New(t)
	for _, tt := range tests {
		got := letters(tt.word)
		r.Equal(tt.want, got)
	}
}
