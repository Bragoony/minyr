package main

import (
	"fmt"
	//"io"
	//"log"
	//"os"
	//"strings"
	//"bufio"

	"github.com/Bragoony/minyr/yr"
)


func main() {
	filePath := "kjevik-temp-celsius-20220318-20230318.csv"

	avg, err := yr.CalculateAverageFourthElement(filePath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Average of fourth elements: %f\n", avg)
}