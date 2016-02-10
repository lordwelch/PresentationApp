// magickwand.go
package main

import (
	"fmt"
	"math"

	"github.com/gographics/imagick/imagick"
)

func resizeImage(mw *imagick.MagickWand, newWidth, newHeight int, keepSpecSize, center bool) (resmw *imagick.MagickWand) {
	var (
		width, height, origHeight, origWidth int
	)
	origHeight = int(mw.GetImageHeight())
	fmt.Println("hahahahahah :-P")
	origWidth = int(mw.GetImageWidth())

	if (origHeight != newHeight) || (origWidth != newWidth) {
		if (round((float64(origWidth) / float64(origHeight)) * float64(newHeight))) <= newWidth {
			width = round((float64(origWidth) / float64(origHeight)) * float64(newHeight))
			height = newHeight
		} else {
			height = round((float64(origHeight) / float64(origWidth)) * float64(newWidth))
			width = newWidth
		}
	} else {
		height = newHeight
		width = newWidth
	}

	resmw = imagick.NewMagickWand()
	if !keepSpecSize {
		resmw.NewImage(uint(width), uint(height), imagick.NewPixelWand())
		center = false
	} else {
		resmw.NewImage(uint(newWidth), uint(newHeight), imagick.NewPixelWand())
		fmt.Println(resmw.GetImageHeight(), resmw.GetImageWidth())
		if center {
			err = mw.ResizeImage(uint(width), uint(height), imagick.FILTER_LANCZOS, 1)
			if err != nil {
				panic(err)
			}
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
