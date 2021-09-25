package main

import (
	"fmt"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/effects"
	"github.com/faiface/beep/speaker"
	"github.com/monzilnepali/mood-cli/utils"
)

func main() {
	streamer, format := utils.GetAudioStreamer("./resources/data_resources_sounds_birds.ogg")
	streamer1, _ := utils.GetAudioStreamer("./resources/data_resources_sounds_rain.ogg")

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	ctrl := &beep.Ctrl{Streamer: beep.Loop(-1, streamer), Paused: false}
	ctrl1 := &beep.Ctrl{Streamer: beep.Loop(-1, streamer1), Paused: false}

	volume := &effects.Volume{
		Streamer: ctrl1,
		Base:     1,
		Volume:   -5,
		Silent:   false,
	}

	speaker.Play(ctrl)
	speaker.Play(volume)

	for {
		fmt.Print("Press [ENTER] to pause/resume. ")
		fmt.Scanln()
		speaker.Lock()
		ctrl.Paused = !ctrl.Paused
		ctrl1.Paused = !ctrl1.Paused
		speaker.Unlock()
	}
}
