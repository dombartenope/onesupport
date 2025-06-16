package actions

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func FilterLogsByOneSignal() {
	files, err := os.ReadDir("./inputs")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	out, err := os.Create("outputs/out.txt")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	writer := bufio.NewWriter(out)
	defer writer.Flush()

	for _, v := range files {
		fmt.Printf("%s\n", v.Name())
		if filepath.Ext(v.Name()) == ".txt" {
			crashName := fmt.Sprintf("\n\n%s\n\n", v.Name())
			_, err = writer.WriteString(crashName)
			if err != nil {
				log.Fatalf("Error writing new file: %s", err)
			}

			fileName := fmt.Sprintf("./inputs/%s", v.Name())
			file, err := os.Open(fileName)
			if err != nil {
				log.Fatalf("error opening file: %s", err)
			}

			/* For VERBOSE & DEBUG logs */
			// ||
			// 	strings.Contains(scanner.Text(), "VERBOSE") ||
			// 	strings.Contains(scanner.Text(), "DEBUG")
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				if strings.Contains(scanner.Text(), "OneSignal") {
					_, err := writer.WriteString(fmt.Sprintf("%s\n", scanner.Text()))
					if err != nil {
						log.Fatalf("Error writing new file: %s", err)
					}

				}
			}
			file.Close()
		}
	}
}
