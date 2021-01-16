package utils

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"

	"github.com/oxodao/vibes/services"
)

// SetPicture convert the given picture in all formats and save it
func SetPicture(prv *services.Provider, file multipart.File) (string, error) {
	rndName := prv.GenerateUID(20)

	f, err := os.OpenFile("./pictures/"+rndName+".jpg", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return "", err
	}

	_, err = io.Copy(f, file)
	if err != nil {
		return "", err
	}

	// Creating all the size of pictures
	// It is probably better to store it in all format than resizing it each call
	// The app calls for picture A LOT
	saveImageAtSize("portrait_medium", rndName)
	saveImageAtSize("medium", rndName)
	saveImageAtSize("xxlarge", rndName)
	saveImageAtSize("small", rndName)

	return rndName, nil
}

func saveImageAtSize(size string, filename string) {
	// size
	/**
		portrait_medium (Used in: Main frame background)
		medium (Used in: small picture in the settings)
	**/
	copy("./pictures/"+filename+".jpg", "./pictures/"+filename+"_"+size+".jpg")

}

// Temporary function that copy an image, this will be removed
// later because picture are supposed to be resized
func copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}
