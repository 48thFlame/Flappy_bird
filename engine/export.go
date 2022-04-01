package engine

import (
	"image"
	"os"

	pix "github.com/faiface/pixel"
)

type Dim struct {
	Width, Height float64
}

// Anything that gets updated every frame such as UI elements and Entities etc.
type Component interface {
	Update(*Game) // ? error
}

type HitBoxAble interface {
	ToRect() pix.Rect
}

func Touching(a, b HitBoxAble) bool {
	return a.ToRect().Intersects(b.ToRect())
}

func TouchingEdge(a HitBoxAble, windowWidth, windowHeight float64) bool {
	rect := a.ToRect()

	return rect.Min.X <= 0 || rect.Max.X >= windowWidth || rect.Min.Y <= 0 || rect.Max.Y >= windowHeight
}

type State int

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
