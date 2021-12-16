package libs

func ParseRuneNumber(char rune) int {
	if char >= 48 && char <= 57 {
		return int(char - 48)
	}

	panic("incorrect rune passed")
}
