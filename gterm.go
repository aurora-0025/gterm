package main

import (
	"bufio"

	"fmt"

	"io"

	"os"

	"strings"

	"github.com/mazznoer/csscolorparser"
)

//var out []string

func generateGradient(colorSlice []float64, length float64) []float64 {
	fmt.Println(length)
	result := []float64{}
	var dataLength = (length - 2) / (float64(len(colorSlice) - 1))
	fmt.Println(dataLength)

	if (length - 2) == float64(len(colorSlice)) {
		result = append(result, colorSlice...)
	} else {
		for index := range colorSlice {
			if index == len(colorSlice)-1 {
				break
			}
			newColorSlice := []float64{colorSlice[index], colorSlice[index+1]}
			if newColorSlice[0] == newColorSlice[1] {
				for i := 0; i < int(dataLength); i++ {
					result = append(result, newColorSlice[0])
				}
			} else {
				//Arithmetic Progression Taught in school finally came to use
				var commonDifference = (newColorSlice[len(newColorSlice)-1] - newColorSlice[0]) / float64(dataLength)
				var color = newColorSlice[0]
				for int(color) != int(newColorSlice[(len(newColorSlice)-1)]) {
					color = color + commonDifference
					result = append(result, color)
				}
			}
		}
	}
	return result
}

func main() {
	//get all args
	args := os.Args[1:]
	var colors []string
	for _, arg := range args {
		colors = append(colors, string(arg))
	}

	info, _ := os.Stdin.Stat()

	if (info.Mode() & os.ModeCharDevice) != 0 {
		fmt.Println("this is a pipe app")
		fmt.Println("eg: echo 'toby rox' | gterm orange red purple")
		return
	}

	reader := bufio.NewReader(os.Stdin)
	var output []rune

	for {
		input, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			if err != io.EOF {
				fmt.Println(err)
				return
			}
			break
		}
		output = append(output, input)
	}

	var rSlice []float64
	var gSlice []float64
	var bSlice []float64

	for _, color := range colors {
		c, err := csscolorparser.Parse(color)
		if err != nil {
			fmt.Println(err)
			return
		}
		var r, g, b, _ = c.RGBA255()
		rSlice = append(rSlice, float64(r))
		gSlice = append(gSlice, float64(g))
		bSlice = append(bSlice, float64(b))

		fmt.Println(rSlice)
		fmt.Println(gSlice)
		fmt.Println(bSlice)

	}

	chars := strings.Split(string(output), "")
	var length float64 = float64(len(chars))

	fmt.Println(generateGradient(rSlice, length))
	// var gGrad = generateGradient(gSlice, length)
	// var bGrad = generateGradient(bSlice, length)

	// fmt.Println(rGrad)

	// for i := 0; i < len(rGrad); i++ {
	// 	var red, green, blue = uint8(rGrad[i]), uint8(gGrad[i]), uint8(bGrad[i])
	// 	var ele = chars[i]
	// 	out = append(out, fmt.Sprintf("\x1b[38;2;%d;%d;%dm%s\x1b[0m", red, green, blue, ele))
	// }
	// fmt.Println(strings.Join(out[:], ""))
}
