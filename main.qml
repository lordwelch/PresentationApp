import QtQuick 2.5
import QtQuick.Controls 1.3
import QtQuick.Dialogs 1.1
import QtQuick.Layouts 1.0
import "qml"

ApplicationWindow {
    id: applicationWindow1
    title: "Presentation App"
    visible: true
    objectName: "applicationWindow1"
    minimumWidth: 500
    minimumHeight: 500
    width: 1000
    height: 600
    AboutDialog { id: aboutDialog }

    Action {
        id: aboutAction
        text: "About"
        onTriggered: aboutDialog.open()
    }

    menuBar: MenuBar {
        Menu {
            title: "&File"
            MenuItem {

            }
            MenuItem {
                text: "Close"
                shortcut: StandardKey.Quit
            }
        }
        Menu {
            title: "&Help"
            MenuItem {
                action: aboutAction
            }
        }
    }

    SplitView {
        id: mainSlider
        objectName: "mainSlider"
        anchors.right: parent.right
        anchors.bottom: parent.bottom
        anchors.top: parent.top
        anchors.left: parent.left
        anchors.rightMargin: 0
        anchors.bottomMargin: 0
        anchors.leftMargin: 0
        anchors.topMargin: 0
        orientation: Qt.Horizontal
        onResizingChanged: col1.width = gridData.width / 2

        Rectangle {
            id: gridRect
            objectName: "gridRect"
            width: 300
            color: "#00000000"
            radius: 1
            border.color: "#000000"
            border.width: 4
            anchors.left: parent.left
            anchors.leftMargin: 0
            anchors.bottom: parent.bottom
            anchors.bottomMargin: 0
            anchors.top: parent.top
            anchors.topMargin: 0

            SplitView {
                id: gridData
                objectName: "gridData"
                anchors.rightMargin: 4
                anchors.leftMargin: 4
                anchors.bottomMargin: 4
                anchors.topMargin: 4
                anchors.fill: parent

                Rectangle {
                    id: col1
                    objectName: "col1"
                    width: gridData.width / 2
                    color: "#e41616"
                    transformOrigin: Item.TopLeft
                    border.width: 0

                    Column {
                        id: data1
                        objectName: "data1"
                        spacing: 1
                        anchors.right: parent.right
                        anchors.bottom: parent.bottom
                        anchors.top: parent.top
                        anchors.left: parent.left
                        //onAdded: children.width = data1.width
                    }
                }

                Rectangle {
                    id: col2
                    objectName: "col2"
                    color: "#4f90e2"
                    border.width: 0

                    Column {
                        id: data2
                        objectName: "data2"
                        anchors.fill: parent
                    }
                }
            }
        }

        Rectangle {
            id: mainView
            objectName: "mainView"
            anchors.right: parent.right
            anchors.rightMargin: 0
            anchors.leftMargin: 0
            anchors.left: gridRect.right
            anchors.bottom: parent.bottom
            anchors.top: parent.top
            z: 1
            clip: false
            visible: true
        }
    }

    Rectangle {
        id: textEdit
        property int cell
        x: 448
        y: 151
        objectName: "textEdit"
        width: 200
        height: 200
        color: "#ffffff"
        visible: false
        Keys.onEscapePressed: {
            x = -100
            y = -100
            visible = false
            focus = true
            enabled = false
            opacity = 0
            textEdit1.focus = false
        }

        TextArea {
            id: textEdit1
            objectName: "textEdit1"
            text: "Text Edit"
            anchors.fill: parent
            clip: true
            textFormat: Text.AutoText
            visible: true
            font.pixelSize: 12
            z: 99
        }
    }
}
