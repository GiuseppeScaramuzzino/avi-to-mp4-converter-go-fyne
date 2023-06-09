package main

import (
	"log"
	"os/exec"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {

	app := app.New()
	window := app.NewWindow("Avi2Mp4 Converter")
	box := container.NewVBox()

	var convert *widget.Button
	var upload *widget.Button
	var file string

	label := widget.NewLabel("Import avi from system")

	upload = widget.NewButtonWithIcon("Upload", theme.UploadIcon(), func() {
		window.SetFullScreen(true)
		fd := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err == nil && reader != nil {
				file = reader.URI().Path()
				label.SetText("Selected file: " + file)
				convert.Show()
				upload.Hide()
			}
			window.SetFullScreen(false)
			window.Resize(fyne.NewSize(600, 0))
		}, window)

		fd.Resize(fyne.NewSize(900, 900))
		fd.SetFilter(storage.NewExtensionFileFilter([]string{".avi"})) // filters to image files
		fd.Show()
	})

	convert = widget.NewButtonWithIcon("Convert", theme.MediaVideoIcon(), func() {
		progress := widget.NewProgressBarInfinite()
		convert.Hide()
		box.Add(progress)
		window.Content().Refresh()
		cmd := exec.Command("ffmpeg", "-i", file, file+".mp4")

		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		} else {
			progress.Hide()
			label.SetText(file + ".mp4")
		}
	})
	convert.Hide()

	box.Add(convert)
	box.Add(upload)
	window.SetContent(container.NewVBox(label, box))

	window.Resize(fyne.NewSize(600, 0))

	window.ShowAndRun()

}
