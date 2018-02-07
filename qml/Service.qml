//https://gist.github.com/elpuri/3753756
import QtQuick 2.4

Item {
    id: rt
    property ListElement def: ListElement {
        property string cellText: "<!DOCTYPE html PUBLIC \"-//W3C//DTD XHTML 1.0 Transitional//EN\"

\"http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd\">
<html >

<head>
<meta content=\"text/html; charset=utf-8\" http-equiv=\"Content-Type\" />
<title>Untitled 1</title>
<style type=\"text/css\">
.auto-style2 {
font-family: \"Times New Roman\", Times, serif;
font-size: small;
}
.auto-style3 {
background-color: #FFFF00;
}
</style>
</head>

<body>
<p><b>Header text</b><br/></p>
<span class=\"auto-style3\">This is paragraph text</span>

<hr />
</body>

</html>"
        property int collectionIndex: 0
        property string imageSource: "image://images/list:;cell:"
    }

    Component.onCompleted: addLst("Haha :-P")
    height: ((lst.count) * 50) + (lst.subCount * 100)

    anchors.right: parent.right
    anchors.left: parent.left
    anchors.leftMargin: 0
    function remove(List, index) {
        lst.subCount--
        nestedModel.get(List).subItems.remove(index, 1)
    }

    function pop(List) {
        lst.subCount--
        nestedModel.get(List).subItems.remove(nestedModel.get(
                                                  List).subItems.count - 1, 1)
    }
    function newdef(index, txt, src) {
        var item = Object.create(def)
        item.collectionIndex = index
        item.text = txt
        item.imageSource = src
        return item
    }
    function remLst() {
        nestedModel.remove(nestedModel.count - 1, 1)
    }

    function apppend(List, obj) {
        lst.subCount++
        nestedModel.get(List).subItems.append(obj)
    }

    function insert(List, index, obj) {
        lst.subCount++
        nestedModel.get(List).subItems.insert(index, obj)
    }

    function get(List, index) {
        return nestedModel.get(List).subItems.get(index)
    }

    function set(List, index, obj) {
        nestedModel.get(List).subItems.set(index, obj)
    }

    function addLst(str) {
        var newCollection
        var i = 0
        var temp = Qt.createComponent("Sublist.qml").createObject(rt, {

                                                                  })

        newCollection = temp.get(0)
        newCollection.name = str
        newCollection.subItems.clear()
        for (i = 0; i < 1; i++) {
            newCollection.subItems.append(newdef(nestedModel.count, "idiot"))
        }

        nestedModel.append(newCollection)
    }

    ListView {
        id: lst
        anchors.fill: parent
        y: 0
        height: ((lst.count) * 55) + (lst.subCount * 100)
        interactive: false
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
                        text: name
                        clip: true
                        anchors.left: parent.left
                        anchors.right: parent.right
                        anchors.rightMargin: 15
                        anchors.leftMargin: 5
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
                                nestedModel.setProperty(index, "collapsed",
                                                        !collapsed)
                                if (!nestedModel.get(index).collapsed) {
                                    lst.subCount = lst.subCount + subItemLoader.subItemModel.count
                                } else {
                                    lst.subCount = lst.subCount - subItemLoader.subItemModel.count
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
                    onStatusChanged: if (status == Loader.Ready) {
                                         item.model = subItemModel
                                     }
                }
            }
        }
    }

    Component {
        id: subItemColumnDelegate
        Column {
            property alias model: subItemRepeater.model

            width: rt.width
            Repeater {
                id: subItemRepeater
                objectName: "repeater"
                delegate: Cell {
                }
            }
        }
    }
    ListModel {
        id: nestedModel
        objectName: "nestedModel"
    }
}
