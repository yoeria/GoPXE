package conf

import (
	"flag"
	"testing"
)

func TestSetup(t *testing.T) {
	Setup()
	expectedResult := "9090"
	actualResult := flag.Lookup("port").Value.String()

	if actualResult != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}
