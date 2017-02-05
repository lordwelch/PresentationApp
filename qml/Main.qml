import QtQuick 2.4
import QtQuick.Dialogs 1.2
import QtQuick.Controls 1.3
import QtQuick.Window 2.0
import QtQuick.Layouts 1.0

ApplicationWindow {
    id: applicationWindow1
    //title: "Presentation App"
    visible: true
    objectName: "applicationWindow1"
    //minimumWidth: 500
    //minimumHeight: 500
    width: 1000
    height: 600
    property variant mlst

    Component.onCompleted: {
//        mlst = Qt.createComponent("Lst.qml").createObject(data1, {})
    }
    FileDialog {
        id: imgpicker
        title: "Choose an image for this slide"
        objectName: "imgpicker"
    }

    AboutDialog {
        id: aboutDialog
    }

    Action {
        id: aboutAction
        text: "About"
        onTriggered: aboutDialog.open()
    }

    Action {
        id: quitAction
        text: "Close"
        onTriggered: Qt.quit()
    }

    menuBar: MenuBar {
        Menu {
            title: "&File"
            MenuItem {
                action: quitAction
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
            onTriggered: imgpicker.open()
        }
    }

    SplitView {
        id: mainSlider
        anchors.fill: parent
        objectName: "mainSlider"
        orientation: Qt.Horizontal

        //onResizingChanged: col1.width = gridData.width / 2
        ScrollView {
            id: scview
            frameVisible: false
            anchors.margins: 4
            horizontalScrollBarPolicy: Qt.ScrollBarAlwaysOff
            verticalScrollBarPolicy: Qt.ScrollBarAlwaysOn
            flickableItem.boundsBehavior: Flickable.StopAtBounds

            Rectangle {
                id: col1
                width: scview.width - 22
                height: data1.childrenRect.height
                objectName: "col1"
                color: "#00000000"
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
                        if ((event.key === Qt.Key_Return)
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

                        if (event.key === Qt.Key_Escape) {
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
                    clip: true
                    height: data1.childrenRect.height
                }
            }
        }

        Rectangle {
            id: mainView
            border.width: 0
            objectName: "mainView"
            //anchors.left: scview.right
            z: 1
            clip: false
            visible: true

            Button {
                id: button1
                objectName: "btnAdd"
                x: 8
                y: 8
                text: "Button add"
//                onClicked: applicationWindow1.mlst.addLst("Nobody")
            }

            Button {
                id: button2
                x: 8
                y: 43
                text: "Button rem"
                objectName: "btnRem"
//                onClicked: applicationWindow1.mlst.remLst()
            }

            Button {
                id: button5
                x: 8
                y: 78
                text: "Button mem"
                objectName: "btnMem"
            }
        }
    }
}
