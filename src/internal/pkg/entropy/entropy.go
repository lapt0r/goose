package entropy

import
(
	"math"
)

func GetCharacterCount(text string) map[string]int {
	var result = make(map[string]int)
	for char := range text {
		var count = result[string(char)]
		count += 1
		result[string(char)] = count
	}
	return result
}

func GetPValues(text string) []float64 {
	var charmap = GetCharacterCount(text)
	var totalcharcount int = len(text)
	var result = make([]float64, totalcharcount)

	var index = 0
	for char := range charmap {
		var p_val = float64(charmap[char]) / float64(totalcharcount)
		result[index] = p_val
		index += 1
	}
	return result
}

func ComputeEntropy(text string) float64 {
	var pvals = GetPValues(text)
	var entropy = float64(0)
	for _,p_val := range pvals {
		entropy -= (p_val * math.Log(p_val))
	}
	return entropy
}