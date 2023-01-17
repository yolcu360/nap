package nap

import (
	"regexp"
	"strings"
)

var re = regexp.MustCompile(`FOR UPDATE|LOCK`)

// ForceMasterNode some queries only works with master. This func helps for that.
func ForceMasterNode(query string) bool {
	return re.MatchString(strings.ToUpper(query))
}
