package checking

import (
	"regexp"
	"strings"
)

var ignoreCommentRegex = regexp.MustCompile(`^leakcheck\s?:\s?ignore$`)

func ContainsIgnoreInlineComment(line string) (bool, error) {
	_, comment, found := strings.Cut(line, "//")
	if !found {
		return false, nil
	}
	comment = strings.TrimSpace(comment)

	matched := ignoreCommentRegex.MatchString(comment)
	return matched, nil
}