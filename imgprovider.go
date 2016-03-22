// imgprovider.go
package main

import (
	"image"
	"strconv"
	"strings"
)

var imgready = false

func imgProvider(id string, width, height int) image.Image {
	if imgready && (len(id) > 0) {
		//fmt.Println("source (provider): ", id)
		i1 := strings.Index(id, `;`)
		i, _ := strconv.Atoi(id[:i1])
		return slides[i].getImage(width, height)

	} else {
		var img1 image.Image = image.NewRGBA(image.Rect(0, 0, 340, 480))
		return img1
	}
}
