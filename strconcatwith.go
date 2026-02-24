package main

func StrConcatWith(strs []string, sep string) string {
	var result []byte
	for i, s := range strs {
		result = append(result, s...)
		if i < len(strs)-1 {
			result = append(result, sep...)
		}
	}
	return string(result)
}