package IC

import (
	"regexp"
	"strings"
)

type IC struct {
	totalLength     int
	characterCounts [26]int
	cleanerRegex    *regexp.Regexp
}

// Compute the Index o Coincidence of the given IC in one shot
func OneShot(text string) float32 {
	re := regexp.MustCompile("[^[:alpha:]]")
	lowered := strings.ToLower(re.ReplaceAllString(text, ""))
	sum := 0

	for char := 'a'; char <= 'z'; char++ {
		count := strings.Count(lowered, string(char))
		sum += count * (count - 1)
	}

	n := len([]rune(lowered))
	if n == 0 {
		return 0.0
	}

	return float32(sum) / float32(n*(n-1))
}

// Reset the internal structures for continuous calculations
func Begin() IC {
	return IC{cleanerRegex: regexp.MustCompile("[^[:alpha:]]")}
}

// Update continuous calculation with new data
func Update(ic *IC, text string) {
	lowered := strings.ToLower(ic.cleanerRegex.ReplaceAllString(text, ""))

	for char := 'a'; char <= 'z'; char++ {
		ic.characterCounts[char-'a'] += strings.Count(lowered, string(char))
	}

	ic.totalLength += len([]rune(lowered))
}

// Compute the final IC for continuous calculation
func Compute(ic *IC) float32 {
	if ic.totalLength < 2 {
		return 0.0
	}

	sum := 0
	for char := 'a'; char <= 'z'; char++ {
		sum += ic.characterCounts[char-'a'] * (ic.characterCounts[char-'a'] - 1)
	}

	return float32(sum) / float32(ic.totalLength*(ic.totalLength-1))
}
