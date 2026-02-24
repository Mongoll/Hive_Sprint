package main

func StrSplitBy(s, sep string) []string {

	if s == "" {
		return []string{}
	}

	if sep == "" {
		return []string{}
	}

	var result []string
	current := ""
	i := 0

	for i < len(s) {

		if i+len(sep) <= len(s) && s[i:i+len(sep)] == sep {

			result = append(result, current)
			current = ""
			i += len(sep)
			continue
		}

		current += string(s[i])
		i++
	}

	result = append(result, current)

	return result
}
