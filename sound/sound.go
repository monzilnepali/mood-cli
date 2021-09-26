package sound

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/faiface/beep"
	"github.com/faiface/beep/effects"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/vorbis"
	"github.com/monzilnepali/mood-cli/constants"
	"github.com/monzilnepali/mood-cli/utils"
)

var streamer beep.StreamSeekCloser
var format beep.Format

type SoundPreset struct {
	Name        string `json:"name"`
	VolumeLevel int    `json:"volume_level"`
}

type Preset struct {
	Name   string        `json:"name"`
	Sounds []SoundPreset `json:"sounds"`
}

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
		streamer, format, err = vorbis.Decode(f)
	case ".mp3":
		streamer, format, err = mp3.Decode(f)
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
		Base:     1,
		//TODO: Calculate the volume level by volumn percentage from preset
		Volume: float64(volume_level),
		Silent: silent,
	}

	return volume
}

//GetComposedSoundsFromPreset load user preset
func GetComposedSoundsFromPreset() (streamers []*effects.Volume, err error) {
	json_string := `
	{
		"name":"first",
		"sounds": [
			{
				"name": "rain",
				"volume_level": 10
			},
			{
				"name": "birds",
				"volume_level": 10
			},
			{
				"name": "waves",
				"volume_level": 10
			}
	]}`

	var preset Preset
	err = json.Unmarshal([]byte(json_string), &preset)
	if err != nil {
		error := fmt.Errorf("Unable to parse preset.json file")
		return nil, error
	}

	var data = make([]*effects.Volume, 2)

	for _, sound := range preset.Sounds {
		soundPath := constants.SoundData[sound.Name]
		volume_level := sound.VolumeLevel
		streamer := getAudioStreamer(soundPath, volume_level)
		data = append(data, streamer)
	}

	return data, nil
}

//GetComposedSounds
func GetComposedSounds(soundList []SoundPreset) (streamers []*effects.Volume, err error) {
	data := make([]*effects.Volume, len(soundList))

	for _, sound := range soundList {
		soundPath := constants.SoundData[sound.Name]
		volume_level := sound.VolumeLevel
		streamer := getAudioStreamer(soundPath, volume_level)
		data = append(data, streamer)
	}

	return data, nil
}

//initialize the speaker
func Play(soundList []*effects.Volume) {
	//default sampling rate = 44100
	//TODO: Make the speaker sample dynamic
	//https://github.com/faiface/beep/wiki/Hello,-Beep!#dealing-with-different-sample-rates
	err := speaker.Init(44100, 4410)
	if err != nil {
		fmt.Print("Unable to initiate speaker")
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
