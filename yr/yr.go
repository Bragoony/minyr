package yr

import (
	"fmt"
	"strconv"
	"strings"
	"errors"
	"github.com/Bragoony/funtemps/conv"
	"bufio"
	"os"
	"io"
	"log"
	"math"
)

func CelsiusToFahrenheitString(celsius string) (string, error) {
	var fahrFloat float64
	var err error
	if celsiusFloat, err := strconv.ParseFloat(celsius, 64); err == nil {
		fahrFloat = conv.CelsiusToFarhenheit(celsiusFloat)
	}
	fahrString := fmt.Sprintf("%.1f", fahrFloat)
	return fahrString, err
}

// Forutsetter at vi kjenner strukturen i filen og denne implementasjon 
// er kun for filer som inneholder linjer hvor det fjerde element
// på linjen er verdien for temperaturaaling i grader celsius
func CelsiusToFahrenheitLine(line string) (string, error) {

        dividedString := strings.Split(line, ";")
	var err error
	
	if (len(dividedString) == 4) {
		dividedString[3], err = CelsiusToFahrenheitString(dividedString[3])
		if err != nil {
			return "", err
		}
	} else {
		return "", errors.New("linje har ikke forventet format")
	}
	return strings.Join(dividedString, ";"), nil


	//return "Kjevik;SN39040;18.03.2022 01:50;42.8", err
}

func CountLines(filename string) (int, error) {
	// Open the file for reading
	file, err := os.Open(filename)
	if err != nil {
		return 0, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	// Create a new scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Loop through the file and count the lines
	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("error scanning file: %v", err)
	}

	return lineCount, nil
}

func EditLastLine(filename string) error {
    file, err := os.OpenFile(filename, os.O_RDWR, 0644)
    if err != nil {
        return err
    }
    defer file.Close()

    buf := make([]byte, 3)
    _, err = file.Seek(-3, io.SeekEnd)
    if err != nil {
        return err
    }
    _, err = file.Read(buf)
    if err != nil {
        return err
    }
    if string(buf) != ";;;" {
        return errors.New("last line doesn't end with ';;;'")
    }

    _, err = file.Seek(-2, io.SeekEnd)
    if err != nil {
        return err
    }
    _, err = file.Write([]byte("endringen er gjort av Brage Kjemperud"))
    if err != nil {
        return err
    }

    return nil
}

func CalculateAverageFourthElement(filePath string) (float64, error) {
	// Open the CSV file
	file, err := os.Open(filePath)
	if err != nil {
		return 0, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Initialize variables to keep track of the sum and count of fourth elements
	sum := 0.0
	count := 0

	// Loop through each line in the file
	lineNumber := 0
	for scanner.Scan() {
		lineNumber++
		if lineNumber == 1 || lineNumber == 16756 {
			continue
		}

		// Split the line into fields
		line := scanner.Text()
		fields := strings.Split(line, ";")
		if len(fields) < 4 {
			return 0, fmt.Errorf("line %d has less than 4 fields", lineNumber)
		}

		// Convert the fourth field to a float and add it to the sum
		num, err := strconv.ParseFloat(fields[3], 64)
		if err != nil {
			return 0, fmt.Errorf("error converting field %d in line %d to float: %v", 3, lineNumber, err)
		}
		sum += num
		count++
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("error reading file: %v", err)
	}

	// Calculate the average of the fourth elements
	if count == 0 {
		return 0, fmt.Errorf("no valid lines found")
	}
	average := sum / float64(count)
	average = math.Round(average*100) / 100

	return average, nil
}

func ConvertCelsiusFileToFahrenheitFile() {
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
			if lineNumber == 16756 {
				_, err = writer.WriteString(line)
				if err != nil {
					log.Fatal(err)
					}
				continue
				}
			newLine, err := CelsiusToFahrenheitLine(line)
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

func ReadLastLine(filePath string) (string, error) {
	// Open the CSV file
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Initialize a variable to store the last line
	var lastLine string

	// Loop through each line in the file and keep updating the last line variable
	for scanner.Scan() {
		lastLine = scanner.Text()
	}

	// Check if there was an error during scanning
	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error reading file: %v", err)
	}

	// Check if the last line contains the expected string
	expectedString := "Data er gyldig per 20.03.2023 (CC BY 4.0), Meteorologisk institutt (MET);endringen er gjort av Brage Kjemperud"
	if strings.Contains(lastLine, expectedString) {
		return lastLine, nil
	}

	// Return an error if the last line doesn't contain the expected string
	return "", fmt.Errorf("last line does not contain the expected string")
}
