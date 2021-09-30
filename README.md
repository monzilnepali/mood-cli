# Mood-cli :notes:
![Go Report Card](https://goreportcard.com/badge/github.com/monzilnepali/mood-cli)
![Linux](https://svgshare.com/i/Zhy.svg)
![release](https://img.shields.io/github/v/release/monzilnepali/mood-cli?include_prereleases)


Listen to chillaxing nature sound from your cli :smile:


## Built With
- [beep](https://github.com/faiface/beep)
- [cobra](https://github.com/spf13/cobra)
- [survey](https://github.com/AlecAivazis/survey)

## Getting started

### Prerequisites
- go 1.17

### Installation
1. Clone the repo
    ```bash
     git clone git@github.com:monzilnepali/mood-cli.git
    ```
2. Install depdendencies
    ```bash
     go install
    ```
3. Run
    ```bash
    go run main.go
    ```



For playback, Beep uses Oto under the hood.
- Linux
ALSA is required. On Ubuntu or Debian, run this command:
```bash
 apt install libasound2-dev
```
