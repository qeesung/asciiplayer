
```                     
                         _    ____   ____ ___ ___ ____  _        _ __   _______ ____
                        / \  / ___| / ___|_ _|_ _|  _ \| |      / \\ \ / / ____|  _ \
                       / _ \ \___ \| |    | | | || |_) | |     / _ \\ V /|  _| | |_) |
                      / ___ \ ___) | |___ | | | ||  __/| |___ / ___ \| | | |___|  _ <
                     /_/   \_\____/ \____|___|___|_|   |_____/_/   \_\_| |_____|_| \_\
```

[![Build Status](https://travis-ci.org/qeesung/asciiplayer.svg?branch=master)](https://travis-ci.org/qeesung/asciiplayer)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/2dba80ebf4c04687a320858c385fe7f8)](https://app.codacy.com/app/qeesung/asciiplayer?utm_source=github.com&utm_medium=referral&utm_content=qeesung/asciiplayer&utm_campaign=Badge_Grade_Dashboard)
[![Coverage Status](https://coveralls.io/repos/github/qeesung/asciiplayer/badge.svg)](https://coveralls.io/github/qeesung/asciiplayer)
[![Go Report Card](https://goreportcard.com/badge/github.com/qeesung/asciiplayer)](https://goreportcard.com/report/github.com/qeesung/asciiplayer)
[![GoDoc](https://godoc.org/github.com/qeesung/asciiplayer?status.svg)](https://godoc.org/github.com/qeesung/asciiplayer)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

ASCIIPlayer is a library that convert gif/video into ASCII gif/video and provide command-line tools for easy use.

# Installation

```bash
go get -u github.com/qeesung/asciiplayer
```

# CLI usage

```
    _    ____   ____ ___ ___ ____  _        _ __   _______ ____
   / \  / ___| / ___|_ _|_ _|  _ \| |      / \\ \ / / ____|  _ \
  / _ \ \___ \| |    | | | || |_) | |     / _ \\ V /|  _| | |_) |
 / ___ \ ___) | |___ | | | ||  __/| |___ / ___ \| | | |___|  _ <
/_/   \_\____/ \____|___|___|_|   |_____/_/   \_\_| |_____|_| \_\
>>>Version  : 1.0.0
>>>Author   : qeesung
>>>HomePage : https://github.com/qeesung/asciiplayer

asciiplayer is a library that can convert gif and video to ASCII image
and provide the cli for easy use.

Usage:
  asciiplayer [command]

Available Commands:
  encode      Encode gif or video to ascii gif or video
  help        Help about any command
  play        Play the gif and video in ASCII mode
  server      Server command setup a server
  version     Show the version

Flags:
  -D, --debug   Switch log level to DEBUG mode
  -h, --help    help for asciiplayer

Use "asciiplayer [command] --help" for more information about a command.
```

## Command play

Play command only work in terminal, decoding the gif or video info multi frames and convert the frames to ASCII character matrix, finally, output the matrix to stdout at a certain frequency.

More detail please run `asciiplayer play --help`

### Play examples

Play the gif， and be able to match the screen size.
```bash
asciiplayer play demo.gif
```

Zoom to the original 1/10 and play it.
```bash
asciiplayer play demo.gif -r 0.1
```

Zoom to the fixed width and fixed height and play it
```bash
asciiplayer play demo.gif -w 100 -h 40
```

## Command encode

Encode command can convert gif or video to a ascii gif or video.

More detail please run `asciiplayer encode --help`

### Encode examples

Encode gif image to ascii gif image 
```bash
asciiplayer encode demo.gif -o output.gif
```

Encode gif image to ascii gif image with custom font size
```bash
asciiplayer encode demo.gif -o output.gif --font_size=5
```

Zoom to the original 1/10, then encode gif image to ascii gif image
```bash
asciiplayer encode demo.gif -o output.gif -r 0.1
```
