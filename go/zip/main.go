package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"

	"archive/zip"
)

type FileContent struct {
	filename string
	content  Content
}

type Content struct {
	item1, item2 string
}

func main() {
	r, err := zip.OpenReader("data.zip")

	if err != nil {
		log.Fatalf("could not open: %v", err)
	}

	defer r.Close()

	for _, f := range r.File {
		file, err := f.Open()

		if err != nil {
			log.Fatalf("could not open the file %s: %v", f.Name, err)
		}

		reader := csv.NewReader(file)
		_, err = reader.Read()

		if err == io.EOF {
			continue
		}

		if err != nil {
			log.Fatal(err)
		}

		for {
			row, err := reader.Read()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatal(err)
			}

			fmt.Print(row)
		}
	}
}
