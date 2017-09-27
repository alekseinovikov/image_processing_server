package resizing

import (
	"mime/multipart"
	"image"
	"image/jpeg"
	"github.com/freedom4live/resize"
	"bytes"
)

// Gets image from the multipart file
func getImage(file multipart.File) (image.Image, error) {
	defer file.Close()

	return jpeg.Decode(file)
}

// Resizing file
func Resize(file multipart.File, width int, height int) ([]byte, error) {
	img, err := getImage(file)
	if nil != err {
		return nil, err
	}

	newImage := resize.Resize(uint(width), uint(height), img, resize.Lanczos3)

	buf := new(bytes.Buffer)
	err = jpeg.Encode(buf, newImage, nil)
	return buf.Bytes(), nil
}

// Making thumbnail
func Thumbnail(file multipart.File, width int, height int) ([]byte, error) {
	img, err := getImage(file)
	if nil != err {
		return nil, err
	}

	newImage := resize.Thumbnail(uint(width), uint(height), img, resize.Lanczos3)

	buf := new(bytes.Buffer)
	err = jpeg.Encode(buf, newImage, nil)
	return buf.Bytes(), nil
}
