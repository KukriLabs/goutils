package arrays

import "strings"

func TrimEmptyStrings(input []string) []string {
	output := []string{}
	if len(input) == 0 {
		return input
	}
	for _, s := range input {
		if strings.TrimSpace(s) != "" {
			output = append(output, s)
		}
	}
	return output
}
