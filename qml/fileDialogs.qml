import QtQuick 2.4
import QtQuick.Dialogs 1.3

FileDialog {
    id: imgDialog
    title: "Please choose an image"
    folder: shortcuts.home
    onAccepted: {

    }
    onRejected: {

    }
    Component.onCompleted: visible = true
}
