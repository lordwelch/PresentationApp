//https://gist.github.com/elpuri/3753756
import QtQuick 2.4
import QtQuick.Controls 1.6
import QtQuick.Controls.Styles 1.4

TreeView {
    id: view
    anchors.fill: parent
    anchors.margins: 2 * 12 + row.height
    model: colors
    alternatingRowColors: false
    style: TreeViewStyle {
      branchDelegate: Rectangle {
        width: 16
        height: 16
        color: styleData.isExpanded ? "green" : "red"
      }
      frame: Rectangle {border {color: "blue"}}
      backgroundColor: "blue"
    }

    TableViewColumn {
        title: "Name"
        role: "display"
        resizable: true
        delegate: Cell {
            text: "hell"
        }
    }
}

/*##^## Designer {
    D{i:0;autoSize:true;height:480;width:640}
}
 ##^##*/
