package sprint

func ShiftBy(r rune, step int) rune {
	runeStep := ((step % 26) + 26) % 26
	newPosition := int(char-'a'+rune(runeStep)) % 26
	return 'a' + rune(newPosition)
}