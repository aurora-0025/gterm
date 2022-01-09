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

func splitLength(x, n int) []int {
	dataLength := []int{}
	if x < n {
		fmt.Println("error")
	} else if x%n == 0 {
		for i := 0; i < n; i++ {
			dataLength = append(dataLength, x/2)
		}
	} else {
		var zp = n - (x % n)
		var pp = x / n
		for i := 0; i < n; i++ {
			if i >= zp {
				dataLength = append(dataLength, (pp + 1))
			} else {
				dataLength = append(dataLength, pp)
			}

		}
	}
	return dataLength
}

func generateGradient(colorSlice []float32, length float32) []float32 {
	fmt.Println(length - 1)
	result := []float32{}
	var dataLength = splitLength(int(length-2), (len(colorSlice) - 1))
	result = append(result, colorSlice[0])

	fmt.Println(dataLength)

	if (length - 1) == float32(len(colorSlice)) {
		result = append(result, colorSlice...)
	} else {
		for index := range colorSlice {
			if index == len(colorSlice)-1 {
				break
			}
			newColorSlice := []float32{colorSlice[index], colorSlice[index+1]}
			if newColorSlice[0] == newColorSlice[1] {
				for i := 0; i < int(dataLength[index]); i++ {
					result = append(result, newColorSlice[0])
				}
			} else {
				//Arithmetic Progression Taught in school finally came to use
				var commonDifference = (newColorSlice[len(newColorSlice)-1] - newColorSlice[0]) / float32(dataLength[index])
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

	var rSlice []float32
	var gSlice []float32
	var bSlice []float32

	for _, color := range colors {
		c, err := csscolorparser.Parse(color)
		if err != nil {
			fmt.Println(err)
			return
		}
		var r, g, b, _ = c.RGBA255()
		rSlice = append(rSlice, float32(r))
		gSlice = append(gSlice, float32(g))
		bSlice = append(bSlice, float32(b))

		fmt.Println(rSlice)
		fmt.Println(gSlice)
		fmt.Println(bSlice)

	}

	chars := strings.Split(string(output), "")
	var length float32 = float32(len(chars))

	var rGrad = generateGradient(rSlice, length)
	var gGrad = generateGradient(gSlice, length)
	var bGrad = generateGradient(bSlice, length)

	// fmt.Println(rGrad)

	for i := 0; i < len(rGrad); i++ {
		var red, green, blue = uint8(rGrad[i]), uint8(gGrad[i]), uint8(bGrad[i])
		var ele = chars[i]
		out = append(out, fmt.Sprintf("\x1b[38;2;%d;%d;%dm%s\x1b[0m", red, green, blue, ele))
	}
	fmt.Println(strings.Join(out[:], ""))
}
