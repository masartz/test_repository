package main

import (
	"log"
	"os"
)

func main() {
	outputFile, err := os.Create("docs/index.html")
	if err != nil {
		log.Fatalf("Failed to create output file: %s", err)
	}
	defer outputFile.Close()
}
