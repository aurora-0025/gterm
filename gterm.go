package main

import (
	"bufio"

	"fmt"

	"io"

	"os"

	"strings"

	"github.com/mazznoer/csscolorparser"
)

var out []string

func sumAray(array []float64) float64 {
	var result float64 = 0
	for _, v := range array {
		result += v
	}
	return result
}

func splitLength(x, n int) []float64 {
	dataLength := []float64{}
	if x < n {
		fmt.Println("error")
	} else if n == 1 {
		dataLength = append(dataLength, float64(x))
	} else if x%n == 0 {
		for i := 0; i < n; i++ {
			dataLength = append(dataLength, float64(x/n))
		}
	} else {
		var zp = n - (x % n)
		var pp = x / n
		for i := 0; i < n; i++ {
			if i >= zp {
				dataLength = append(dataLength, float64(pp+1))
			} else {
				dataLength = append(dataLength, float64(pp))
			}

		}
	}
	return dataLength
}

func generateGradient(colorSlice []float64, length int) []float64 {
	result := []float64{}
	var dataLength = (splitLength(length-1, (len(colorSlice) - 1)))
	result = append(result, colorSlice[0])

	if (length - 1) == (len(colorSlice)) {
		result = append(result, colorSlice...)
	} else {
	EXIT:
		for index := range colorSlice {
			if index == len(colorSlice)-1 {
				break EXIT
			} else {
				newColorSlice := []float64{colorSlice[index], colorSlice[index+1]}
				if newColorSlice[0] == newColorSlice[1] {
					for i := 0; float64(i) < float64(dataLength[index]); i++ {
						result = append(result, newColorSlice[0])
					}
				} else {
					//Arithmetic Progression Taught in school finally came to use
					var commonDifference float64 = float64(float64(newColorSlice[len(newColorSlice)-1]-newColorSlice[0]) / float64(dataLength[index]))
					var color = newColorSlice[0]
					for int(color+0.5) != int(newColorSlice[(len(newColorSlice)-1)]) {
						color = (color + commonDifference)
						result = append(result, color)
					}
				}
			}
		}
		for (len(result) - 1) != int(sumAray(dataLength)) {
			result = append(result, result[len(result)-1])
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

	}

	chars := strings.Split(string(output), "")
	var length int = (len(chars))

	var rGrad = generateGradient(rSlice, length)
	var gGrad = generateGradient(gSlice, length)
	var bGrad = generateGradient(bSlice, length)

	for i := 0; i < len(rGrad); i++ {
		var red, green, blue = uint8(rGrad[i]), uint8(gGrad[i]), uint8(bGrad[i])
		var ele = chars[i]
		out = append(out, fmt.Sprintf("\x1b[38;2;%d;%d;%dm%s\x1b[0m", red, green, blue, ele))
	}
	fmt.Println(strings.Join(out[:], ""))
}
