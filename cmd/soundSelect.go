package cmd

import (
	"log"

	"github.com/AlecAivazis/survey/v2"
	"github.com/monzilnepali/mood-cli/constants"
	"github.com/monzilnepali/mood-cli/sound"
)

func PromptSoundSelect() {
	selectSoundList := []string{}
	prompt := &survey.MultiSelect{
		Message: "Select:",
		Options: constants.SoundList,
	}
	err := survey.AskOne(prompt, &selectSoundList)
	if err != nil {
		log.Fatal(err)
	}
	streamingSoundList := make([]sound.SoundPreset, 0)

	for _, soundName := range selectSoundList {
		streamingSoundList = append(streamingSoundList, sound.SoundPreset{
			Name:        soundName,
			VolumeLevel: 10,
		})
	}

	composedSounds, err := sound.GetComposedSounds(streamingSoundList)

	if err != nil {
		log.Fatal(err)
	}

	sound.Play(composedSounds)
}
