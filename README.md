# Colorblind

This package implements the WCAG 2.0 grading and score system. 

This project was initially a mix of another two NodeJS projects ports [relative-luminance](https://github.com/tmcw/relative-luminance) and [wcag-contrast](https://github.com/tmcw/wcag-contrast) to Golang both created by [Tom MacWright](https://github.com/tmcw). 


## Installation

Use go get.

```bash
go get github.com/afonsolopez/colorblind
```

Then import the validator package into your own code.
```bash
import "github.com/afonsolopez/colorblind"
```

## Usage

```go
# Convert hexadecimal color codes to RGB values
colorblind.HextoRGB("#DE605F") # [0 0 0 255] 

# Get the contrast score based on a RGB value
colorblind.ScoreRGB([4]int{0, 0, 0, 255}, [4]int{255, 255, 255, 255}) # 21

# Get the contrast score based on a hexadecimal color code 
colorblind.ScoreHex("#663399", "#FFF") # 21

# Get the WCAG 2.0 grading based on a contrast score
colorblind.Grading(21) # AAA
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.
