package main

import (
	"github.com/dombartenope/onesupport.git/actions"
	"github.com/dombartenope/onesupport.git/userinput"
)

func main() {

	userChoice := userinput.InitialPrompt()
	switch userChoice {
	case 1:
		actions.FilterLogsByOneSignal()
	case 2:
		actions.FindSomeStringInFile()
	}

}
