package actions

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/dombartenope/onesupport.git/userinput"
)

func FindSomeStringInFile() {
	files, err := os.ReadDir("./inputs")
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	searchTerm := userinput.SearchPrompt()

	for _, v := range files {
		fileName := fmt.Sprintf("inputs/%s", v.Name())
		file, err := os.Open(fileName)
		if err != nil {
			log.Fatalf("Error for var \"file\": %s", err)
		}

		if strings.Contains(file.Name(), ".txt") || strings.Contains(file.Name(), ".log") {
			//Only create txt output if the file is a .txt or .log
			outTXT, err := os.Create("outputs/out.txt")
			if err != nil {
				log.Fatalf("error creating outTXT: %s", err)
			}

			//bufio for txt reading
			writer := bufio.NewWriter(outTXT)
			defer writer.Flush()
			//READ FILE
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				//CHECK FOR SEARCH TERM
				if strings.Contains(scanner.Text(), searchTerm) {
					//WRITE TO OUTPUT FILE
					_, err := writer.WriteString(fmt.Sprintf("%s\n", scanner.Text()))
					if err != nil {
						log.Fatalf("Error for var \"txt writer\": %s", err)
					}
				}
			}
		} else if strings.Contains(file.Name(), ".csv") {
			//Only create csv output if the file is a .csv
			outCSV, err := os.Create("outputs/out.csv")
			if err != nil {
				log.Fatalf("error creating outCSV: %s", err)
			}

			//csv reader for individual header and rows
			reader := csv.NewReader(file)
			writer := csv.NewWriter(outCSV)
			defer writer.Flush()

			//READ HEADERS
			header, err := reader.Read()
			if err != nil {
				log.Fatalf("Error for var \"header\": %s", err)
			}
			err = writer.Write(header)
			if err != nil {
				log.Fatalf("Error for var \"err\": %s", err)
			}

			//READ ROWS
			rows, err := reader.ReadAll()
			if err != nil {
				log.Fatalf("Error for var \"rows\": %s", err)
			}
			for _, row := range rows {
				for _, r := range row {
					//CHECK FOR SEARCH TERM
					if strings.Contains(r, searchTerm) {
						//WRITE TO OUTPUT FILE
						err := writer.Write(row)
						if err != nil {
							log.Fatalf("Error for var \"csv writer\": %s", err)
						}

					}
				}
			}
		}

	}
}
