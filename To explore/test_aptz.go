package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	// Challenge 1, find the neeares beetwen the pair and mix them, e.g
	// from {{0, 1}, {2, 3}, {4, 6}, {6, 10}, {11, 20}, {30, 40}}
	// to {{0, 1}, {2, 3}, {4, 10}, {11, 20}, {30, 40}}
	fmt.Println(nearest([][]int{{0, 1}, {2, 3}, {4, 6}, {6, 10}, {11, 20}, {30, 40}}))

	// Challenge 2, remove dups
	fmt.Println(removeDuplicates([]int{1, 100, 100, 3, 5, 3, 4, 4, 4, 4, 4, 5, 7, 3, 6, 9, 7}))

	// Challenge 3, Convert a multidimentional data structure to a one dimention one. e.g
	// From map[string]interface{}{
	//			"a": map[string]interface{}{
	//				"b": map[string]interface{}{
	//					"c": 42,
	//					"d": "foo",
	//				},
	//				"d": "bar",
	//			},
	//			"e": "baz",
	//		}
	// to
	//
	// map[a_b_c:42 a_b_d:foo a_d:bar e:baz]
	//
	fmt.Println(
		mapKeys(map[string]interface{}{
			"a": map[string]interface{}{
				"b": map[string]interface{}{
					"c": 42,
					"d": "foo",
				},
				"d": "bar",
			},
			"e": "baz",
		}),
	)

	// Challenge 4 check if palindromes
	fmt.Println(
		palindromes([]string{
			"ana",
			"hannah",
			"pedro",
			"emoome",
		}),
	)
}


func palindromes(names []string) (output []string) {
	for _, name := range names {
		if isPalindrome(name) {
			output = append(output, name)
		}
	}
	return output
}
func isPalindrome(name string) bool {
	keys := strings.Split(name, "")
	total := len(keys)
	half := int(math.Floor(float64(total / 2)))

	for i := 0; i < half; i++ {
		if keys[i] != keys[total-1-i] {
			return false
		}
	}
	return true
}

func mapKeys(input map[string]interface{}) map[string]interface{} {
	return mapKeysSingle(input, []string{}, map[string]interface{}{})
}
func mapKeysSingle(input map[string]interface{}, keys []string, output map[string]interface{}) map[string]interface{} {

	for key, item := range input {
		newKeys := append(keys, key)
		switch item.(type) {
		case map[string]interface{}:
			output = mapKeysSingle(item.(map[string]interface{}), newKeys, output)
		default:
			output[strings.Join(newKeys, "_")] = item
		}

	}
	return output
}

func removeDuplicates(list []int) (output []int) {
	counterMap := map[int]int{}
	for _, item := range list {
		value := counterMap[item]
		if value == 1 {
			output = append(output, item)
		}
		counterMap[item] = counterMap[item] + 1
	}
	return output
}

//[[0 1] [2 3] [4 10] [11 20] [30 40]]
func nearest(list [][]int) (output [][]int) {

	for i := 0; i < len(list); i++ {
		item := list[i]
		toIndex := getIndex(list, i)
		i = toIndex
		output = append(output, []int{item[0], list[toIndex][1]})
	}

	return output
}

func getIndex(list [][]int, index int) int {
	for i := index; i < len(list)-1; i++ {
		from := list[i]
		to := list[i+1]
		if from[1] >= to[0] {
			continue
		}
		return i
	}
	return index
}

//
//func compareTo2(list [][]int, index int) int {
//
//	prevIndex := index
//	i := index + 1
//	total := len(list)
//	if total == i {
//		return index
//	}
//	for ; i < total; i++ {
//		from := list[prevIndex]
//		to := list[i]
//		if from[1] >= to[0] {
//			prevIndex = i
//			continue
//		}
//		break
//	}
//	return i - 1
//}