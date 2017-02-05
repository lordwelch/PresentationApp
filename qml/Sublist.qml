import QtQuick 2.4
import QtQuick.Controls 1.3

ListModel {
  id: nestedModel1
  objectName: "nestedModel1"
  ListElement {
    name: "Cars"
    collapsed: true
    subItems: [
      ListElement {
        itemName: "idiot"
      }
    ]
  }
}
