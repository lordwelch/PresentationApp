import QtQuick 2.4
import QtQuick.Controls 1.4
import QtQuick.Layouts 1.1
import QtQuick.Dialogs 1.2
//import Qt.labs.controls 1.0

ApplicationWindow {
    minimumHeight: 480
    minimumWidth: 640

    ColorDialog {
        id: textClrDialog
        //objectname: "textClrDialog"
        // @disable-check M16
        title: "Please choose a color for the text"
        // @disable-check M16
        showAlphaChannel: true
    }

    ColorDialog {
        id: outlineClrDialog
        //objectname: "outlineClrDialog"
        // @disable-check M16
        title: "Please choose a color for the text"
        // @disable-check M16
        showAlphaChannel: true
    }

    menuBar: MenuBar {
        Menu {
            title: "&File"
            MenuItem {
                text: "Close"
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
                //action: aboutAction
            }
        }
    }

    RowLayout {
        id: rowLayout1
        enabled: true
        smooth: true
        antialiasing: true
        anchors.fill: parent

        RowLayout {
            id: rowlayout3
            Layout.fillHeight: true
            Layout.alignment: Qt.AlignTop
            Layout.maximumWidth: 225

            ColumnLayout {
                id: columnlayout2
                Layout.alignment: Qt.AlignLeft | Qt.AlignTop
                Layout.fillHeight: true

                Label {
                    id: label1
                    text: qsTr("Verses")
                    Layout.alignment: Qt.AlignLeft | Qt.AlignTop
                }
                ListView {
                    id: listView1
                    clip: true
                    highlight: Rectangle {
                        color: "lightsteelblue"
                        radius: 5
                    }
                    width: 110
                    Layout.fillHeight: true
                    Layout.alignment: Qt.AlignLeft | Qt.AlignTop
                    focus: true
                    keyNavigationWraps: true
                    boundsBehavior: Flickable.StopAtBounds
                    model: GO.verseLen

                    delegate: Item {
                        x: 5
                        width: 80
                        height: 40

                        Text {
                            text: GO.verses(index)
                            anchors.verticalCenter: parent.verticalCenter
                            font.bold: true
                        }
                    }
                }
            }

            ColumnLayout {
                id: columnlayout3
                Layout.alignment: Qt.AlignLeft | Qt.AlignTop
                Layout.fillHeight: true

                Label {
                    id: label2
                    text: qsTr("Verse Order")
                    Layout.alignment: Qt.AlignLeft | Qt.AlignTop
                }
                ListView {
                    id: listView2
                    clip: true
                    highlight: Rectangle {
                        color: "lightsteelblue"
                        radius: 5
                    }
                    width: 110
                    Layout.fillHeight: true
                    Layout.alignment: Qt.AlignLeft | Qt.AlignTop
                    boundsBehavior: Flickable.StopAtBounds
                    model: GO.orderLen
                    delegate: Item {
                        x: 5
                        width: 80
                        height: 40

                        Text {
                            text: GO.verseOrder(index)
                            anchors.verticalCenter: parent.verticalCenter
                            font.bold: true
                        }
                    }
                }
            }
        }

        ColumnLayout {
            id: columnlayout4
            Layout.fillWidth: true
            Layout.alignment: Qt.AlignLeft | Qt.AlignTop
            Layout.fillHeight: true

            RowLayout {
                id: rowLayout3
                Layout.alignment: Qt.AlignLeft | Qt.AlignTop
                Layout.maximumHeight: 30
                Layout.minimumHeight: 30
                Layout.preferredHeight: 30
                Layout.fillWidth: true

                ToolButton {
                    id: textColorPicker
                    objectName: "textColorPicker"
                    text: "Text Color"
                    Layout.alignment: Qt.AlignLeft | Qt.AlignTop
                    tooltip: "Pick the color of the text"
                }

                ToolButton {
                    id: outlineColorPicker
                    objectName: "outlineColorPicker"
                    text: "Outline Color"
                    Layout.alignment: Qt.AlignLeft | Qt.AlignTop
                    tooltip: "Pick the color of the text outline"
                }

                ComboBox {
                    id: fontPicker
                    objectName: "fontPicker"
                    Layout.alignment: Qt.AlignLeft | Qt.AlignTop
                    model: GO.fontLen
                    // @disable-check M16
                    /*delegate:Text {
                        text: GO.fontList(index)
                    }*/

                }

                SpinBox {
                    id: fontSize
                    Layout.alignment: Qt.AlignLeft | Qt.AlignTop
                    objectName: "fontSize"
                    maximumValue: 1000
                    value: 1
                    suffix: "Pt"
                }

                SpinBox {
                    id: outlineSize
                    stepSize: 0.1
                    decimals: 1
                    Layout.alignment: Qt.AlignLeft | Qt.AlignTop
                    objectName: "outlineSize"
                    maximumValue: 10
                    value: 1
                }
            }
            RowLayout {
                id: rowLayout2
                Layout.preferredHeight: 30
                Layout.maximumHeight: 30
                Layout.minimumHeight: 30
                Layout.alignment: Qt.AlignLeft | Qt.AlignTop
                Layout.fillHeight: true
                Layout.fillWidth: true

                ComboBox {
                    id: versePicker
                    objectName: "versePicker"
                    model: GO.verseLen
                    // @disable-check M16
                   /* delegate: Text {
                        text: GO.verses(index)
                    }*/
                }

                ComboBox {
                    id: imgPicker
                    objectName: "imgPicker"
                    model: GO.imgLen
                    // @disable-check M16
                    /*delegate: Text {
                        text: GO.img(index)
                    }*/
                }
                TextArea {
                    id: textEdit1
                    width: 80
                    height: 20
                    text: qsTr("Text Edit")
                    textFormat: Text.AutoText
                    Layout.fillHeight: true
                    Layout.fillWidth: true
                    Layout.alignment: Qt.AlignLeft | Qt.AlignTop
                    font.pixelSize: 12
                    selectByKeyboard: true
                    selectByMouse: true
                }
            }
        }
    }
}
