//https://gist.github.com/elpuri/3753756
import QtQuick 2.4

Item {
    id: tst4
    height: 50 + ((tst3.count - 1) * 50) + (tst3.subCount * 40)
    width: 200
    anchors.right: parent.right
    anchors.rightMargin: 0
    anchors.left: parent.left
    anchors.leftMargin: 0
    Component.onCompleted: {
     addLst()
    }

    function addLst() {
        var tstm
	tstm = nestedModel.get(0)
	tstm.subItems = [ { itemName: "test" }, { itemName: "notest" } ]

        nestedModel.append(tstm)

    }

    ListView {
        id: tst3
        anchors.fill: parent
        property int subCount: 0
        model: nestedModel
        delegate: Component {
            id: categoryDelegate
            Column {
                anchors.right: parent.right
                anchors.left: parent.left

                //width: 200
                Rectangle {
                    id: categoryItem
                    anchors.right: parent.right
                    anchors.left: parent.left
                    border.color: "black"
                    border.width: 5
                    color: "white"
                    height: 50

                    //width: 200
                    Text {
                        anchors.verticalCenter: parent.verticalCenter
                        x: 15
                        font.pixelSize: 24
                        text: categoryName
                    }

                    Rectangle {
                        color: "red"
                        width: 30
                        height: 30
                        anchors.right: parent.right
                        anchors.rightMargin: 15
                        anchors.verticalCenter: parent.verticalCenter

                        MouseArea {
                            anchors.fill: parent

                            // Toggle the 'collapsed' property
                            onClicked: {
                                nestedModel.setProperty(index, "collapsed", !collapsed)
                                if (!nestedModel.get(index).collapsed) {
                                    tst3.subCount = tst3.subCount + subItemLoader.subItemModel.count
                                } else {
                                    tst3.subCount = tst3.subCount - subItemLoader.subItemModel.count
                                }
                            }
                        }
                    }
                }

                Loader {
                    id: subItemLoader

                    // This is a workaround for a bug/feature in the Loader element. If sourceComponent is set to null
                    // the Loader element retains the same height it had when sourceComponent was set. Setting visible
                    // to false makes the parent Column treat it as if it's height was 0.
                    visible: !collapsed
                    property variant subItemModel: subItems
                    sourceComponent: subItemColumnDelegate
                    onStatusChanged: if (status == Loader.Ready)
                                         item.model = subItemModel
                }
            }
        }
    }

    Component {
        id: subItemColumnDelegate
        Column {
            property alias model: subItemRepeater.model

            width: tst4.width
            Repeater {
                id: subItemRepeater
                delegate: Rectangle {
                    color: "#cccccc"
                    height: 40
                    anchors.right: parent.right
                    anchors.left: parent.left
                    //width: 200
                    border.color: "black"
                    border.width: 2

                    Text {
                        anchors.verticalCenter: parent.verticalCenter
                        x: 30
                        font.pixelSize: 18
                        text: itemName
                    }
                }
            }
        }
    }
    ListModel {
        id: nestedModel
        objectName: "nestedModel"
        ListElement {
            categoryName: "Cars"
            collapsed: true
            subItems: [
                ListElement {
                    itemName: "Nisan"
                },
                ListElement {
                    itemName: "Toyota"
                },
                ListElement {
                    itemName: "Chevy"
                },
                ListElement {
                    itemName: "Audi"
                }
            ]
        }
        ListElement {
            categoryName: "Cars"
            collapsed: true
            subItems: [
                ListElement {
                    itemName: "Nissa"
                },
                ListElement {
                    itemName: "Toyota"
                },
                ListElement {
                    itemName: "Chevy"
                },
                ListElement {
                    itemName: "Audi"
                }
            ]
        }
        ListElement {
            categoryName: "Cars"
            collapsed: true
            subItems: [
                ListElement {
                    itemName: "Nissan"
                },
                ListElement {
                    itemName: "Toota"
                },
                ListElement {
                    itemName: "Chevy"
                },
                ListElement {
                    itemName: "Audi"
                }
            ]
        }
    }
}

