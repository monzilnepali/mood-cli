package cmd

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/monzilnepali/mood-cli/file"
	"github.com/monzilnepali/mood-cli/sound"
)

func PromptPresetSelect() {
	availablePresets, isExist := file.ReadPreset()

	if !isExist {
		fmt.Println("No preset")
		return
	}

	var availablePresetNameList []string
	for presetName := range availablePresets {
		availablePresetNameList = append(availablePresetNameList, presetName)
	}

	if !isExist {
		fmt.Println("No preset")
		return
	}

	selectedPreset := ""
	prompt := &survey.Select{
		Message: "Select preset",
		Options: availablePresetNameList,
	}

	survey.AskOne(prompt, &selectedPreset)
	sound.Play(availablePresets[selectedPreset], selectedPreset)
}
