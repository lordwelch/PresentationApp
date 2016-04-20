import QtQuick 2.4
import QtQuick.Controls 1.3
import QtQuick.Window 2.0

ApplicationWindow {
    id: songEdit
    title: "Song Editor"
    visible: true
    objectName: "SongEdit"

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
                action: aboutAction
            }
        }
    }
    GroupBox {
        anchors.top: parent
        anchors.bottom: textArea
        ComboBox {
            id: verseSelector
            model: ["V1", "V2"]
            anchors.left: parent
            anchors.right: imageSelector
            anchors.top: parent
            anchors.bottom: parent
        }
        ComboBox {
            id: imageSelector
            anchors.left: verseSelector
            anchors.right: parent
            anchors.top: parent
            anchors.bottom: parent
        }
    }
    SplitView {

        ListView {
            width: 180
            height: 200
            boundsBehavior: Flickable.StopAtBounds

            model: ListModel {
                ListElement {
                    name: "v1"
                }
                ListElement {
                    name: "v2"
                }
                ListElement {
                    name: "v3"
                }
                ListElement {
                    name: "v4"
                }
                ListElement {
                    name: "v5"
                }
            }
            delegate: Component {
                id: contactsDelegate
                Text {
                    id: contactInfo
                    text: name
                }
            }
            focus: true
        }
        ListView {
            width: 180
            height: 200
            boundsBehavior: Flickable.StopAtBounds

            model: ListModel {
                ListElement {
                    name: "v1"
                }
                ListElement {
                    name: "v2"
                }
                ListElement {
                    name: "v3"
                }
                ListElement {
                    name: "v4"
                }
                ListElement {
                    name: "v5"
                }
            }
            delegate: Component {
                id: contactsDelegate1
                Text {
                    id: contactInfo
                    text: name
                }
            }
            focus: true
        }
    }
    TextArea {
            id: textArea
                }
}
