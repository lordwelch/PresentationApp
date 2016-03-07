// imgprovider.go
package main

import (
	"fmt"
	"image"
	"strconv"
)

/*var imgproviderstr = `import QtQuick 2.4

Image {
    source: "image://images/`*/

func imgProvider(id string, width, height int) image.Image {
	var i int
	fmt.Println("id: ", id)
	i, _ = strconv.Atoi(id)
	fmt.Println("haha: ", i)
	return slides[i].getImage(width, height)
}
