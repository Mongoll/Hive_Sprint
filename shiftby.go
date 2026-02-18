package sprint

import "fmt"

func ShiftBy(r rune, step int) rune {
	runeStep := ((step % 26) + 26) % 26 // we need positiv argument, move to -1=step forward fo 25 (always from 0 to 25)
	newPosition := int(r-'a'+rune(runeStep)) % 26 // interpret letter to position -> add step -> move out of alphabet
	return 'a' + rune(newPosition) // interpret position to letter
}

func main(){
	fmt.Println(string(ShiftBy('a', 4)))
}