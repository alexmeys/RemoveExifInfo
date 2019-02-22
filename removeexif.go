//removeexif.go
// Alex Meys

package main

import (
    	"os"
	"io/ioutil"
	"log"
	"image"
	"image/jpeg"
)

func main() {

    // read current directory 
    bestanden, err := ioutil.ReadDir(".")
    if err != nil {
        log.Fatal("Failed, %s\n", err)
    }
	for _, file := range bestanden {	
		bestand, err := os.Open(file.Name())
		if err != nil {
			log.Fatal(err)
		}
		defer bestand.Close()

		img, _, err := image.Decode(bestand)
		if err != nil {
			log.Fatal("Failed, %s\n", err)
		}

		outfile, err := os.Create(file.Name())
		if err != nil {
			log.Fatal("Failed, %s\n", err)
		}
		defer outfile.Close()

		if err := jpeg.Encode(outfile, img, nil); err != nil {
			log.Fatal("Failed, %s\n", err)
		}
	}
}
