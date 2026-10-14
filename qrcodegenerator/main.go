package main

import (
	"fmt"

	"github.com/yeqown/go-qrcode/v2"
	"github.com/yeqown/go-qrcode/writer/standard"
)

func main() {
	createQRCode("https://github.com/manuelarte", "qr-github-manuelarte.jpeg")
	createQRCode("https://github.com/manuelarte/workshop-create-your-first-linter", "repository.jpeg")
}

func createQRCode(url string, filename string) {
	qrc, err := qrcode.New(url)
	if err != nil {
		fmt.Printf("could not generate QRCode: %v", err)
		return
	}

	w, err := standard.New(fmt.Sprintf("../slides/public/%s", filename))
	if err != nil {
		fmt.Printf("standard.New failed: %v", err)
		return
	}

	// save file
	if err = qrc.Save(w); err != nil {
		fmt.Printf("could not save image: %v", err)
	}
}
