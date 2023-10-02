package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

func report() string {
	data, err := ioutil.ReadFile("./data.txt")

	if err != nil {
		panic(err)
	}

	return string(data)
}

func main(){
	data := strings.Split(report(), "\n")
	higher := make([]int, 0)

	for i := 1; i < len(data); i++ {
		current, _ := strconv.Atoi(data[i])
		previous, _ := strconv.Atoi(data[i-1])

		if(current > previous){
			higher = append(higher, current)
		}
	}

	fmt.Println(len(higher))
}
