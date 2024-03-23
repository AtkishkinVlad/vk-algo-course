package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"bytes"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	zeroCount := 0

	scanner.Scan()
	nStr := scanner.Text()
	n, err := strconv.Atoi(nStr)

	if err != nil || n <= 0 {
		fmt.Println("Error: invalid integer or negative value:", nStr)
		return
	}

	resArr := make([]int, 0, n)
	arr := make([]int, n)

	scanner.Scan()
	inputStr := scanner.Text()
	fields := strings.Fields(inputStr)

	if len(fields) != n {
		fmt.Println("Error: incorrect number of elements:", len(fields), "!=", n)
		return
	}

	for i, field := range fields {
		num, err := strconv.Atoi(field)
		if err != nil {
			fmt.Println("Error: invalid integer value in the array:", num)
			return
		}
		arr[i] = num
	}

	for _, currNum := range arr {
		if currNum != 0 {
			resArr = append(resArr, currNum)
		} else {
			zeroCount += 1
		}
	}

	for zeroCount > 0 {
		resArr = append(resArr, 0)
		zeroCount -= 1
	}

	var buf bytes.Buffer

	for i, n := range resArr {
		if i != 0 {
			buf.WriteString(" ")
		}
		buf.WriteString(strconv.Itoa(n))
	}

	fmt.Println(buf.String())
}
