import QtQuick 2.4

ListModel {
    id: nestedModel
    objectName: "nestedModel"
    function get1() {
        console.log(get(0))
        return get(0)
    }
    ListElement {
        title: "Cars"
        collapsed: true
        subItems: [
            ListElement {
                itemName: "tst"
            },
            ListElement {
                itemName: "Tota"
            },
            ListElement {
                itemName: "vy"
            },
            ListElement {
                itemName: "Audio Adrenaline"
            }
        ]
    }
}
