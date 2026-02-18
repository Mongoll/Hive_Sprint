package sprint

import "fmt"

func StrConcat(s1, s2, delim string) string {
	return s1+delim+s2
}

 func main(){
	fmt.Println(StrConcat("r", "ProgrammerHumour", "/"))
}