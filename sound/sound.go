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

func GetAudioStreamer(audioPath string, volume_level int) (volume *effects.Volume) {

	f, err := os.Open(audioPath)
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
		Volume:   float64(volume_level),
		Silent:   silent,
	}

	return volume
}

//LoadPreset load user preset all sound and play
func LoadPreset() (streamers []*effects.Volume, err error) {
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
		streamer := GetAudioStreamer(soundPath, volume_level)
		data = append(data, streamer)
	}

	return data, nil
}
