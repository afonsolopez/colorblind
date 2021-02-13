package colorblind

import (
	"regexp"
	"strconv"
	"strings"
)

//HextoRGB convert any Hex color code to RGB
func HextoRGB(hex string) [4]int {
	//Remove hashtag
	re := regexp.MustCompile("#")
	s := re.ReplaceAllString(hex, "")

	//Declare Alpha value
	alpha := 255

	//Handle each number of Hex digits

	//In case of 8 digits Hex
	if len(s) == 8 {
		//New Alpha value is calculated, based on the last 2 digits
		alpha64, _ := strconv.ParseInt(s[6:8], 16, 64)
		alpha = int(alpha64)
		//Assign the first 6 digits as Hex value
		s = s[:6]
	}

	//In case of 4 digits Hex
	if len(s) == 4 {
		//New Alpha value is calculated, based on the last digit
		alpha64, _ := strconv.ParseInt(strings.Repeat(s[3:4], 2), 16, 64)
		alpha = int(alpha64)
		//Assign the first 3 digits as Hex value
		s = s[:3]
	}

	//In case of 3 digits Hex or after a 4 digits conversion
	if len(s) == 3 {
		//Assign 2 times each digit in lecture order as the Hex value
		s = string(s[0]) + string(s[0]) + string(s[1]) + string(s[1]) + string(s[2]) + string(s[2])
	}
	// Convert the Hex values to decimal
	num, _ := strconv.ParseInt(s, 16, 64)
	red := num >> 16
	green := (num >> 8) & 255
	blue := num & 255

	//Create the RGB array
	rgbarray := [4]int{int(red), int(green), int(blue), alpha}

	return rgbarray
}
