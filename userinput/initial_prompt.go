package userinput

import (
	"log"

	"github.com/manifoldco/promptui"
)

func InitialPrompt() int {

	choices := map[string]int{
		"1 - Filter logs by OneSignal":       1,
		"2 - Search for something in a file": 2,
	}

	prompt := promptui.Select{
		Label: "Select an option",
		Items: []string{
			"1 - Filter logs by OneSignal",
			"2 - Search for something in a file",
		},
	}

	_, res, err := prompt.Run()
	if err != nil {
		log.Fatalf("Error for var \"res\": %s", err)
	}
	return choices[res]
}
