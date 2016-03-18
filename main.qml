import QtQuick 2.4
import QtQuick.Dialogs 1.2
import QtQuick.Controls 1.3
import QtQuick.Window 2.0
import "qml"
import QtQuick.Layouts 1.2

ApplicationWindow {
    id: applicationWindow1
    title: "Presentation App"
    visible: true
    objectName: "applicationWindow1"
    minimumWidth: 500
    minimumHeight: 500
    width: 1000
    height: 600
    property bool cls: false

    function openFileDialog() {
        imgpicker.open()
        return imgpicker.fileUrl
    }

    onClosing: if (!cls) {
                   close.accepted = false
               }

    FileDialog {
        id: imgpicker
        title: "Choose an image for this slide"
        folder: shortcuts.pictures
    }

    AboutDialog {
        id: aboutDialog
    }

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


        Menu {
            objectName: "mnuCtx"
            title: "new image..."
            MenuItem {
                objectName: "mnuImgPick"
                text: "new Image..."
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
            border.width: 4
            anchors.left: parent.left
            anchors.leftMargin: 0
            anchors.bottom: parent.bottom
            anchors.bottomMargin: 0
            anchors.top: parent.top
            anchors.topMargin: 0
            property int count: 0

            ScrollView {
                id: scview
                anchors.fill: parent
                anchors.margins: 4
                horizontalScrollBarPolicy: Qt.ScrollBarAlwaysOff
                verticalScrollBarPolicy: Qt.ScrollBarAlwaysOn

                SplitView {
                    id: gridData
                    objectName: "gridData"
                    width: scview.width - 1
                    height: gridRect.count * 101

                    Rectangle {
                        id: col1
                        objectName: "col1"
                        width: gridData.width / 2
                        color: "#00000000"
                        transformOrigin: Item.TopLeft
                        border.width: 0
    Rectangle {
        id: textEdit
        property int cell
        x: 232
        y: 622
        objectName: "textEdit"
        width: 200
        height: 200
        color: "#ffffff"
        visible: false
        property bool txt: true
        Keys.onPressed: {
            if ((event.key == Qt.Key_Return)
                    && (event.modifiers & Qt.ControlModifier)) {
                txt = true

                x = -100
                y = -100
                visible = false
                focus = true
                enabled = false
                opacity = 0
                textEdit1.focus = false

                event.accepted = true
            }

            if (event.key == Qt.Key_Escape) {
                txt = false
                x = -100
                y = -100
                visible = false
                focus = true
                enabled = false
                opacity = 0
                textEdit1.focus = false

                event.accepted = true
            }
        }

        TextArea {
            id: textEdit1
            objectName: "textEdit1"
            anchors.fill: parent
            clip: true
            textFormat: Text.AutoText
            visible: true
            font.pixelSize: 12
            z: 99
        }
    }

                        Column {
                            id: data1
                            objectName: "data1"
                            spacing: 1
                            anchors.fill: parent
                        }
                    }

                    Rectangle {
                        id: col2
                        objectName: "col2"
                        color: "#00000000"
                        border.width: 0

                        Column {
                            id: data2
                            spacing: 1
                            objectName: "data2"
                            anchors.fill: parent
                        }
                    }
                }
            }
        }

        Rectangle {
            id: mainView
            border.width: 0
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

            Button {
                id: button1
                objectName: "btnAdd"
                x: 8
                y: 8
                text: qsTr("Button")
            }

            Button {
                id: button2
                x: 8
                y: 39
                text: qsTr("Button")
                objectName: "btnRem"
            }

            Button {
                id: button3
                x: 8
                y: 70
                text: qsTr("Button")
                objectName: "btnMem"
            }
        }
    }

}
