package main

import (
    "fmt"
    "bufio"
    "os"
    "strconv"
    "strings"
	"bytes"
)

func insertionSort(arr []int) []int {
    for i := 1; i < len(arr); i++ {
        key := arr[i]
        j := i - 1
        for ; j >= 0 && arr[j] > key; j-- {
            arr[j + 1] = arr[j]
        }
        arr[j + 1] = key
    }
    return arr
}

func getInputs() []int {
    reader := bufio.NewReader(os.Stdin)
    inputStr, _ := reader.ReadString('\n')
    n, _ := strconv.Atoi(strings.TrimSpace(inputStr))

    arr := make([]int, n)
    inputStr, _ = reader.ReadString('\n')
    inputs := strings.Split(inputStr, " ")
    for i := 0; i < n; i++ {
        arr[i], _ = strconv.Atoi(inputs[i])
    }
    return arr
}

func printArray(arr []int) {
    var buf bytes.Buffer

	for i, n := range arr {
		if i != 0 {
			buf.WriteString(" ")
		}
		buf.WriteString(strconv.Itoa(n))
	}

	fmt.Println(buf.String())
}

func main() {
    unsortedArr := getInputs()
    sortedArr := insertionSort(unsortedArr)
    printArray(sortedArr)
}