package main

import (
	"fmt"
	"log"
	"os"

	"github.com/faiface/beep/speaker"
	"github.com/monzilnepali/mood-cli/sound"
)

// 44100 4410

func main() {

	//initialize the speaker
	//TODO: Make the speaker sample dynamic
	// https://github.com/faiface/beep/wiki/Hello,-Beep!#dealing-with-different-sample-rates
	speaker.Init(44100, 4410)

	soundList, err := sound.LoadPreset()

	if err != nil {
		log.Fatal(err)
	}

	for _, audiodata := range soundList {
		if audiodata != nil {
			speaker.Play(audiodata)
		}

	}

	for {
		fmt.Print("Press [ENTER] to stop. ")
		fmt.Scanln()
		os.Exit(0)

	}

}
