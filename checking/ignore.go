package checking

import (
	"regexp"
	"strings"
)

var ignoreCommentRegex = regexp.MustCompile(`^leakcheck\s?:\s?ignore$`)

func ContainsIgnoreComment(line string) bool {
	_, comment, found := strings.Cut(line, "//")
	if !found {
		return false
	}
	comment = strings.TrimSpace(comment)

	matched := ignoreCommentRegex.MatchString(comment)
	return matched
}
