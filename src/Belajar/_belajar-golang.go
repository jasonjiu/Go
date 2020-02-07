package main

import (
	"fmt"
	"strings"
)

func main() {
	var firstName = "jason"
	var lastName = "anggarah"
	var notUsed = "satu"
	_ = notUsed
	_ = firstName
	_ = lastName

	fmt.Println("your fav child")
	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(firstName + " " + lastName)
	closure()
	filterMain()
}

func closure() {
	var getMinMax = func(n []int) (int, int) {
		var min, max int
		for i, e := range n {
			switch {
			case i == 0:
				max, min = e, e
			case e > max:
				max = e
			case e < min:
				min = e
			}
		}
		return min, max
	}
	var number = []int{2, 3, 4, 3, 4, 2, 3}
	var min, max = getMinMax(number)
	fmt.Printf("data : %v\nmin : %v\nmax : %v\n", number, min, max)
}

func filter(data []string, callback func(string) bool) []string {
	var result []string
	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}
	return result
}

func filterMain() {
	var data = []string{"wick", "jason", "ethan"}
	var dataContainsO = filter(data, func(each string) bool {
		return strings.Contains(each, "o")
	})

	var dataLength5 = filter(data, func(each string) bool {
		return len(each) == 5
	})

	fmt.Println("data asli \t\t:", data)
	fmt.Println("filter ada huruf O \t:", dataContainsO)
	fmt.Println("filter jumlah huruf 5 \t:", dataLength5)
}
