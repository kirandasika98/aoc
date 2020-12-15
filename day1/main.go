package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	nums := []int{}
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, num)
	}

	fmt.Println(result(find2020Trio(nums)))
}

func find2020Pair(nums []int) (num1 int, num2 int) {
	for _, i := range nums {
		num1 = i
		for _, j := range nums {
			num2 = j

			if isMatch(num1, num2) {
				return
			}
		}
	}
	return
}

func isMatch(nums ...int) bool {
	sum := 0
	for _, i := range nums {
		sum += i
	}
	return sum == 2020
}

func find2020Trio(nums []int) (num1, num2, num3 int) {
	for _, i := range nums {
		num1 = i
		for _, j := range nums {
			num2 = j
			for _, k := range nums {
				num3 = k
				if isMatch(num1, num2, num3) {
					return
				}
			}
		}
	}
	return
}

func result(nums ...int) (result int) {
	result = 1
	for _, i := range nums {
		result *= i
	}
	return
}
