
# :tv: ASCIIPlayer

[![Build Status](https://travis-ci.org/qeesung/asciiplayer.svg?branch=master)](https://travis-ci.org/qeesung/asciiplayer)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/2dba80ebf4c04687a320858c385fe7f8)](https://app.codacy.com/app/qeesung/asciiplayer?utm_source=github.com&utm_medium=referral&utm_content=qeesung/asciiplayer&utm_campaign=Badge_Grade_Dashboard)
[![Coverage Status](https://coveralls.io/repos/github/qeesung/asciiplayer/badge.svg)](https://coveralls.io/github/qeesung/asciiplayer)
[![Go Report Card](https://goreportcard.com/badge/github.com/qeesung/asciiplayer)](https://goreportcard.com/report/github.com/qeesung/asciiplayer)
[![GoDoc](https://godoc.org/github.com/qeesung/asciiplayer?status.svg)](https://godoc.org/github.com/qeesung/asciiplayer)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

```                  
                       _    ____   ____ ___ ___ ____  _        _ __   _______ ____
                      / \  / ___| / ___|_ _|_ _|  _ \| |      / \\ \ / / ____|  _ \
                     / _ \ \___ \| |    | | | || |_) | |     / _ \\ V /|  _| | |_) |
                    / ___ \ ___) | |___ | | | ||  __/| |___ / ___ \| | | |___|  _ <
                   /_/   \_\____/ \____|___|___|_|   |_____/_/   \_\_| |_____|_| \_\
```

ASCIIPlayer is a library and tool that can play picture(png,jpeg,gif) and video(mp4,avi) in ASCII mode and can convert the picture and video into ASCII picture. 

![fire](https://github.com/qeesung/asciiplayer/blob/master/docs/images/ascii_cat.gif?raw=true#center)  


## Table of contents

- [Features](#features)
- [How it works?](#how-it-works)
- [CLI usage](#cli-usage)
    - [Command Play](#command-play)
    - [Command Encode](#command-encode)
    - [Command Server](#command-server)
- [Library usage](#library-usage)
- [Examples](#examples)
- [Todo](#todo)
- [License](#license)

## Features

- Support playing a PNG, JPEG, GIF type picture in the terminal in ASCII mode, as well as MP4, AVI and many other types of video(playing video still working in progress). More details can be accessed at [play](#command-play)
- Support encoding a common PNG, JPEG, GIF type picture into a ASCII picture, as well as MP4, AVI and many other types of video
(encoding video still working in progress). More details can be accessed at [encode](#command-encode)
- Support for building a HTTP server, and you can share your ASCII picture and video to others(servering vedio still working in progress). More details can be accessed at [server](#command-server)

## How it works?

ASCIIPlayer is base on [Image2ASCII](https://github.com/qeesung/image2ascii)(which is a library that converts image into ASCII image). 

- Firstly, we need to decode the input media (picture, gif, vedio) into multi frames, different media use different decoders. For example, if you input a gif file, we will use GifDecoder to decode the gif into multi frames.
- Secondly, build a Image2ASCII converter that can convert the pictures into ASCII images (it would be a long string , or a ascii pixel matrix).
- Finally, Display the ASCII images in different ways, there are three ways to display the ASCII images:
  - (**Encode Mode**) Display in a file, we need to render the ascii pixel matrix into a file, just draw it pixel by pixel. 
  - (**Play Mode**) Display in the terminal, simplely output the ascii image(string type) to the stdout at a certain frequency.
  - (**Server Mode**) Just like display in the terminal, but it displays at remote clinet, we need to setup a http server, and then flush the ascii image to remote client, only when the response received by the client is exported to the terminal can it work properly.

```
                 +---------------+                                                  +---------+
                 |               |                                                  |         |
          +------> Gif Decoder   |                                              +---> Encoder +---> file
          |      |               |                                              |   |         |
          |      +---------------+                                              |   +---------+
          |      +---------------+                +-------------+               |   +---------+
          |      |               |                |             |               |   |         |
Input File+------> Image Decoder +---> Frames +-->+ Image2ASCII +->ASCII Frames-+----> Player  +---> stdout
          |      |               |                |             |               |   |         |
          |      +---------------+                +-------------+               |   +---------+
          |      +---------------+                                              |   +---------+
          |      |               |                                              |   |         |
          +------> Video Decoder |                                              +---> Server  +---> socket
                 |               |                                                  |         |
                 +---------------+                                                  +---------+
```

## Installation

```bash
go get -u github.com/qeesung/asciiplayer
```

## CLI usage

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

### Command play

Play command only work in terminal, decoding the gif or video info multi frames and convert the frames to ASCII character matrix, finally, output the matrix to stdout at a certain frequency.

More detail please run `asciiplayer play --help`

![play tutorial gif](https://github.com/qeesung/asciiplayer/blob/master/docs/images/play_tutorial.gif?raw=true)

#### Play examples

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

Play the png image
```
asciiplayer play demo.png
```

### Command encode

Encode command can convert gif or video to a ascii gif or video.

More detail please run `asciiplayer encode --help`

![encode tutorial gif](https://github.com/qeesung/asciiplayer/blob/master/docs/images/encode_tutorial.gif?raw=true)

ascii_eye.gif
![eye gif](https://github.com/qeesung/asciiplayer/blob/master/docs/images/ascii_eye.gif?raw=true)

#### Encode examples

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

Encode jpeg image to ascii png image
```bash
asciiplayer encode demo.jpeg -o output.png
```

### Command server

Setup a http server, and share your ascii image with others. Setup a http server, then access through curl command.

Setup server
```bash
$ asciiplayer server demo.gif
# Server available on : http://0.0.0.0:8080
```

Access from remote
```bash
$ curl http://hostname:8080
# play ascii image here
```

More detail please run `asciiplayer server --help`

![server tutorial gif](https://github.com/qeesung/asciiplayer/blob/master/docs/images/server_tutorial.gif?raw=true)


#### Server examples

Setup a http server with default port and host
```bash
asciiplayer server demo.gif
```

Setup a http server with the custom port
```bash
asciiplayer server demo.gif --port 8888
```

Setup a http server and share the ascii png image
```bash
asciiplayer server demo.png
```

## Library usage

Please access the godoc https://godoc.org/github.com/qeesung/asciiplayer

## Examples

Encoding gif sample

| Raw Image | ASCII Image|
|:--:|:--:|
|![](https://github.com/qeesung/asciiplayer/blob/master/docs/images/fire.gif?raw=true)|![](https://github.com/qeesung/asciiplayer/blob/master/docs/images/ascii_fire.gif?raw=true)|
|![](https://github.com/qeesung/asciiplayer/blob/master/docs/images/eye.gif?raw=true)|![](https://github.com/qeesung/asciiplayer/blob/master/docs/images/ascii_eye.gif?raw=true)|
|![](https://github.com/qeesung/asciiplayer/blob/master/docs/images/pounch.gif?raw=true)|![](https://github.com/qeesung/asciiplayer/blob/master/docs/images/ascii_pounch.gif?raw=true)|


Encoding jpeg sample

| Raw Image | ASCII Image|
|:--:|:--:|
|![](https://github.com/qeesung/asciiplayer/blob/master/docs/images/zolo.jpg?raw=true)|![](https://github.com/qeesung/asciiplayer/blob/master/docs/images/ascii_zolo.png?raw=true)|
|![](https://github.com/qeesung/asciiplayer/blob/master/docs/images/gg.jpg?raw=true)|![](https://github.com/qeesung/asciiplayer/blob/master/docs/images/ascii_gg.jpg?raw=true)|

## Todo

- Support playing, encoding, servering video
- Accelerating the encoding process

## License

This project is under the MIT License. See the [LICENSE](https://github.com/qeesung/asciiplayer/blob/master/LICENSE) file for the full license text.
