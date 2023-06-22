package main

import (
	"image/jpeg"
	"log"
	"os"

	"github.com/kolesa-team/go-webp/encoder"
	"github.com/kolesa-team/go-webp/webp"
)

func convertJpegToWebp(path string) error {
	file, err := os.Open(path)
	if err != nil {
		log.Println(err)
		return err
	}
	image, err := jpeg.Decode(file)
	if err != nil {
		log.Println("here")
		return err
	}
	log.Println(path)
	output, err := os.Create("./" + path + ".webp")
	if err != nil {
		log.Println("f")
		return err
	}
	defer output.Close()

	options, err := encoder.NewLossyEncoderOptions(encoder.PresetPhoto, 100)
	if err != nil {
		return err
	}
	if err := webp.Encode(output, image, options); err != nil {
		return err
	}
	return nil
}
