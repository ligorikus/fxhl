package texture

import (
	"image"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/ftrvxmtrx/tga"
)

func NewImage(filepath string) image.Image {
	imgFile, err := os.Open(getProjectRoot() + "/" + filepath)
	if err != nil {
		log.Fatalf("texture %q not found on disk: %v\n", filepath, err)
	}

	img, err := tga.Decode(imgFile)
	if err != nil {
		panic(err)
	}

	return img
}

func NewRgba(img image.Image) *image.RGBA {
	rgba := image.NewRGBA(img.Bounds())
	if rgba.Stride != rgba.Rect.Size().X*4 {
		panic("unsupported stride")
	}
	return rgba
}

func getProjectRoot() string {
	_, b, _, _ := runtime.Caller(0)
	projectRoot := filepath.Join(filepath.Dir(b), "../../")
	return projectRoot
}
