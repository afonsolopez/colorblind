package colorblind

import "math"

func luminance(a, b float64) float64 {
	l1 := math.Max(a, b)
	l2 := math.Min(a, b)
	return (l1 + 0.05) / (l2 + 0.05)
}

//ScoreRGB calculate the score between two RGB triplets
func ScoreRGB(a [4]int, b [4]int) float64 {
	return luminance(Relativeluminance(a), Relativeluminance(b))
}

//ScoreHex calculate the score between two hexadecimal strings
func ScoreHex(a, b string) float64 {
	return ScoreRGB(HextoRGB(a), HextoRGB(b))
}

//Grading the scores based on the WCAG criteria
func Grading(score float64) string {
	if score >= 7 {
		return "AAA"
	} else if score >= 4.5 {
		return "AA"
	} else if score >= 3 {
		return "A"
	}

	return "FAIL"
}
