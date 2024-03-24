package main

import (
    "fmt"
    "bufio"
    "os"
    "strconv"
    "strings"
	"bytes"
)

func mergeSort(items []int) []int {
    if len(items) < 2 {
        return items
    }
    first := mergeSort(items[:len(items)/2])
    second := mergeSort(items[len(items)/2:])
    return merge(first, second)
}

func merge(a []int, b []int) []int {
    final := []int{}
    i := 0
    j := 0
    for i < len(a) && j < len(b) {
        if a[i] < b[j] {
            final = append(final, a[i])
            i++
        } else {
            final = append(final, b[j])
            j++
        }
    }
    for ; i < len(a); i++ {
        final = append(final, a[i])
    }
    for ; j < len(b); j++ {
        final = append(final, b[j])
    }
    return final
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
    sortedArr := mergeSort(unsortedArr)
    printArray(sortedArr)
}