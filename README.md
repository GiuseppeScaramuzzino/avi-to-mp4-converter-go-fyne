# Avi2Mp4 converter (Go&Fyne)

This repository contains the source code for the tool `avi-to-mp4-converter-go-fyne`, a simple yet powerful AVI to MP4 video converter built using Go and Fyne.io, a cross-platform GUI toolkit. This application provides a user-friendly desktop interface for converting AVI video files to MP4 format.

## Pre-requisites

The application relies on FFmpeg, a suite of libraries and programs for handling multimedia data, for video conversion. Therefore, you need to have FFmpeg installed on your system to use this tool.

### Installing FFmpeg

#### On Ubuntu

```bash
sudo apt update
sudo apt install ffmpeg
```
#### On Windows

1. Download the FFmpeg build from the [official site](https://www.ffmpeg.org/download.html).
2. Unzip the downloaded file to a location of your choice.
3. Add the `/bin` directory from the unzipped files to your system path.

#### On MacOs

```bash
brew install ffmpeg
```

## Installation

To install `avi-to-mp4-converter-go-fyne`, please follow the steps below:

1. Clone the repository:

```bash
git clone https://github.com/gscaramuzzino/avi-to-mp4-converter-go-fyne.git
```

2. Navigate into the cloned repository:

```bash
cd avi-to-mp4-converter-go-fyne
```

3. Build the application:

```bash
go build
```

## Usage

After building the application, you can run the binary:

```bash
./avi-to-mp4-converter-go-fyne
```

Then, use the GUI to select the AVI file you wish to convert and click the convert button.

## Contributing

We welcome contributions from the community. If you wish to contribute, please follow the standard GitHub pull request process:

1. Fork the repository.
2. Make your changes on a new branch.
3. Submit a pull request with your changes.

We will review all pull requests and provide feedback as necessary.
