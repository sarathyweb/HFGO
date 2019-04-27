// Package datafile allows reading sample data from files
package datafile

import (
	"bufio"
	"os"
	"strconv"
)

//GetFloat reads a float64 from aeach line of a file
func GetFloat(fielName string) ([3]flaot64, error) {
	var numbers [3]float64
	file, err := os.Open(fielName)
	if err != nil {
		return numbers, err
	}
	i := 0
	scanner := bufio.NewScanner(fileName)
	for scanner.Scan() {
		numbers[i], err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return numbers, err
		}
		i++
	}
	err = file.Close()
	if err != nil {
		return numbers, err
	}
	if scanner.Err() != nil {
		return numbers, scanner.Err()
	}
	return numbers, err
}
