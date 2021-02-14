package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"image/png"
	"log"
	"os"
)

func main() {
	sf, err := os.Open("testdata/sample.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer sf.Close()

	img, _, err := image.Decode(sf)
	if err != nil {
		log.Fatal(err)
	}

	df, err := os.Create("testdata/sample.png")
	defer func() {
		if err := df.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	png.Encode(df, img)
	fmt.Println("Success!")
}
