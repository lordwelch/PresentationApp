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

    SplitView {
        ListView {
            width: 180; height: 200
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
                ListElement {
                    name: "v6"
                }
                ListElement {
                    name: "v7"
                }
                ListElement {
                    name: "v8"
                }
                ListElement {
                    name: "v9"
                }
                ListElement {
                    name: "v10"
                }
                ListElement {
                    name: "v11"
                }
                ListElement {
                    name: "v12"
                }
                ListElement {
                    name: "v13"
                }
                ListElement {
                    name: "v14"
                }
                ListElement {
                    name: "v15"
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
    }
}
