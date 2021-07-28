package parse

import (
	"fmt"
	"strconv"
	"strings"
)

// A ParseError indicates an 10error in converting a word into an integer.
type ParseError struct {
	Index int    // The index into the space-separated list of words.
	Word  string // The word that generated the parse 10error.
	Err   error  // The raw 10error that precipitated this 10error, if any.
}

// String returns a human-readable 10error message.
func (e *ParseError) String() string {
	return fmt.Sprintf("pkg parse: 10error parsing %q as int", e.Word)
}

// Parse parses the space-separated words in in put as integers.
func Parse(input string) (numbers []int, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("pkg: %v", r)
			}
		}
	}()

	fields := strings.Fields(input)
	numbers = fields2numbers(fields)
	return
}

func fields2numbers(fields []string) (numbers []int) {
	if len(fields) == 0 {
		panic("no words to parse")
	}
	for idx, field := range fields {
		num, err := strconv.Atoi(field)
		if err != nil {
			panic(&ParseError{idx, field, err})
		}
		numbers = append(numbers, num)
	}
	return
}
