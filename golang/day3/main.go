package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func report() string {
	data, err := ioutil.ReadFile("./data.txt")
	if err != nil {
		panic(err)
	}

	return string(data)
}

func generateNewArray(arr [][]int) [][]int {
	newArr := make([][]int, len(arr[0]))

	for i := 0; i < len(arr[0]); i++ {
		var temp []int
		for j := 0; j < len(arr); j++ {
			temp = append(temp, arr[j][i])
		}
		newArr[i] = temp
	}
	return newArr
}

func bitOperation(arr [][]int) (gamma, epsilon int) {
	for _, v := range arr {
		count := map[int]int{0: 0, 1: 0}
		for _, val := range v {
			count[val]++
		}
		if count[0] > count[1] {
			gamma = gamma<<1 | 0
			epsilon = epsilon<<1 | 1
		} else {
			gamma = gamma<<1 | 1
			epsilon = epsilon<<1 | 0
		}
	}

	return gamma, epsilon
}

func main() {
	data := strings.Split(report(), "\n")
	var arr [][]int

	for _, row := range data {
		if row == "" {
			continue
		}
		var newRow []int
		for _, val := range strings.Split(row, "") {
			num, _ := strconv.Atoi(val)
			newRow = append(newRow, num)
		}
		arr = append(arr, newRow)
	}

	if len(arr) == 0 {
		fmt.Println("No data found.")
		return
	}

	newArr := generateNewArray(arr)
	gamma, epsilon := bitOperation(newArr)
	powerConsumption := gamma * epsilon

	fmt.Printf("The power consumption is %d\n", powerConsumption)
}
