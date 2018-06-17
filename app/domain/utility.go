package domain

import (
	"unicode/utf8"
)

func calculateScore(chip *UserChip) int {
	return utf8.RuneCountInString(chip.Description)
}
