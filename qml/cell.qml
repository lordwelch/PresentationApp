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
        anchors.leftMargin: 2

        MouseArea {
            id: cellMouse
            hoverEnabled: true
            enabled: true
            objectName: "cellMouse"
            anchors.fill: parent
            onFocusChanged: if (focus) {
                                selected()
                            }

            onClicked: {
                focus = true
                selected()
                mouseXChanged(mouse)
            }

            onMouseXChanged: if (containsMouse) {
                                 parent.parent.border.color = "skyblue"
                                 parent.parent.color = "darkblue"
                                 parent.color = "white"
                             } else if (focus) {
                                 parent.color = "black"
                             }

            onExited: {
                parent.parent.border.color = "white"
                parent.parent.color = "white"
                parent.color = "black"
                if (focus) {
                    focusChanged(focus)
                }
            }

            function selected() {
                parent.parent.border.color = "blue"
                parent.parent.color = "gainsboro"
            }
        }
    }
}
