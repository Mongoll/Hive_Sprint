package sprint

import "string"

func StrConcat(s1, s2, delim string) string {
	var builder strings.Builder
	builder.WriteString(s1)
	builder.WriteString(delim)
	builder.WriteString(s2)
	return builder.String()
}