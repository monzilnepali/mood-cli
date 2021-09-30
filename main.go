package main

import (
	"embed"

	"github.com/monzilnepali/mood-cli/cmd"
	"github.com/monzilnepali/mood-cli/sound"
)

//TODO: Remove the audio resource bind from build and use http download

//go:embed resources/*
var resources embed.FS

func main() {
	cmd.Execute()
}

func init() {
	sound.Resources = resources
}
