package main

import (
	"log"
	"image"
	"os"
	"errors"
	"image/jpeg"
	"path"
	"github.com/artyom/smartcrop"
)

type SubImager interface {
	SubImage(r image.Rectangle) image.Image
}

func main() {
	img, err := openImage("original.jpg")
	if err != nil {
		log.Fatal(err)
	}

	img, err = crop(img, 1000, 1000)
	if err != nil {
		log.Fatal(err)
	}

	err = saveImage(img, ".", "cropped.jpg")
	if err != nil {
		log.Fatal(err)
	}

	img, err = openImage("cropped.jpg")
	if err != nil {
		log.Fatal(err)
	}

	sat := saturate(img)
	err = saveImage(sat, ".", "saturated.jpg")
	if err != nil {
		log.Fatal(err)
	}

	mult := multiply(img)
	err = saveImage(mult, ".", "multiplied.jpg")
	if err != nil {
		log.Fatal(err)
	}

}

func openImage(path string) (image.Image, error) {
	imgFile, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "Cannot open "+path)
	}

	img, err := jpeg.Decode(imgFile)
	if err != nil {
		return nil, errors.Wrap(err, "Decoding the image failed.")
	}

	return img, nil
}

func saveImage(img image.Image, pname, fname string) error {
	fpath := path.Join(pname, fname)

	f, err := os.Create(fpath)
	if err != nil {
		return errors.Wrap(err, "Cannot create file: "+fpath)
	}
	err = jpeg.Encode(f, img, &jpeg.Options{Quality: 85})
	if err != nil {
		return errors.Wrap(err, "Failed to encode the image as JPEG")
	}
	return nil
}

func crop(img image.Image, width, height int) (image.Image, error) {
	rect, err := smartcrop.Crop(img, width, height)
	if err != nil {
		return nil, errors.Wrap(err, "Smartcrop failed")
	}
	si, ok := (img).(SubImager)
	if !ok {
		return nil, errors.New("crop(): img does not support SubImage()")
	}
	subImg := si.SubImage(rect)
	return subImg, nil
}