import QtQuick 2.5
import QtQuick.Controls 1.3

ApplicationWindow {
    title: "Presentation App"
    visible: true
    objectName: qsTr("")
    minimumWidth: 500
    minimumHeight: 500

    Grid {
        id: grid1
        x: 155
        y: 157
        width: 142
        height: 143
        clip: false
        columns: 2
        antialiasing: true
        z: 0
        rotation: 0
        scale: 1
        transformOrigin: Item.Center
        Column {
            TextArea {
                text: "test"
            }

        }

        Column {
        }
    }

}

