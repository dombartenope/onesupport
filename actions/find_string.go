package actions

import (
	"bufio"
	"encoding/csv"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/dombartenope/onesupport.git/userinput"
)

func FindSomeStringInFile() {
	files, err := os.ReadDir("./inputs")
	if err != nil {
		log.Fatalf("error reading inputs: %s", err)
	}

	//What output files do we need?
	var needCSV, needTXT bool
	for _, f := range files {
		if f.IsDir() {
			continue
		}
		switch ext := strings.ToLower(filepath.Ext(f.Name())); ext {
		case ".csv":
			needCSV = true
		case ".txt", ".log":
			needTXT = true
		}
		// if we already know we need both, break early
		if needCSV && needTXT {
			break
		}
	}

	searchTerm := userinput.SearchPrompt()

	// Prepare output writers based on what we need
	var (
		csvWriter *csv.Writer
		txtWriter *bufio.Writer
	)

	if needCSV {
		outCSV, err := os.Create("outputs/out.csv")
		if err != nil {
			log.Fatalf("error creating out.csv: %s", err)
		}
		defer outCSV.Close()
		csvWriter = csv.NewWriter(outCSV)
		defer csvWriter.Flush()
	}

	if needTXT {
		outTXT, err := os.Create("outputs/out.txt")
		if err != nil {
			log.Fatalf("error creating out.txt: %s", err)
		}
		defer outTXT.Close()
		txtWriter = bufio.NewWriter(outTXT)
		defer txtWriter.Flush()
	}

	// CSV header guard so we write headers only once
	headerWritten := false

	// Process each file
	for _, v := range files {
		if v.IsDir() {
			continue
		}
		name := v.Name()
		path := filepath.Join("inputs", name)
		file, err := os.Open(path)
		if err != nil {
			log.Fatalf("error opening %s: %s", path, err)
		}
		defer file.Close()

		ext := strings.ToLower(filepath.Ext(name))
		switch ext {
		case ".csv":
			if csvWriter == nil {
				// we decided we didn't need CSV output
				continue
			}
			reader := csv.NewReader(file)
			if !headerWritten {
				header, err := reader.Read()
				if err != nil {
					log.Fatalf("reading header from %s: %s", name, err)
				}
				if err := csvWriter.Write(header); err != nil {
					log.Fatalf("writing header: %s", err)
				}
				headerWritten = true
			}
			rows, err := reader.ReadAll()
			if err != nil {
				log.Fatalf("reading rows from %s: %s", name, err)
			}
			for _, row := range rows {
				for _, cell := range row {
					if strings.Contains(cell, searchTerm) {
						if err := csvWriter.Write(row); err != nil {
							log.Fatalf("writing row: %s", err)
						}
						break
					}
				}
			}

		case ".txt", ".log":
			if txtWriter == nil {
				// we decided we didn't need TXT output
				continue
			}
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				line := scanner.Text()
				if strings.Contains(line, searchTerm) {
					if _, err := txtWriter.WriteString(line + "\n"); err != nil {
						log.Fatalf("writing text: %s", err)
					}
				}
			}
			if err := scanner.Err(); err != nil {
				log.Fatalf("scanning %s: %s", name, err)
			}
		}
	}
}
