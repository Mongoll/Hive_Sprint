package main

func StrSplitBy(s, sep string) []string {
	var strarr []string

	if len(sep) == 0 {
		if len(s) == 0 {
			return []string{}
		}
		return []string{}
	}

	tempstr := ""
	i := 0
	for i < len(s) {
		if i+len(sep) <= len(s) && s[i:i+len(sep)] == sep {
			strarr = append(strarr, tempstr)
			tempstr = ""
			i += len(sep)
			continue
		}
		tempstr += string(s[i])
		i++
	}
	strarr = append(strarr, tempstr)
	return strarr
}