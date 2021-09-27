package main

import "github.com/monzilnepali/mood-cli/cmd"

type SoundPreset struct {
	Name        string `json:"name"`
	VolumeLevel int    `json:"volume_level"` //percentage level
}

type Preset struct {
	Name   string        `json:"name"`
	Sounds []SoundPreset `json:"sounds"`
}

func main() {
	cmd.Execute()
}
