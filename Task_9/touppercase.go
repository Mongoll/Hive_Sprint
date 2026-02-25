package main

func ToUpperCase(s string) string {
	result := []rune(s)
	for i, v := range result {
		if v >= 'a' && v <= 'z' {
			result[i] = v - 32
		} else {
			result[i] = v
		}
	}
	return string(result)
}