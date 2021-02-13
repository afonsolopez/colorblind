package colorblind

import "math"

const rc float64 = 0.2126
const gc float64 = 0.7152
const bc float64 = 0.0722

const lowc float64 = 1 / 12.92

func adjustGamma(x float64) float64 {
	return math.Pow((x+0.055)/1.055, 2.4)
}

//Relativeluminance is "the relative brightness of any point in a colorspace"
func Relativeluminance(rgb [4]int) float64 {
	rsrgb := float64(rgb[0]) / float64(255)
	gsrgb := float64(rgb[1]) / float64(255)
	bsrgb := float64(rgb[2]) / float64(255)

	var r float64

	if rsrgb <= 0.03928 {
		r = rsrgb * lowc
	} else {
		r = adjustGamma(rsrgb)
	}

	var g float64

	if gsrgb <= 0.03928 {
		g = gsrgb * lowc
	} else {
		g = adjustGamma(gsrgb)
	}

	var b float64

	if bsrgb <= 0.03928 {
		b = bsrgb * lowc
	} else {
		b = adjustGamma(bsrgb)
	}

	return float64(r*rc + g*gc + b*bc)
}
