package main

func IsLower(s string) bool {
	for _, v := range s {
		if v < 'a' || v > 'z' {
			if v < 'A' || v > 'Z' {
				return false
			}
		}

	}
	return true
}
