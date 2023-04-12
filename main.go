package main

import (
	//"fmt"
	//"io"
	"log"
	"os"
	//"strings"
	"bufio"

	"github.com/Bragoony/minyr/yr"
)


func main() {
	src, err := os.Open("table.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer src.Close()
	dest, err := os.OpenFile("kjevik-temp-fahr-20220318-20230318.csv", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer dest.Close()

	lineNumber := 0

	scanner := bufio.NewScanner(bufio.NewReader(src))
	writer := bufio.NewWriter(dest)

	for scanner.Scan(){
		lineNumber++
		line := scanner.Text()
		if lineNumber == 1 {
			_, err = writer.WriteString(line + "\n")
			if err != nil {
				log.Fatal(err)
				}
			continue
			}
			if lineNumber == 27 {
				_, err = writer.WriteString(line)
				if err != nil {
					log.Fatal(err)
					}
				continue
				}
			newLine, err := yr.CelsiusToFahrenheitLine(line)
			_, err = writer.WriteString(newLine + "\n")
			if err != nil {
				log.Fatal(err)
				}
	}
	err = writer.Flush()
	if err != nil {
		log.Fatal(err)
	}

}


/*
func main() {
	filePath := "table.csv"

	avg, err := yr.CalculateAverageFourthElement(filePath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Average of fourth elements: %.2f\n", avg)
}
*/