import QtQuick 2.0

Rectangle {
    enabled: true
    objectName: "cellRect"
    property int index: 0
    width: 100
    height: 100
    color: "#00000000"
    anchors.right: parent.right
    anchors.left: parent.left

    Text {
        id: cellText
        enabled: true
        objectName: "cellText"
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
            enabled: true
            objectName: "cellMouse"
            anchors.fill: parent
        }
    }
}