import QtQuick 2.4
import QtQuick.Controls 1.6 as Quick
import QtQuick.Controls 2.4
import QtQuick.Dialogs 1.3
import QtQuick.Window 2.11
import QtQuick.Layouts 1.11

ApplicationWindow {
    id: applicationWindow1
    title: "Presentation App"
    visible: true
    objectName: "applicationWindow1"
    minimumWidth: 500
    minimumHeight: 500

    FileDialog {
        id: imgpicker
        objectName: "imgpicker"
        title: "Choose an image for this slide"
    }

    Quick.SplitView {
        id: spview
        anchors.fill: parent
        Rectangle {
            id: preview
            objectName: "col1"
            border.width: 0
            Layout.minimumWidth: 150
            Layout.fillWidth: true

            Flickable {
                id: scview
                objectName: "scview"
                anchors.fill: parent
                boundsBehavior: Flickable.OvershootBounds
                flickableDirection: Flickable.VerticalFlick
                pixelAligned: true
                //verticalScrollBarPolicy: Qt.ScrollBarAlwaysOn
                //horizontalScrollBarPolicy: Qt.ScrollBarAlwaysOff
                //highlightOnFocus: false
                //frameVisible: true
                contentHeight: contentItem.childrenRect.height

                Quick.SplitView {
                    anchors.fill: parent
                }

                Rectangle {
                    id: textEdit
                    objectName: "textEdit"
                    visible: false
                    property bool keepText: true
                    Keys.onPressed: {
                        if ((event.key == Qt.Key_Return)
                                && (event.modifiers & Qt.ControlModifier)) {
                            keepText = true
                            textEdit1.focus = false
                            event.accepted = true
                        }

                        if (event.key == Qt.Key_Escape) {
                            keepText = false
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
                        hoverEnabled: false
                    }
                }
            }
        }

        Rectangle {
            id: mainView
            objectName: "mainView"
            Layout.minimumWidth: 100
            Layout.fillWidth: false

            Button {
                id: button1
                objectName: "btnAdd"
                x: 8
                y: 8
                text: qsTr("Add")
                onClicked: sv.addLst("fail")
            }

            Button {
                id: button2
                x: 8
                y: 49
                text: qsTr("Remove")
                objectName: "btnRem"
            }

            Button {
                id: button3
                x: 8
                y: 90
                text: qsTr("Button ")
                objectName: "btnMem"
            }
        }
    }
    /*
    menuBar: MenuBar {
        Menu {
            title: "&File"
            MenuItem {
                text: "Close"
                shortcut: StandardKey.Quit
            }
        }
        Menu {
            title: "&Edit"
            MenuItem {
                text: "quick edit"
                objectName: "mnuEdit"
            }
        }

        Menu {
            title: "Window"

            MenuItem {
                text: "Display"
                objectName: "mnuDisplay"
            }
        }

        Menu {

            MenuItem {
                text: "&help"
            }
        }
    }
*/
    Menu {
        objectName: "mnuCtx"
        title: "new image..."
        MenuItem {
            objectName: "mnuImgPick"
            text: "new Image..."
            onTriggered: imgpicker.open()
        }
    }
}
