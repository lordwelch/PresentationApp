// PresentationApp project qml.go
package main

import (
	"image"
)

//imgProvider() for preview images in QML
func imgProvider(id string, width, height int) image.Image {
	var img1 image.Image
	if imgready && (len(id) > 0) {
		//fmt.Println("source (provider): ", id)
		// i1 := strings.Index(id, `;`)
		// i, _ := strconv.Atoi(id[:i1])
		// img1 = slides[i].getImage(width, height)
	} else {
		img1 = image.NewRGBA(image.Rect(0, 0, 340, 480))
	}
	return img1

}
