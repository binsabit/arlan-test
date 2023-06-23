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
	//decoding file to jpeg
	image, err := jpeg.Decode(file)
	if err != nil {
		log.Println("here")
		return nil, err
	}

	//creating io.Reader interface to store file in bytes in memory
	var b bytes.Buffer

	output := bufio.NewWriter(&b)

	options, err := encoder.NewLossyEncoderOptions(encoder.PresetPhoto, 100)
	if err != nil {
		return nil, err
	}
	//encode image to webp
	if err := webp.Encode(output, image, options); err != nil {
		return nil, err
	}
	output.Flush()
	return []byte(b.String()), nil
}
