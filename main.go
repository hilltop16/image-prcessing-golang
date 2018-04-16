package main

import "log"

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
	Fun with bild.
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

	shrp := sharpen(sat)
	err = saveImage(shrp, ".", "sharpened.jpg")
	if err != nil {
		log.Fatal(err)
	}

	pri := primitivePicture(sat)
	err = saveImage(pri, ".", "primitive.jpg")
	if err != nil {
		log.Fatal(err)
	}
}