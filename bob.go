package bob

import "strings"

func MyFunction(input string) interface{} {
	switch {
	case strings.ToUpper(input) == input && strings.HasSuffix(input, "?"):
		return "Calm down, I know what I'm doing\n"
	case strings.ToUpper(input) == input:
		return "Whoa, chill out\n"
	case strings.HasSuffix(input, "?"):
		return "sure\n"
	case " " == input, "\n" == input, len(input) == 0:
		return "Fine. Be that way!\n"
	default:
		return "Whatever!!!!\n"
	}
}
