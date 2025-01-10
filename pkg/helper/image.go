package helper

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"image"
	"os"

	"github.com/anthonynsimon/bild/blur"
	"github.com/anthonynsimon/bild/imgio"
	"github.com/anthonynsimon/bild/transform"
)

func OpenImage(filepath string) (image.Image, error) {
	img, err := imgio.Open(filepath)
	if err != nil {
		return nil, err
	}
	return img, nil
}

func Thumbnail(img image.Image, size int) *image.RGBA {
	return transform.Resize(img, size, size, transform.Linear)
}

func Placeholder(img image.Image) *image.RGBA {
	return blur.Gaussian(Thumbnail(img, 360), 50)
}

func ImageAsBytes(img image.Image, quality int) []byte {
	encoder := imgio.JPEGEncoder(10)
	var b bytes.Buffer
	foo := bufio.NewWriter(&b)
	encoder(foo, img)
	return b.Bytes()
}

func Base64Image(img image.Image, quality int) string {
	return base64.StdEncoding.EncodeToString(ImageAsBytes(img, quality))

}

func SaveAsImage(img image.Image, filepath string, quality int) error {
	return os.WriteFile(filepath, ImageAsBytes(img, quality), 0655)
}
