package cmd

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
)

func Menu() {
	menuItem := ""
	prompt := &survey.Select{
		Message: "Select:",
		Options: []string{"Listen", "Available presets"},
	}
	survey.AskOne(prompt, &menuItem)

	switch menuItem {
	case "Listen":
		PromptSoundSelect()
	case "Available presets":
		fmt.Println("soon...")
	}

}
