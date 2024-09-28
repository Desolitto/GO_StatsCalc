package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
)

var (
	ErrNotInteger = errors.New("input error: invalid data")
	ErrOutOfRange = errors.New("input error: number out of range (-100000 to 100000)")
	ErrEmptyInput = errors.New("input error: empty value entered")
)

func Scanner(r io.Reader) ([]int, error) {
	var slice []int
	scanner := bufio.NewScanner(r)
	fmt.Println("Enter numbers, each on a new line (press Ctrl+D to finish)")
	for scanner.Scan() {
		line := scanner.Text()
		input, err := strconv.Atoi(line)
		if err != nil {
			return nil, ErrNotInteger
		}
		if input <= -100000 || input >= 100000 {
			return nil, ErrOutOfRange
		}
		slice = append(slice, input)
	}
	if len(slice) == 0 {
		return nil, ErrEmptyInput
	}
	if err := scanner.Err(); err != nil {
		return nil, errors.New("input reading error")
	}
	return slice, nil
}

func mean(slice []int) float64 {
	var mean float64
	for _, elem := range slice {
		mean += float64(elem)
	}
	return mean / float64(len(slice))
}

func mode(slice []int) int {
	frequency := make(map[int]int)
	maxFrequency := 0
	mode := slice[0]

	for _, num := range slice {
		frequency[num]++
		if frequency[num] > maxFrequency {
			maxFrequency = frequency[num]
			mode = num
		}
	}

	return mode
}

func median(slice []int) float64 {
	sort.Ints(slice)
	mid := len(slice) / 2
	if len(slice)%2 == 1 {
		return float64(slice[mid])
	} else {
		return (float64(slice[mid]) + float64(slice[mid-1])) / 2.0
	}
}

func sd(slice []int) float64 {
	mean := mean(slice)
	sumSquares := 0.0
	for _, num := range slice {
		diff := num - int(mean)
		sumSquares += float64(diff * diff)
	}
	mean = sumSquares / float64(len(slice)-1)
	return math.Sqrt(mean)
}

func applyDefaultFlags(meanFlag, medianFlag, modeFlag, sdFlag *bool) {
	if !*meanFlag && !*medianFlag && !*modeFlag && !*sdFlag {
		*meanFlag = true
		*medianFlag = true
		*modeFlag = true
		*sdFlag = true
	}
}

func main() {
	meanFlag := flag.Bool("mean", false, "Calculate mean value")
	modeFlag := flag.Bool("mode", false, "Calculate mode")
	medianFlag := flag.Bool("median", false, "Calculate median")
	sdFlag := flag.Bool("sd", false, "Calculate standard deviation")

	flag.Parse()

	slice, err := Scanner(os.Stdin)
	if err != nil {
		fmt.Println(err)
		return
	}

	applyDefaultFlags(meanFlag, medianFlag, modeFlag, sdFlag)

	if *meanFlag {
		fmt.Printf("Mean: %.2f\n", mean(slice))
	}
	if *medianFlag {
		fmt.Printf("Median: %.2f\n", median(slice))
	}
	if *modeFlag {
		fmt.Printf("Mode: %d\n", mode(slice))
	}
	if *sdFlag {
		fmt.Printf("SD: %.2f\n", sd(slice))
	}
}
