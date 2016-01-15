import QtQuick 2.0

Rectangle {
    property int index: 0

    Text {
        id: cellText
        height: 75
        text: "hello this is text\nhaha\nhdsjfklfhaskjd"
        textFormat: Text.AutoText
        clip: true
        font.bold: false
        anchors.fill: parent
        wrapMode: Text.WrapAtWordBoundaryOrAnywhere
        horizontalAlignment: Text.AlignLeft
        verticalAlignment: Text.AlignTop
        anchors.right: parent.right
        anchors.rightMargin: 0
        anchors.left: parent.left
        anchors.leftMargin: 0
        font.pixelSize: 12
        MouseArea {
            id: cellMouse
            anchors.fill: parent
        }
    }

