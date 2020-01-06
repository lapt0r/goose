package entropy

import
(
	"math"
)

func GetCharacterCount(text string) map[rune]int {
	var result = make(map[rune]int)
	for _,char := range text {
		var count = result[char]
		count += 1
		result[char] = count
	}
	return result
}

func GetPValues(text string) []float64 {
	var charmap = GetCharacterCount(text)
	var totalcharcount int = len(text)
	var result = make([]float64, len(charmap))

	var index = 0
	for _, count := range charmap {
		var p_val = float64(count) / float64(totalcharcount)
		result[index] = p_val
		index += 1
	}
	return result
}

func GetShannonEntropy(text string) float64 {
	var pvals = GetPValues(text)
	var entropy = float64(0)
	for _,p_val := range pvals {
		entropy -= (p_val * math.Log(p_val))
	}
	return entropy
}