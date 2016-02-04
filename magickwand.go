// magickwand.go
package main

import (
	"math"

	"github.com/gographics/imagick/imagick"
)

func resizeImage(mw *imagick.MagickWand, newWidth, newHeight int, keepExactSize, center bool) (resmw *imagick.MagickWand) {
	var (
		width, height int
	)
	origHeight := int(mw.GetImageHeight())
	origWidth := int(mw.GetImageWidth())

	if (origHeight != newHeight) || (origWidth != newWidth) {
		if (round((origHeight / origWidth) * newWidth)) <= newHeight {
			height = round((float64(origHeight) / float64(origWidth))) * newWidth
			width = newWidth
		} else {
			width = round((float64(origWidth) / float64(origHeight))) * newHeight
			height = newHeight
		}
	} else {
		height = newHeight
		width = newWidth
	}

	if !keepExactSize {
		resmw.SetSize(width, height)
		center = false
	} else {
		resmw.SetSize(newWidth, newHeight)
		if center {
			err = mw.ResizeImage(width, height, imagick.FILTER_LANCZOS, 1)
			if err != nil {
				panic(err)
			}
			resmw.CompositeImage(mw, imagick.COMPOSITE_OP_SRC_OVER, uint((width-newWidth)/2), uint((height-newHeight)/2))
		} else {
			resmw.CompositeImage(mw, imagick.COMPOSITE_OP_SRC_OVER, 0, 0)
		}
	}
	mw.Destroy()
	return
}

func round(a float64) int {
	if a < 0 {
		return int(math.Ceil(a - 0.5))
	}
	return int(math.Floor(a + 0.5))
}
