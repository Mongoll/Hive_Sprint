package sprint

//import "fmt"

func ReverseAlphabetValue(ch rune) rune {
	
	return 'a' + 'z' - ch //for example: if ch=a unicode(97) then 97+122-97(ch)=122 -> z
}

func main(){
	fmt.Println(string(ReverseAlphabetValue('a')))
}