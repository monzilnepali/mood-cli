package sound

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/faiface/beep"
	"github.com/faiface/beep/effects"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/vorbis"
	"github.com/kyokomi/emoji"
	"github.com/monzilnepali/mood-cli/constants"
	"github.com/monzilnepali/mood-cli/utils"
)

var streamer beep.StreamSeekCloser

// var format beep.Format

type SoundPreset struct {
	Name        string `json:"name"`
	VolumeLevel int    `json:"volume_level"` //percentage level
}

type Preset map[string][]SoundPreset

func getAudioStreamer(audioPath string, volume_level int) (volume *effects.Volume) {

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Open(dir + audioPath)
	if err != nil {
		log.Fatal(err)
	}

	fileExtension := filepath.Ext(audioPath)

	if !utils.Include(constants.SupportedAudioFormat, fileExtension) {
		error := fmt.Errorf("invalid audio format")
		log.Fatal(error)
	}

	switch fileExtension {
	case ".ogg":
		streamer, _, err = vorbis.Decode(f)
	case ".mp3":
		streamer, _, err = mp3.Decode(f)
	}

	if err != nil {
		log.Fatal(err)
	}

	//Loop the audio
	ctrl := &beep.Ctrl{Streamer: beep.Loop(-1, streamer), Paused: false}

	var silent bool = false
	if volume_level == 0 {
		silent = true
	}

	volume = &effects.Volume{
		Streamer: ctrl,
		Base:     2,
		Volume:   -float64(100-volume_level) / 100.0 * 5,
		Silent:   silent,
	}

	return volume
}

//initialize the speaker
func Play(soundList []SoundPreset, label string) {

	//getcompose sound
	composedSounds := make([]*effects.Volume, 0)
	for _, sound := range soundList {
		soundPath := constants.SoundData[sound.Name]
		volume_level := sound.VolumeLevel
		streamer := getAudioStreamer(soundPath, volume_level)
		composedSounds = append(composedSounds, streamer)
	}

	//default sampling rate = 44100
	//TODO: Make the speaker sample dynamic
	//https://github.com/faiface/beep/wiki/Hello,-Beep!#dealing-with-different-sample-rates
	err := speaker.Init(44100, 4410)
	if err != nil {
		fmt.Print("Unable to initiate speaker")
		log.Fatal(err)
	}

	for _, audiodata := range composedSounds {
		if audiodata != nil {
			speaker.Play(audiodata)
		}
	}

	emojiString := emoji.Sprint(":sound:", label)
	fmt.Println(emojiString + "\n")
	for {
		fmt.Print("Press [ENTER] to stop. ")
		fmt.Scanln()
		os.Exit(0)
	}
}
