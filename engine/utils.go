package engine

import (
	"image"
	"os"

	pix "github.com/faiface/pixel"
)

// Anything that gets updated every frame such as UI elements and Entities etc.
type Component interface {
	Update(*Game) // ? error
}

type State int

func loadPicture(path string) (pix.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pix.PictureDataFromImage(img), nil
}
