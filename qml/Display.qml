import QtQuick 2.4
import QtQuick.Controls 1.3

ApplicationWindow {
    flags: Qt.MaximumSize
    Component.onCompleted: visible = true

    Image {
        id: image1
        objectName: "displayImage"
        sourceSize.height: 768
        sourceSize.width: 1024
        antialiasing: true
        anchors.fill: parent
    }

}
