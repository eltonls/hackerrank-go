package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

/*
 * Complete the 'plusMinus' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func plusMinus(arr []int32) {
    // Write your code here
    var arrSize int = len(arr)
	var positives int
	var negatives int
	var zeroes int
	
	// Find the number of negatives/positives/zeroes
    for i := 0; i < arrSize; i++ {
		if arr[i] > 0 {
			positives += 1;	
		} else if arr[i] < 0 {
			negatives += 1;
		} else {
			zeroes += 1;
		}
	}

	// iterate to find the ratios and print:
	for k := 0; k < 3; k++ {
		switch k {
		case 0:
			var ratio float64 = float64(positives) / float64(arrSize) 
			fmt.Printf("%.6f\n",ratio);
		case 1:
			var ratio float64 = float64(negatives) / float64(arrSize)
			fmt.Printf("%.6f\n",ratio)
		case 2:
			var ratio float64 = float64(zeroes) / float64(arrSize)
			fmt.Printf("%.6f\n",ratio)
		}
	}
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    n := int32(nTemp)

    arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var arr []int32

    for i := 0; i < int(n); i++ {
        arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
        checkError(err)
        arrItem := int32(arrItemTemp)
        arr = append(arr, arrItem)
    }

    plusMinus(arr)
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
