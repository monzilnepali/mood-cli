package utils

import (
	"fmt"
	"log"
	"strconv"

	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/terminal"
)

// Include check whether item exist in slice
func Include(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func PromptVolumeInput(label string) (number int) {
	value := ""
	for {
		volumePrompt := &survey.Input{
			Message: fmt.Sprintf("Enter volume level for %s [0-100]:", label),
		}
		promptErr := survey.AskOne(volumePrompt, &value)
		// https://github.com/AlecAivazis/survey#why-isnt-ctrl-c-working
		if promptErr == terminal.InterruptErr {
			log.Fatal("exit")
		}
		volume_level, err := strconv.ParseInt(value, 10, 0)

		if err != nil {
			fmt.Print("Invalid volume level")
			continue
		}

		if volume_level <= 100 && volume_level >= 0 {
			return int(volume_level)
		}
		fmt.Println("Valid volume level range [0-100]")
	}

}
