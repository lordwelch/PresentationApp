// magickwand.go
package main

import (
	"math"

	"gopkg.in/gographics/imagick.v2/imagick"
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

func round(a float64) int {
	if a < 0 {
		return int(math.Ceil(a - 0.5))
	}
	return int(math.Floor(a + 0.5))
}
