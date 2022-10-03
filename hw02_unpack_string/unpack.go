package hw02unpackstring

import (
	"log"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

func Unpack(PackedString string) string {
	var lastRune, lastLetter rune
	var result, num strings.Builder
	var esc bool
	result.Reset()
	num.Reset()
	lastRune = 0
	lastLetter = 0
	for i, curRune := range PackedString {
		// early return
		if unicode.IsDigit(curRune) && i == 0 {
			return ""
		}
		if unicode.IsLetter(curRune) {
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
			result.WriteRune(curRune)
			lastLetter = curRune
			lastRune = curRune
		}
		if unicode.IsDigit(curRune) {
			if esc {
				result.WriteRune(curRune)
				lastLetter = curRune
				lastRune = curRune
				esc = false
			} else {
				if unicode.IsLetter(lastRune) {
					num.Reset()
				}
				num.WriteRune(curRune)
				lastRune = curRune
				if i == utf8.RuneCountInString(string(PackedString))-1 {
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

	return ""
}
