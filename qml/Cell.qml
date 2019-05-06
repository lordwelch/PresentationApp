import QtQuick 2.4
import QtQuick.Layouts 1.11

Rectangle {
    id: itm
    height: 100
    anchors.right: parent.right
    anchors.left: parent.left
            property alias text: cellText.text
    Rectangle {
        id: half1
        height: 100
        Layout.fillWidth: true
        Layout.minimumWidth: 100
        Rectangle {
            objectName: "cellRect"
            property int index: 0
            anchors.fill: parent
            border.width: 2
            border.color: "black"
            Text {
                id: cellText
                enabled: true
                objectName: "cellText"
                // text: "itm.model.text"
                renderType: Text.NativeRendering
                clip: true
                wrapMode: Text.WrapAtWordBoundaryOrAnywhere
                anchors.fill: parent
                anchors.right: parent.right
                anchors.rightMargin: 0
                anchors.left: parent.left
                anchors.leftMargin: 2
            }
        }
    }
}
