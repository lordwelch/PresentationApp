import QtQuick 2.4
import QtQuick.Dialogs 1.1

MessageDialog {
    icon: StandardIcon.Information
    text: "Presentation App                                     \nVersion: Alpha"
    detailedText: "Presentation App for use in a church service\nMade in 2016 by Timmy Welch."
    title: "About"
    height: 100
    width: 200
    standardButtons: StandardButton.Close
}
