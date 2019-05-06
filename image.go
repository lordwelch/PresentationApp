// PresentationApp project imagick.go
package main

import (
	"image"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/disintegration/imaging"
	"github.com/flopp/go-findfont"
	"github.com/fogleman/gg"
)

/*resizeImage() src fullsize image
newWidth, newHeight = size to be resized to
keepSpecSize = return image with exactly the size specified or just the size of the resized image
center = center the image
*/
func resizeImage(src image.Image, newWidth, newHeight int, keepSpecSize, center bool) (dst image.Image) {
	imaging.Fit(src, newWidth, newHeight, imaging.Lanczos)

	if keepSpecSize {
		//blank image
		dst = image.NewNRGBA(image.Rect(0, 0, newWidth, newHeight))
		if center {
			dst = imaging.OverlayCenter(dst, src, 1)
		} else {
			dst = imaging.Overlay(dst, src, image.Pt(0, 0), 1)
		}
	}

	return dst

}

// adding text to image copied from example
func (cl *Cell) imgtext(width, height int) image.Image {
	ctx := gg.NewContextForImage(cl.image.resized)
	ctx.SetColor(cl.Font.color)

	if (cl.Font.name != "") || (cl.Font.name != "none") {
		data, err := ioutil.ReadFile(cl.Font.name)
		if err != nil {
			return image.Rectangle{}
		}
		if err := ctx.LoadFontData(data, cl.Font.size); err != nil {
			panic(err)
		}
	}
	ctx.DrawStringWrapped(cl.text, 0, 0, 0, 0, float64(width), 1, gg.AlignCenter)
	return ctx.Image()
}

func findFonts() {
	paths := findfont.List()
	for i, v := range paths {
		_, paths[i] = filepath.Split(v)
	}
	QML.FontList = strings.Join(paths, "\n")
}
