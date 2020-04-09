package reflectorfilter

import (
	"math"
	"regexp"
	"strings"

	"gopkg.in/vmarkovtsev/go-lcss.v1"
)

// IsReflected : Checks to see if a potential match is a partially-reflected string
func IsReflected(potentialMatch string) bool {
	matcher, _ := regexp.Compile("(:|:=|->|=)")
	potentialMatch = strings.ToLower(strings.TrimSpace(potentialMatch))
	operator := matcher.FindString(potentialMatch)
	if operator != "" {
		substrings := strings.SplitN(potentialMatch, operator, 2)
		left, right := substrings[0], substrings[1]
		common := lcss.LongestCommonSubstring([]byte(left), []byte(right))
		reflectorPercentage := float64(len(common)) / math.Min(float64(len(left)), float64(len(right)))
		return reflectorPercentage > 0.7
	}
	return false
}
