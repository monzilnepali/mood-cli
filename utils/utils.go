package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/vorbis"
	"github.com/monzilnepali/mood-cli/constants"
)

var streamer beep.StreamSeekCloser
var format beep.Format

func GetAudioStreamer(audioPath string) (streamer beep.StreamSeekCloser, format beep.Format) {

	f, err := os.Open(audioPath)
	if err != nil {
		log.Fatal(err)
	}

	fileExtension := filepath.Ext(audioPath)

	if !Include(constants.SupportedAudioFormat, fileExtension) {
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

	return streamer, format
}

func loadPreset() {

}

// Include check whether item exist in slice
func Include(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}
