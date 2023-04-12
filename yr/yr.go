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

func studentString(string) (string, error) {
	var err error

	tekst := ("Endringen er gjort av Brage")
	return tekst, err
}

func studentLine(line string) (string, error) {

	dividedString := strings.Split(line, ";")
	var err error

	if (len(dividedString) == 4) {
		dividedString[3], err = studentString(dividedString[3])
		if err != nil {
			return "", err
		}
	} else {
		return "", errors.New("linje har ikke forventet format")
	}
	return strings.Join(dividedString, ";"), nil
}

func countLines(filename string) (int, error) {
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

func editLastLine(filename string) error {
	// Open the file for reading and writing
	file, err := os.OpenFile(filename, os.O_RDWR, 0644)
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	// Seek to the end of the file
	_, err = file.Seek(0, io.SeekEnd)
	if err != nil {
		return fmt.Errorf("error seeking to end of file: %v", err)
	}

	// Read the last line
	reader := bufio.NewReader(file)
	lastLine, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		return fmt.Errorf("error reading last line: %v", err)
	}

	// Truncate the last line by deleting the last 3 characters
	newLastLine := lastLine[:len(lastLine)-3]

	// Write the modified last line
	newLastLine += "Made by Brage\n"
	_, err = file.WriteAt([]byte(newLastLine), int64(len(newLastLine)-len(lastLine)))
	if err != nil {
		return fmt.Errorf("error writing modified last line: %v", err)
	}

	return nil
}
