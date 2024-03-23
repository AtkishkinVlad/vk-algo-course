package main

import (
    "fmt"
    "bufio"
    "os"
    "strconv"
    "strings"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
	res := 0

    scanner.Scan()
    nStr := scanner.Text()
    n, err := strconv.Atoi(nStr)

	if err != nil || n <= 0 {
        fmt.Println("Error: invalid integer or negative value:", nStr)
        return
    }

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

    scanner.Scan()

	elemStr := scanner.Text()
    elem, err := strconv.Atoi(elemStr)

	if err != nil {
        fmt.Println("Error: invalid integer value for 'element':", elemStr)
        return
    }

	for _, currNum := range arr {
        if currNum != elem {
			res += 1
		}
    }

    fmt.Println(res)
}