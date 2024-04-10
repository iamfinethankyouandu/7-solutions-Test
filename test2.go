package main

import (
	"fmt"
	"math"
	"slices"
	"strings"
)

// ข้อนี้ยากจังครับจัดเลยพี่
func Decode(txt string) string {
	arr := strings.Split(txt, "")
	result := []int{0}
	lastChar := ""

	for _, c := range arr {
		if c == "L" {
			result = append(result, result[len(result)-1]-1)
			lastChar = c
		} else if c == "R" {
			if result[len(result)-1]+2 > 2 {
				result[len(result)-1]--
				if lastChar == "=" {
					result[len(result)-2]--
				}
				result = append(result, result[len(result)-1]+1)
				lastChar = c
				continue
			}
			if result[len(result)-1] < 0 {
				result = append(result, result[len(result)-1]+1)
				lastChar = c
				continue
			}
			result = append(result, result[len(result)-1]+2)
			lastChar = c
		} else {
			result = append(result, result[len(result)-1])
			lastChar = c
		}
	}

	if min := slices.Min(result); min < 0 {
		positiveNumber := int(math.Abs(float64(min)))
		for i := range result {
			result[i] += positiveNumber
		}
	}

	return strings.Trim(strings.Replace(fmt.Sprint(result), " ", "", -1), "[]")
}

func incrementEachElement(numbers *[]int) {
	for i := range *numbers {
		(*numbers)[i]--
	}
}
