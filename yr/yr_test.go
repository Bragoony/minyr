package yr

import (
	"testing"
)

func TestCelsiusToFahrenheitString(t *testing.T) {
    type test struct {
		input string
		want string
    }
    tests := []test{
	    {input: "6", want: "42.8"},
	    {input: "0", want: "32.0"},
    }

    for _, tc := range tests {
	    got, _ := CelsiusToFahrenheitString(tc.input)
	    if !(tc.want == got) {
	    	t.Errorf("expected %s, got: %s", tc.want, got)
	    }
    }
}

// Forutsetter at vi kjenner strukturen i filen og denne implementasjon 
// er kun for filer som inneholder linjer hvor det fjerde element
// på linjen er verdien for temperatrmaaling i grader celsius
func TestCelsiusToFahrenheitLine(t *testing.T) {
    type test struct {
		input string
		want string
    }
    tests := []test{
	    {input: "Kjevik;SN39040;18.03.2022 01:50;6", want: "Kjevik;SN39040;18.03.2022 01:50;42.8"},
	    {input: "Kjevik;SN39040;07.03.2023 18:20;0", want: "Kjevik;SN39040;07.03.2023 18:20;32.0"},
		{input: "Kjevik;SN39040;08.03.2023 02:20;-11", want: "Kjevik;SN39040;08.03.2023 02:20;12.2"},
    }

    for _, tc := range tests {
	    got, _ := CelsiusToFahrenheitLine(tc.input)
	    if !(tc.want == got) {
	     	t.Errorf("expected %s, got: %s", tc.want, got)
	    }
    }

	
}
func TestCountLines(t *testing.T) {
	expected := 16756
	filename := "../kjevik-temp-celsius-20220318-20230318.csv"
	lineCount, err := CountLines(filename)
	if err != nil {
		t.Errorf("error counting lines in %s: %v", filename, err)
	}
	if lineCount != expected {
		t.Errorf("unexpected number of lines in %s, expected %d but got %d", filename, expected, lineCount)
	}
}

func TestCalculateAverageFourthElement(t *testing.T) {
	// Test with a sample CSV file
	filePath := "../table.csv"
	expectedAverage := 4.92

	// Call the function to calculate the average
	average, err := CalculateAverageFourthElement(filePath)
	if err != nil {
		t.Errorf("error calculating average: %v", err)
	}

	// Check if the average is equal to the expected value
	if average != expectedAverage {
		t.Errorf("average %v does not match expected value %v", average, expectedAverage)
	}
}