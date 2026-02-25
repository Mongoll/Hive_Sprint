package main

func SplitWhitespaces(s string) []string {
	var strarr []string
	tempstr := ""
	for _, v := range s {
		if v == ' ' || v == '\t' || v == '\n' {
			if tempstr != "" {
				strarr = append(strarr, tempstr)
				tempstr = ""
			}
		} else {
			tempstr += string(v)
		}
	}
	if tempstr != "" {
		strarr = append(strarr, tempstr)
	}
	return strarr
}
