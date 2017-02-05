import QtQuick 2.4

Rectangle {
    id: rectangle1
    height: 100
    anchors.left: parent.left
    anchors.right: parent.right

    Rectangle {
        id: cellRect
        objectName: "cellRect"
        property int index: 0
        anchors.bottom: parent.bottom
        anchors.left: parent.left
        anchors.top: parent.top
        border.width: 2
        border.color: "black"
        width: rectangle1.width / 2

        Text {
            id: displayText
            enabled: true
            objectName: "cellText"
            text: cellText //"Hello\nMy\nName\nIs\n\"Timmy\""
            anchors.fill: parent
            anchors.leftMargin: 3
            clip: true
            wrapMode: Text.WrapAtWordBoundaryOrAnywhere

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
                    focus = true
                    parent.parent.border.color = "blue"
                    parent.parent.color = "gainsboro"
                    parent.color = "black"
                    cellHover()
                }
            }
        }
    }
    Rectangle {
        anchors.left: cellRect.right
        anchors.right: parent.right
        anchors.bottom: parent.bottom
        anchors.top: parent.top
        anchors.leftMargin: 0
        Image {
            id: img
            antialiasing: true
            source: imgSource
            objectName: "cellImg"
            property int index: 0
            anchors.fill: parent
            fillMode: Image.Stretch
            cache: false
            MouseArea {
                id: cellMse
                anchors.fill: parent
                hoverEnabled: true
                enabled: true
                objectName: "cellMouse"
                acceptedButtons: Qt.AllButtons
            }
        }
    }
}
