# StatsCalc

StatsCalc is a command-line program written in Go that calculates basic statistical metrics such as mean, median, mode, and standard deviation for a given set of numbers.

## Features

- **Mean**: Calculates the average of the input numbers.
- **Median**: Finds the middle value of the sorted input numbers.
- **Mode**: Determines the most frequently occurring value(s).
- **Standard Deviation (SD)**: Measures the amount of variation or dispersion in the input numbers.

## Usage

### Input
- The program accepts a list of integers from the standard input. Numbers should be within the range of -100000 to 100000.
- To end input, press `Ctrl+D`.

### Flags
The program supports optional flags to output specific statistical metrics:

- `--mean`: Output the mean.
- `--median`: Output the median.
- `--mode`: Output the mode.
- `--sd`: Output the standard deviation.

### Example
```sh
$ ./statscalc
Введите числа в диапазоне от -100000 до 100000
Чтобы завершить ввод: Ctrl+D
10
20
30
40
50
^D
______________
Mean: 30.00
Median: 30.00
Mode: 10.00
SD: 14.14