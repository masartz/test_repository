package main

import (
	"log"
	"os"
	"path/filepath"
)

func main() {

	outputFile, err := os.Create("docs/index.html")
	if err != nil {
		log.Fatalf("Failed to create output file: %s", err)
	}
	defer outputFile.Close()

	// parse dir and create index.html
	paths := dirwalk("misc_dir")
	d := []byte{}
	for _, path := range paths {
		d = append(d, []byte(path)...)
	}

	_, err = outputFile.Write(d)
	if err != nil {
		log.Fatal(err)
	}
}

func dirwalk(dir string) []string {
	files, err := os.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			paths = append(paths, dirwalk(filepath.Join(dir, file.Name()))...)
			continue
		}
		paths = append(paths, filepath.Join(dir, file.Name()))
	}

	return paths
}
