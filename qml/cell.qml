import QtQuick 2.0

Rectangle {
    enabled: true
    objectName: "cellRect"
    property int index: 0
    width: 100
    height: 100
    border.width: 2
    anchors.right: parent.right
    anchors.left: parent.left
    onFocusChanged: if (focus) {
                        border.color = "gainsboro"
                        color = "blue"
                    }

    Text {
        id: cellText
        enabled: true
        objectName: "cellText"
        text: "hello this is text\nhaha\nhdsjfklfhaskjd"
        clip: true
        wrapMode: Text.WrapAtWordBoundaryOrAnywhere
        anchors.fill: parent
        anchors.right: parent.right
        anchors.rightMargin: 0
        anchors.left: parent.left
        anchors.leftMargin: 0
        onFocusChanged: if (focus) {
                            parent.border.color = "gainsboro"
                            parent.color = "blue"
                        }

        MouseArea {
            id: cellMouse
            hoverEnabled: true
            enabled: true
            objectName: "cellMouse"
            anchors.fill: parent
            onFocusChanged: if (focus) {
                                parent.parent.border.color = "gainsboro"
                                parent.parent.color = "blue"
                            }
            onClicked: focus = true
            onEntered: if (containsMouse) {
                           parent.parent.border.color = "skyblue"
                           parent.parent.color = "darkblue"
                           parent.color = "white"
                       }
            onExited: {
                parent.parent.border.color = "white"
                parent.parent.color = "white"
                parent.color = "black"
            }
        }
    }
}
