package userinput

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func SearchPrompt() string {
	fmt.Println("Enter the search term you want to find in the file:")
	reader := bufio.NewReader(os.Stdin)
	searchPrompt, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Error for var \"searchPrompt\": %s", err)
	}
	searchTerm := strings.TrimSpace(searchPrompt)

	return searchTerm
}
