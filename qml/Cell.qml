import QtQuick 2.4

Rectangle {
    objectName: "cellRect"
    property int index: 0
    height: 100
    border.width: 2
    border.color: "black"
    anchors.right: parent.right
    anchors.left: parent.left

    Text {
        id: cellText
        enabled: true
        objectName: "cellText"
        text: ""
        clip: true
        wrapMode: Text.WrapAtWordBoundaryOrAnywhere
        anchors.fill: parent
        anchors.right: parent.right
        anchors.rightMargin: 0
        anchors.left: parent.left
        anchors.leftMargin: 2

        MouseArea {
            id: cellMouse
            hoverEnabled: true
            enabled: true
            objectName: "cellMouse"
            anchors.fill: parent
            acceptedButtons: Qt.AllButtons

            onMouseXChanged: cellHover()
            onExited: focusChanged(focus)

            function cellHover() {
                if (containsMouse) {
                    parent.parent.border.color = "skyblue"
                    parent.parent.color = "darkblue"
                    parent.color = "white"
                } else if (focus) {
                    parent.color = "black"
                }
            }

            function notSelected() {

                parent.parent.border.color = "black"
                parent.parent.color = "white"
                parent.color = "black"
                cellHover()
            }

            function selected() {
                parent.parent.border.color = "blue"
                parent.color = "black"
                parent.parent.color = "gainsboro"
                cellHover()
            }
        }
    }
    Image {
    id: img
    antialiasing: true
    source: "image://images/"
    objectName: "cellImg"
    property int index: 0
    height: 100
    transformOrigin: Item.TopLeft
    fillMode: Image.PreserveAspectFit
    anchors.right: parent.right
    anchors.left: parent.left
    //cache: false
    MouseArea {
            id: imgMouse
            hoverEnabled: true
            enabled: true
            objectName: "cellMouse"
            anchors.fill: parent
            acceptedButtons: Qt.AllButtons
            }
}
}
