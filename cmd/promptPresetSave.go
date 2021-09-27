package cmd

import (
	"github.com/AlecAivazis/survey/v2"
)

func PromptPresetSave() (preset_name string, isSave bool) {
	isSaveToPreset := false
	saveToPresetPrompt := &survey.Confirm{
		Message: "Save as preset ?",
	}
	survey.AskOne(saveToPresetPrompt, &isSaveToPreset)
	if isSaveToPreset {
		presetName := ""
		presetNamePrompt := &survey.Input{
			Message: "Preset name",
		}
		survey.AskOne(presetNamePrompt, &presetName, survey.WithValidator(survey.Required))

		return presetName, true
	}
	return "", false
}
