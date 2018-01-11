package main

import (
	"bytes"
	"flag"
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"reflect"

	"github.com/nfnt/resize"
)

var ASCIISTR = "MND8OZ$7I?+=~:,.."

func Init() (image.Image, int) {
	width := flag.Int("w", 80, "use -w <width>")
	fpath := flag.String("p", "test.jpg", "Use -p <filesource>")
	flag.Parse()
	f, err := os.Open(*fpath)
	if err != nil {
		log.Fatal(err)
	}

	img, _, err := image.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	f.Close()
	return img, *width

}

func ScaleImage(img image.Image, w int) (image.Image, int, int) {
	sz := img.Bounds()
	h := (sz.Max.Y * w * 10) / (sz.Max.X * 16)
	img = resize.Resize(uint(w), uint(h), img, resize.Lanczos3)
	return img, w, h
}

func Convert2Ascii(img image.Image, w, h int) []byte {
	table := []byte(ASCIISTR)
	buf := new(bytes.Buffer)

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			g := color.GrayModel.Convert(img.At(j, i))
			y := reflect.ValueOf(g).FieldByName("Y").Uint()
			pos := int(y * 16 / 255)
			_ = buf.WriteByte(table[pos])
		}
		_ = buf.WriteByte('\n')
	}
	return buf.Bytes()
}

/*
func main() {
	p := Convert2Ascii(ScaleImage(Init()))
	fmt.Print(string(p))
}
*/
