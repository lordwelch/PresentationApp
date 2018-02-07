import QtQuick 2.2
import QtQuick.Dialogs 1.0

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
