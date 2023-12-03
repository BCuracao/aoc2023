package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	sum := 0
	firstNum, lastNum := "", ""

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		for _, char := range line {
			if unicode.IsNumber(rune(char)) {
				firstNum = string(char)
				break
			}
		}
		for i := len(line) - 1; i >= 0; i-- {
			if unicode.IsNumber(rune(line[i])) {
				lastNum = string(line[i])
				break
			}
		}

		conc := firstNum + lastNum
		temp, _ := strconv.Atoi(conc)
		sum += temp
	}

	fmt.Println(sum)
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file: ", err)
	}
}
