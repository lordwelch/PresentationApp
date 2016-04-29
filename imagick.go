// PresentationApp project imagick.go
package main

import (
	"strings"
	"bytes"
	"image"
	"log"
	"math"
	"os/exec"

	"gopkg.in/gographics/imagick.v2/imagick"
)

var (
	fontlst []string
)

/*resizeImage() mw fullsize image
newwidth, newheight = size to be resized to
keepSpecSize = return image with exactly the size specified or just the size of the resized image
center = senter the image
*/
func resizeImage(mw *imagick.MagickWand, newWidth, newHeight int, keepSpecSize, center bool) (resmw *imagick.MagickWand) {
	var (
		width, height, origHeight, origWidth int
	)
	origHeight = int(mw.GetImageHeight())
	origWidth = int(mw.GetImageWidth())

	//check if requested size is the same as current size
	if (origHeight != newHeight) || (origWidth != newWidth) {
		// width / height * newheight = newwidth
		if (round((float64(origWidth) / float64(origHeight)) * float64(newHeight))) <= newWidth {
			width = round((float64(origWidth) / float64(origHeight)) * float64(newHeight))
			height = newHeight
		} else {
			// height / width * newwidth = newheight
			height = round((float64(origHeight) / float64(origWidth)) * float64(newWidth))
			width = newWidth
		}
	} else {
		height = newHeight
		width = newWidth
	}

	//new magickwand for resized image
	resmw = imagick.NewMagickWand()

	if !keepSpecSize {
		resmw.NewImage(uint(width), uint(height), imagick.NewPixelWand())
		center = false
	} else {
		//blank image
		resmw.NewImage(uint(newWidth), uint(newHeight), imagick.NewPixelWand())
		if center {
			err = mw.ResizeImage(uint(width), uint(height), imagick.FILTER_LANCZOS, 1)
			if err != nil {
				panic(err)
			}
			//centers image
			resmw.CompositeImage(mw, imagick.COMPOSITE_OP_SRC_OVER, round(float64(newWidth-width)/float64(2)), round(float64(newHeight-height)/float64(2)))
		} else {
			resmw.CompositeImage(mw, imagick.COMPOSITE_OP_SRC_OVER, 0, 0)
		}
	}
	mw.Destroy()
	return resmw

}

//getImage() from imagick to image.RGBA
func (cl *cell) getImage(width, height int) (img *image.RGBA) {
	mw := cl.img.GetImage()
	if (width == 0) || (height == 0) {
		width = int(mw.GetImageWidth())
		height = int(mw.GetImageHeight())
	}

	mw = resizeImage(mw, width, height, true, true)
	img = image.NewRGBA(image.Rect(0, 0, int(width), int(height)))
	if img.Stride != img.Rect.Size().X*4 {
		panic("unsupported stride")
	}

	Tpix, _ := mw.ExportImagePixels(0, 0, uint(width), uint(height), "RGBA", imagick.PIXEL_CHAR)
	img.Pix = Tpix.([]uint8)
	mw.Destroy()
	return
}

// adding text to image copied from example
func (cl *cell) imgtext() {
	mw := imagick.NewMagickWand()
	defer mw.Destroy()
	dw := imagick.NewDrawingWand()
	defer dw.Destroy()
	pw := imagick.NewPixelWand()
	defer pw.Destroy()
	pw.SetColor("none")

	// Create a new transparent image
	mw.NewImage(0, 0, pw)

	// Set up a 72 point white font
	pw.SetColor("white")
	dw.SetFillColor(pw)
	dw.SetFont("Verdana-Bold-Italic")
	dw.SetFontSize(72)

	// Add a black outline to the text
	pw.SetColor("black")
	dw.SetStrokeColor(pw)

	// Turn antialias on - not sure this makes a difference
	dw.SetTextAntialias(true)

	// Now draw the text
	dw.Annotation(25, 65, cl.text)

	// Draw the image on to the mw
	mw.DrawImage(dw)

	// Trim the image down to include only the text
	mw.TrimImage(0)

	// equivalent to the command line +repage
	mw.ResetImagePage("")

	// Make a copy of the text image
	cw := mw.Clone()

	// Set the background colour to blue for the shadow
	pw.SetColor("blue")
	mw.SetImageBackgroundColor(pw)

	// Opacity is a real number indicating (apparently) percentage
	mw.ShadowImage(70, 4, 5, 5)

	// Composite the text on top of the shadow
	mw.CompositeImage(cw, imagick.COMPOSITE_OP_OVER, 5, 5)
	cw.Destroy()

	// and write the result
	mw.WriteImage("text_shadow.png")
}

func findfonts() {
	cmd := exec.Command("grep", "-ivE", `\-Oblique$|-Bold$|-Italic$|-Light$`)
	cmd.Stdin = strings.NewReader(strings.Join(imagick.QueryFonts("*"), "\n"))
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Print(err)
	}
	fontlst = strings.Split(out.String(), "\n")
}

func round(a float64) int {
	if a < 0 {
		return int(math.Ceil(a - 0.5))
	}
	return int(math.Floor(a + 0.5))
}
