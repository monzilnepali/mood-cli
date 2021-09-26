package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/kyokomi/emoji"
	"github.com/monzilnepali/mood-cli/constants"
	"github.com/monzilnepali/mood-cli/sound"
	"github.com/monzilnepali/mood-cli/utils"
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
		volume_level := utils.PromptVolumeInput(soundName)
		streamingSoundList = append(streamingSoundList, sound.SoundPreset{
			Name:        soundName,
			VolumeLevel: volume_level,
		})
	}

	composedSounds, err := sound.GetComposedSounds(streamingSoundList)

	if err != nil {
		log.Fatal(err)
	}

	str := strings.Join(selectSoundList, ", ")
	emojiString := emoji.Sprint("\n :sound:", str)
	fmt.Println(emojiString)
	sound.Play(composedSounds)
}
