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
	ErrNotInteger = errors.New("ошибка ввода: невалидные данные")
	ErrOutOfRange = errors.New("ошибка ввода: число вне диапазона (-100000 до 100000)")
	ErrEmptyInput = errors.New("ошибка ввода: введено пустое значение")
)

func Scanner(r io.Reader) ([]int, error) {
	// func Scanner() {
	var slice []int
	scanner := bufio.NewScanner(r)
	fmt.Println("Введите числа, каждое c новой строки (нажмите Ctrl+D для завершения)")
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
		return nil, errors.New("ошибка чтения ввода")
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
	meanFlag := flag.Bool("mean", false, "Рассчитать среднее значение")
	modeFlag := flag.Bool("mode", false, "Рассчитать моду")
	medianFlag := flag.Bool("median", false, "Рассчитать медиану")
	sdFlag := flag.Bool("sd", false, "Рассчитать стандартное отклонение")

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
