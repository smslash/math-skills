package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"git/ssengerb/math-skills/formula"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Error: incorrect number of arguments")
	}

	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalln("Error:", err)
	}
	data = append(data, 10)

	var numbers []int
	var str string

	for i := 0; i < len(data); i++ {
		if data[i] != 10 {
			str += string(data[i])
		} else {
			if str == "" {
				continue
			}
			n, err := strconv.Atoi(str)
			if err != nil {
				log.Fatalln("Error: file contains non convertible value")
			}
			numbers = append(numbers, n)
			str = ""
		}
	}

	fmt.Println("Average:", formula.Average(numbers))
	fmt.Println("Median:", formula.Median(numbers))
	fmt.Println("Variance:", formula.Variance(numbers))
	fmt.Println("Standard Deviation:", formula.StandardDeviation(numbers))
}
