package main

import (
	"bytes"
	"strconv"
)

func addBinary(a string, b string) string {
	i := len(a) - 1
	j := len(b) - 1
	carry := 0
	var buffer bytes.Buffer

	for i >= 0 || j >= 0 {
		sum := carry
		if i >= 0 {
			parta := getInt(i, a)
			sum += parta
		}
		if j >= 0 {
			partb := getInt(j, b)
			sum += partb
		}

		result := sum % 2
		carry = sum / 2

		buffer.WriteString(strconv.Itoa(result))
		i--
		j--
	}

	// note: need to check if there is any remaining carry
	if carry != 0 {
		buffer.WriteString("1")
	}

	return reverseStr(buffer.String())
}

func reverseStr(s string) string {
	sBytes := []byte(s)
	left := 0
	right := len(s) - 1

	for left < right {
		sBytes[left], sBytes[right] = sBytes[right], sBytes[left]
		left++
		right--
	}

	return string(sBytes)
}

func getInt(index int, s string) int {
	return int(s[index] - '0')
}
