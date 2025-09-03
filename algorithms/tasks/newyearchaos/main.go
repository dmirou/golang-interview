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
 * Complete the 'minimumBribes' function below.
 *
 * q has items (positions) from 1 to n
 * [2 1 5 3 4] -> 3 -> 12~21, 45~54, 35~53
 * [2 5 1 3 4] -> Too chaotic
 */
func minimumBribes(q []int32) {
	// [1]
	if len(q) == 1 {
		fmt.Println(0)
		return
	}

	// [1 2] -> 0
	// 1 > 2, false, go next
	// totalB = 0, i = 0, curB = 0
	// 	j = 1, j < 2
	// 		q[1] > q[2], false
	// totalB = 0

	// [2 1] -> 1
	//
	// totalB = 0, i = 0, curB = 0
	// 	j = 1, j < 2
	// 		q[1] > q[2], true
	//		curB = 1
	//		curB < 2, true
	// totalB = 0

	curB := 0
	totalB := 0

	for i := 0; i < len(q)-1; i++ {
		// if position wasn't changed, don't need to compare
		// with other items
		if q[i] == int32(i+1) && totalB == 0 {
			continue
		}

		if q[i]-int32(i+1) > 2 {
			fmt.Println("Too chaotic")
			return
		}

		curB = 0
		for j := i + 1; j < len(q); j++ {
			if q[i] > q[j] {
				curB++
				if curB > 2 {
					fmt.Println("Too chaotic")
					return
				}
			}
		}
		totalB += curB
	}

	fmt.Println(totalB)

	// [1 2 3]
	// 1 > 2, false, go next
	// 2 > 3, false, go next

	// [2 1 3]
	// 2 > 1, true, bribes++,
	// 2 > 3, false, go next
	// 1 > 3, false, go next

	// [3, 1, 2]
	// 3 > 1, true, bribes++
	// 3 > 2, true, bribes++, go next
	// 1 > 2, false, go next

	// [3, 2, 1]
	// 3 > 2, true, bribes++
	// 3 > 1, true, bribes++, go next
	// 2 > 1, true, bribes++, go next

	// [4, 3, 1, 2]
	// 4 > 3, true, bribes++
	// 4 > 1, true, bribes++
	// 4 > 2, true, bribes ==2, too chaotic

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		n := int32(nTemp)

		qTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		var q []int32

		for i := 0; i < int(n); i++ {
			qItemTemp, err := strconv.ParseInt(qTemp[i], 10, 64)
			checkError(err)
			qItem := int32(qItemTemp)
			q = append(q, qItem)
		}

		minimumBribes(q)
	}
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
