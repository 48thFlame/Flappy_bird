package engine

import (
	"image"
	"os"

	pix "github.com/faiface/pixel"
)

func Touching(a, b HitBoxable) bool {
	return a.ToRect().Intersects(b.ToRect())
}

func TouchingEdge(a HitBoxable, windowWidth, windowHeight float64) bool {
	rect := a.ToRect()

	return rect.Min.X <= 0 || rect.Max.X >= windowWidth || rect.Min.Y <= 0 || rect.Max.Y >= windowHeight
}

func LoadPicture(path string) (pix.Picture, error) {
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
