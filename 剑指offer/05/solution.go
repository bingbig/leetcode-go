package solution

import (
	"strings"
)

func replaceSpace(s string) string {
	return strings.ReplaceAll(s, " ", "%20")
}
