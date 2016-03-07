import QtQuick 2.4

Image {
    id: img
    antialiasing: true
    source: "image://images/"
    objectName: "cellImg"
    property int index: 0
    height: 150
    transformOrigin: Item.TopLeft
    fillMode: Image.PreserveAspectFit
    anchors.right: parent.right
    anchors.left: parent.left
    Layout.minimumHeight: 100
    Layout.maximumHeight: 200
}
