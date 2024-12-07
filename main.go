package main

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/code128"
	"github.com/boombuler/barcode/datamatrix"
	"github.com/boombuler/barcode/pdf417"
	"github.com/boombuler/barcode/qr"
	"github.com/sqweek/dialog"
)

func generateBarcode(content string, codeType string) (image.Image, error) {
	var barcodeImg barcode.Barcode
	var err error

	switch codeType {
	case "QR Code":
		barcodeImg, err = qr.Encode(content, qr.M, qr.Auto)
	case "Code 128":
		barcodeImg, err = code128.Encode(content)
	case "Datamatrix":
		barcodeImg, err = datamatrix.Encode(content)
	case "PDF 417":
		barcodeImg, err = pdf417.Encode(content, 3)
	default:
		return nil, fmt.Errorf("unsupported code type selected")
	}

	if err != nil {
		return nil, err
	}

	barcodeImg, err = barcode.Scale(barcodeImg, 600, 600)
	if err != nil {
		return nil, err
	}

	return barcodeImg, nil
}

func imageToResource(img image.Image) fyne.Resource {
	buf := new(bytes.Buffer)
	if err := png.Encode(buf, img); err != nil {
		log.Println("Failed to encode image to resource:", err)
		return nil
	}
	return fyne.NewStaticResource("barcode.png", buf.Bytes())
}

func main() {

	application := app.New()
	window := application.NewWindow("Barcode Generator by LewdLillyVT")

	input := widget.NewEntry()
	input.SetPlaceHolder("Enter content for the barcode")

	codeTypes := []string{"QR Code", "Code 128", "Datamatrix", "PDF 417"}
	selectType := widget.NewSelect(codeTypes, func(selected string) {
		log.Println("Selected code type:", selected)
	})
	selectType.SetSelected("QR Code")

	errorLabel := widget.NewLabel("")
	errorLabel.Hide()

	var generatedBarcode image.Image
	generateBtn := widget.NewButton("Generate", func() {
		if input.Text == "" {
			errorLabel.SetText("Error: Barcode content cannot be empty!")
			errorLabel.Show()
			return
		}

		selectedType := selectType.Selected
		var err error
		generatedBarcode, err = generateBarcode(input.Text, selectedType)
		if err != nil {
			errorLabel.SetText("Error: " + err.Error())
			errorLabel.Show()
			return
		}

		errorLabel.Hide()
	})

	saveBtn := widget.NewButton("Save", func() {
		if generatedBarcode == nil {
			errorLabel.SetText("Error: No barcode generated yet!")
			errorLabel.Show()
			return
		}

		filePath, err := dialog.File().Filter("PNG files", "png").Save()
		if err != nil {
			errorLabel.SetText("Error: Save dialog failed.")
			errorLabel.Show()
			return
		}

		if filePath == "" {
			errorLabel.SetText("Error: No file path selected.")
			errorLabel.Show()
			return
		}

		file, err := os.Create(filePath)
		if err != nil {
			errorLabel.SetText("Error: Failed to create file.")
			errorLabel.Show()
			return
		}
		defer file.Close()

		if err := png.Encode(file, generatedBarcode); err != nil {
			errorLabel.SetText("Error: Failed to save image.")
			errorLabel.Show()
		} else {

			errorLabel.Hide()
		}
	})

	content := container.NewVBox(
		input,
		selectType,
		generateBtn,
		saveBtn,
		errorLabel,
	)

	window.SetContent(content)
	window.Resize(fyne.NewSize(400, 600))
	window.ShowAndRun()
}
