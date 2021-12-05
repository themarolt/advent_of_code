package libs

import (
	"log"
	"os"
	"strings"
)

func ReadTxtFileLines(filename string) []string {
	binaryContent, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error while reading file '%v'. Error: %v", filename, err)
	}

	return strings.Split(string(binaryContent), "\n")
}
