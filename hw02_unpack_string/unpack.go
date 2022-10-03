package hw02unpackstring

import (
	"log"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

type Unpacker interface {
	Unpack() string
}

type PackedString string

func (s PackedString) Unpack() string {
	var lastRune, lastLetter rune
	var result, num strings.Builder
	var esc bool
	result.Reset()
	num.Reset()
	lastRune = 0
	lastLetter = 0
	for i, curRune := range s {
		// early return
		if unicode.IsDigit(curRune) && i == 0 {
			return ""
		}
		// if letter writing it to result and optionally unpacking previous sequence
		if unicode.IsLetter(curRune) {
			// letter after digit
			if unicode.IsDigit(lastRune) {
				numRunes, err := strconv.Atoi(num.String())
				if err != nil {
					log.Fatal(err)
				}
				for j := 0; j < numRunes-1; j++ {
					result.WriteRune(lastLetter)
				}
				num.Reset()
			}
			// any letter
			result.WriteRune(curRune)
			lastLetter = curRune
			lastRune = curRune
		}
		// write to digit sequence or flush letters to result
		if unicode.IsDigit(curRune) {
			// escape digit
			if esc {
				result.WriteRune(curRune)
				lastLetter = curRune
				lastRune = curRune
				esc = false
			} else {
				// first digit of new digit sequence
				if unicode.IsLetter(lastRune) {
					num.Reset()
				}
				num.WriteRune(curRune)
				lastRune = curRune
				// last digit in input string
				if i == utf8.RuneCountInString(string(s))-1 {
					numRunes, err := strconv.Atoi(num.String())
					if err != nil {
						log.Fatal(err)
					}
					for j := 0; j < numRunes-1; j++ {
						result.WriteRune(lastLetter)
					}
				}
			}

		}
		if curRune == '\\' {
			if lastRune == '\\' {
				result.WriteRune(curRune)
				lastLetter = curRune
				lastRune = curRune
				esc = false

			} else {
				esc = true
				lastRune = curRune
			}
		}
	}

	return result.String()
}
