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

func convert(v string) int {
	value, _ := strconv.Atoi(v)
	return value
}

const (
	FORWARD string = "FORWARD"
	DOWN    = "DOWN"
	UP      = "UP"
)

type Submarine struct {
	depth int
	horizontal int
}

func main() {
	data := strings.Split(report(), "\n")
	dataSplit := make([][]string, 0)
	submarine := Submarine{0, 0}

	for i := 0; i < len(data)-1; i++ {
		temp := strings.Split(data[i], " ")
		v := []string{strings.ToUpper(temp[0]), temp[1]}
		dataSplit = append(dataSplit, v)
	}


	for i := 0; i < len(dataSplit); i++ {
		command := dataSplit[i][0]
		value := dataSplit[i][1]

		if command == FORWARD {
			submarine.horizontal += convert(value)
		} else if command == UP {
			submarine.depth -= convert(value)
		} else if command == DOWN {
			submarine.depth += convert(value)
		}
	}

	fmt.Println(submarine.depth*submarine.horizontal)
}
