// imgprovider.go
package main

import (
	"image"
	"strconv"
)

<<<<<<< Updated upstream
/*var imgproviderstr = `import QtQuick 2.4

Image {
    source: "image://images/`*/
var imgready = false

func imgProvider(id string, width, height int) image.Image {
	if imgready {
		i, _ := strconv.Atoi(id)
		return slides[i].getImage(width, height)

	} else {
		var img1 image.Image = image.NewRGBA(image.Rect(0, 0, 340, 480))
		return img1
	}
=======
func imgProvider(id string, width, height int) image.Image {
	i, _ := strconv.Atoi(id)
	return slides[i].getImage(width, height)
>>>>>>> Stashed changes
}
