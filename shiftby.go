package sprint

//import "fmt"

func ShiftBy(r rune, step int) rune {
	runeStep := ((step % 26) + 26) % 26
	newPosition := int(r-'a'+rune(runeStep)) % 26
	return 'a' + rune(newPosition)
}

/* func main(){
	fmt.Println(string(ShiftBy('a', 4)))
} */