package file

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/monzilnepali/mood-cli/sound"
)

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func GetConfigDirPath() (dirPath string) {
	homeDirPath, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	dirPath = path.Join(homeDirPath, "mood-cli")
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err := os.Mkdir(dirPath, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}
	return dirPath
}

func WriteToPreset(jsonDate sound.Preset) {
	configDirPath := GetConfigDirPath()
	presetFilePath := path.Join(configDirPath, "preset.json")

	file, _ := json.MarshalIndent(jsonDate, "", " ")
	err := ioutil.WriteFile(presetFilePath, file, 0644)
	if err != nil {
		fmt.Println("Error on saving preset")
	}

}

func ReadPreset() (preset sound.Preset, presetExist bool) {
	configDirPath := GetConfigDirPath()
	presetFilePath := path.Join(configDirPath, "preset.json")
	isPresetExist := fileExists(presetFilePath)
	if !isPresetExist {
		return nil, false
	}

	var savedPresetList sound.Preset
	data, err := ioutil.ReadFile(presetFilePath)

	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(data, &savedPresetList)
	if err != nil {
		log.Fatal(err)
	}

	return savedPresetList, true
}

func UpdatePreset(preset_name string, sounds []sound.SoundPreset) {
	updatedPreset := sound.Preset{}

	previousPreset, isExist := ReadPreset()
	if isExist {
		updatedPreset = previousPreset
	}
	updatedPreset[preset_name] = sounds
	WriteToPreset(updatedPreset)
}
