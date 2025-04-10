package floz

import "strings"

func parsePath(path string) []string {
	parts := make([]string, 0)
	vpath := strings.Split(path, "/")
	for _, v := range vpath {
		if v != "" {
			parts = append(parts, v)
			if v[0] == '*' {
				break
			}
		}
	}
	return parts
}
