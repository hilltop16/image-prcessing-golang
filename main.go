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
}