package main

import (
	"bufio"
	"bytes"
	"image/jpeg"
	"io"
	"log"

	"github.com/kolesa-team/go-webp/encoder"
	"github.com/kolesa-team/go-webp/webp"
)

func convertJpegToWebp(file io.Reader) ([]byte, error) {

	image, err := jpeg.Decode(file)
	if err != nil {
		log.Println("here")
		return nil, err
	}
	var b bytes.Buffer

	output := bufio.NewWriter(&b)

	options, err := encoder.NewLossyEncoderOptions(encoder.PresetPhoto, 100)
	if err != nil {
		log.Println("hello")

		return nil, err
	}
	if err := webp.Encode(output, image, options); err != nil {
		log.Println("hello")
		return nil, err
	}
	output.Flush()
	return []byte(b.String()), nil
}
