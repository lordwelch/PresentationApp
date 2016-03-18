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
            acceptedButtons: Qt.AllButtons

            onMouseXChanged: if (containsMouse) {
                                 parent.parent.border.color = "skyblue"
                                 parent.parent.color = "darkblue"
                                 parent.color = "white"
                             } else if (focus) {
                                 parent.color = "black"
                             }

            onExited: focusChanged(focus)

            function notSelected() {

                parent.parent.border.color = "black"
                parent.parent.color = "white"
                parent.color = "black"
            }

            function selected() {
                parent.parent.border.color = "blue"
                parent.parent.color = "gainsboro"
            }
        }
    }
}
