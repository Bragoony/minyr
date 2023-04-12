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
/*
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

	scanner := bufio.NewScanner(bufio.NewReader(src))
		numLines := 0
		for scanner.Scan(){
			numLines++	
		}

		scanner = bufio.NewScanner(bufio.NewReader(src))
		src.Seek(0, 0)
	writer := bufio.NewWriter(dest)

	for i := 1; scanner.Scan(); i++{
		line := scanner.Text()
		if i != 1 || i != numLines{
			newLine, err := yr.CelsiusToFahrenheitLine(line)
			_, err = writer.WriteString(newLine + "\n")
			if err != nil {
				log.Fatal(err)
			}
			continue
		}
	}
	err = writer.Flush()
	if err != nil {
		log.Fatal(err)
	}

}
*/
func main() {
	filePath := "table.csv"

	avg, err := yr.CalculateAverageFourthElement(filePath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Average of fourth elements: %.2f\n", avg)
}
