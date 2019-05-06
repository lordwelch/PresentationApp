package main

// This file is automatically generated by github.com/limetext/qml-go/cmd/genqrc

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/limetext/qml-go"
)

func init() {
	var r *qml.Resources
	var err error
	err = qrcRepackResources()
	if err != nil {
		panic("cannot repack qrc resources: " + err.Error())
	}
	r, err = qml.ParseResources(qrcResourcesRepacked)
	if err != nil {
		panic("cannot parse bundled resources data: " + err.Error())
	}
	qml.LoadResources(r)
}

func qrcRepackResources() error {
	subdirs := []string{"qml"}
	var rp qml.ResourcesPacker
	for _, subdir := range subdirs {
		err := filepath.Walk(subdir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				return nil
			}
			data, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			if filepath.Ext(info.Name()) == ".conf" {
				rp.Add("/", data)
			} else {
				rp.Add(filepath.ToSlash(path), data)
			}
			return nil
		})
		if err != nil {
			return err
		}
	}
	qrcResourcesRepacked = rp.Pack().Bytes()
	return nil
}

var qrcResourcesRepacked []byte
var qrcResourcesData = "qres\x00\x00\x00\x01\x00\x00X\xed\x00\x00\x00\x14\x00\x00W\xdb\x00\x00\x026import QtQuick 2.4\n\nListModel {\n    id: nestedModel\n    objectName: \"nestedModel\"\n    function get1() {\n        console.log(get(0))\n        return get(0)\n    }\n    ListElement {\n        title: \"Cars\"\n        collapsed: true\n        subItems: [\n            ListElement {\n                itemName: \"tst\"\n            },\n            ListElement {\n                itemName: \"Tota\"\n            },\n            ListElement {\n                itemName: \"vy\"\n            },\n            ListElement {\n                itemName: \"Audio Adrenaline\"\n            }\n        ]\n    }\n}\n\x00\x00\x01\xfdimport QtQuick 2.4\n\nImage {\n    id: img\n    antialiasing: true\n    source: \"image://images/\"\n    objectName: \"cellImg\"\n    property int index: 0\n    height: 100\n    transformOrigin: Item.TopLeft\n    fillMode: Image.PreserveAspectFit\n    anchors.right: parent.right\n    anchors.left: parent.left\n    //cache: false\n    MouseArea {\n        id: cellMouse\n        hoverEnabled: true\n        enabled: true\n        objectName: \"cellMouse\"\n        anchors.fill: parent\n        acceptedButtons: Qt.AllButtons\n    }\n}\n\x00\x00\x1e.import QtQuick 2.4\nimport QtQuick.Controls 1.3\nimport QtQuick.Layouts 1.1\nimport QtQuick.Dialogs 1.2\n\nApplicationWindow {\n    minimumHeight: 480\n    minimumWidth: 640\n\n    ColorDialog {\n        id: textClrDialog\n        //objectname: \"textClrDialog\"\n        title: \"Please choose a color for the text\"\n        showAlphaChannel: true\n    }\n\n    ColorDialog {\n        id: outlineClrDialog\n        //objectname: \"outlineClrDialog\"\n        title: \"Please choose a color for the text\"\n        showAlphaChannel: true\n    }\n\n    menuBar: MenuBar {\n        Menu {\n            title: \"&File\"\n            MenuItem {\n                text: \"Close\"\n            }\n        }\n        Menu {\n            title: \"&Edit\"\n            MenuItem {\n                text: \"quick edit\"\n                objectName: \"mnuEdit\"\n            }\n        }\n\n        Menu {\n            title: \"Window\"\n\n            MenuItem {\n                text: \"Display\"\n                objectName: \"mnuDisplay\"\n            }\n        }\n\n        Menu {\n            title: \"&Help\"\n            MenuItem {\n                //action: aboutAction\n            }\n        }\n    }\n\n    RowLayout {\n        id: rowLayout1\n        enabled: true\n        smooth: true\n        antialiasing: true\n        anchors.fill: parent\n\n        RowLayout {\n            id: rowlayout3\n            Layout.fillHeight: true\n            Layout.alignment: Qt.AlignTop\n            Layout.maximumWidth: 225\n\n            ColumnLayout {\n                id: columnlayout2\n                Layout.alignment: Qt.AlignLeft | Qt.AlignTop\n                Layout.fillHeight: true\n\n                Label {\n                    id: label1\n                    text: qsTr(\"Verses\")\n                    Layout.alignment: Qt.AlignLeft | Qt.AlignTop\n                }\n                ListView {\n                    id: lstVerses\n                    objectName: \"lstVerses\"\n                    clip: true\n                    highlight: Rectangle {\n                        color: \"lightsteelblue\"\n                        radius: 5\n                    }\n                    width: 110\n                    Layout.fillHeight: true\n                    Layout.alignment: Qt.AlignLeft | Qt.AlignTop\n                    focus: true\n                    keyNavigationWraps: true\n                    boundsBehavior: Flickable.StopAtBounds\n                    model: go.verseLen\n\n                    delegate: Item {\n                        x: 5\n                        width: 80\n                        height: 40\n\n                        Text {\n                            text: go.verses(index)\n                            anchors.verticalCenter: parent.verticalCenter\n                            font.bold: true\n                        }\n                    }\n                }\n            }\n\n            ColumnLayout {\n                id: columnlayout3\n                Layout.alignment: Qt.AlignLeft | Qt.AlignTop\n                Layout.fillHeight: true\n\n                Label {\n                    id: label2\n                    text: qsTr(\"Verse Order\")\n                    Layout.alignment: Qt.AlignLeft | Qt.AlignTop\n                }\n                ListView {\n                    id: lstOrder\n                    objectName: \"lstOrder\"\n                    clip: true\n                    highlight: Rectangle {\n                        color: \"lightsteelblue\"\n                        radius: 5\n                    }\n                    width: 110\n                    Layout.fillHeight: true\n                    Layout.alignment: Qt.AlignLeft | Qt.AlignTop\n                    boundsBehavior: Flickable.StopAtBounds\n                    model: go.orderLen\n                    delegate: Item {\n                        x: 5\n                        width: 80\n                        height: 40\n\n                        Text {\n                            text: go.verseOrder(index)\n                            anchors.verticalCenter: parent.verticalCenter\n                            font.bold: true\n                        }\n                    }\n                }\n            }\n        }\n\n        ColumnLayout {\n            id: columnlayout4\n            Layout.fillWidth: true\n            Layout.alignment: Qt.AlignLeft | Qt.AlignTop\n            Layout.fillHeight: true\n\n            RowLayout {\n                id: rowLayout3\n                Layout.alignment: Qt.AlignLeft | Qt.AlignTop\n                Layout.maximumHeight: 30\n                Layout.minimumHeight: 30\n                Layout.preferredHeight: 30\n                Layout.fillWidth: true\n\n                ToolButton {\n                    id: textColorPicker\n                    objectName: \"textColorPicker\"\n                    text: \"Text Color\"\n                    Layout.alignment: Qt.AlignLeft | Qt.AlignTop\n                    tooltip: \"Pick the color of the text\"\n                }\n\n                ToolButton {\n                    id: outlineColorPicker\n                    objectName: \"outlineColorPicker\"\n                    text: \"Outline Color\"\n                    Layout.alignment: Qt.AlignLeft | Qt.AlignTop\n                    tooltip: \"Pick the color of the text outline\"\n                }\n\n                ComboBox {\n                    id: fontPicker\n                    objectName: \"fontPicker\"\n                    Layout.alignment: Qt.AlignLeft | Qt.AlignTop\n                    model: go.fontList.split(\"\\n\")\n                    /*// @disable-check M16\n                    delegate:Text {\n                        text: go.fontList(index)\n                    }*/\n                }\n\n                SpinBox {\n                    id: fontSize\n                    objectName: \"fontSize\"\n                    Layout.alignment: Qt.AlignLeft | Qt.AlignTop\n                    maximumValue: 1000\n                    value: 1\n                    suffix: \"Pt\"\n                }\n\n                SpinBox {\n                    id: outlineSize\n                    stepSize: 0.1\n                    decimals: 1\n                    Layout.alignment: Qt.AlignLeft | Qt.AlignTop\n                    objectName: \"outlineSize\"\n                    maximumValue: 10\n                    value: 1\n                }\n            }\n            RowLayout {\n                id: rowLayout2\n                Layout.preferredHeight: 30\n                Layout.maximumHeight: 30\n                Layout.minimumHeight: 30\n                Layout.alignment: Qt.AlignLeft | Qt.AlignTop\n                Layout.fillHeight: true\n                Layout.fillWidth: true\n\n                ComboBox {\n                    id: versePicker\n                    objectName: \"versePicker\"\n                    model: go.verses.split(\"\\n\")\n                    /*// @disable-check M16\n                    delegate: Text {\n                        text: go.verses(index)\n                    }*/\n                }\n\n                ComboBox {\n                    id: imgPicker\n                    objectName: \"imgPicker\"\n                    //model: go.img.split(\"\\n\")\n                    /*// @disable-check M16\n                    delegate: Text {\n                        text: go.img(index)\n                    }*/\n                }\n                TextArea {\n                    id: txtVerse\n                    objectName: \"txtVerse\"\n                    width: 80\n                    height: 20\n                    text: qsTr(\"Text Edit\")\n                    textFormat: Text.AutoText\n                    Layout.fillHeight: true\n                    Layout.fillWidth: true\n                    Layout.alignment: Qt.AlignLeft | Qt.AlignTop\n                    font.pixelSize: 12\n                    selectByKeyboard: true\n                    selectByMouse: true\n                }\n            }\n        }\n    }\n}\n\x00\x00\x00\xebimport QtQuick 2.2\nimport QtQuick.Dialogs 1.0\n\nFileDialog {\n    id: imgDialog\n    title: \"Please choose an image\"\n    folder: shortcuts.home\n    onAccepted: {\n\n    }\n    onRejected: {\n\n    }\n    Component.onCompleted: visible = true\n}\n\x00\x00\x01nimport QtQuick 2.4\nimport QtQuick.Controls 1.3 as Quick\nimport QtQuick.Controls 2.1\n\nApplicationWindow {\n    flags: Qt.MaximumSize\n    Component.onCompleted: visible = true\n\n    Image {\n        id: image1\n        objectName: \"displayImage\"\n        sourceSize.height: 768\n        sourceSize.width: 1024\n        antialiasing: true\n        anchors.fill: parent\n    }\n}\n\x00\x00\x12<import QtQuick 2.3\nimport QtQuick.Controls 1.3 as Quick\nimport QtQuick.Controls 2.1\nimport QtQuick.Dialogs 1.2\nimport QtQuick.Window 2.0\nimport QtQuick.Layouts 1.0\n\nApplicationWindow {\n    id: applicationWindow1\n    title: \"Presentation App\"\n    visible: true\n    objectName: \"applicationWindow1\"\n    minimumWidth: 500\n    minimumHeight: 500\n\n    FileDialog {\n        id: imgpicker\n        objectName: \"imgpicker\"\n        title: \"Choose an image for this slide\"\n    }\n\n    Quick.SplitView {\n        id: spview\n        anchors.fill: parent\n        Rectangle {\n            id: col1\n            x: 300\n            y: 0\n            objectName: \"col1\"\n            color: \"#00000000\"\n            transformOrigin: Item.TopLeft\n            border.width: 0\n            Layout.fillWidth: true\n\n            Rectangle {\n                id: textEdit\n                property int cell\n                x: 232\n                y: 622\n                objectName: \"textEdit\"\n                width: 200\n                height: 200\n                color: \"#ffffff\"\n                visible: false\n                property bool txt: true\n                Keys.onPressed: {\n                    if ((event.key == Qt.Key_Return)\n                            && (event.modifiers & Qt.ControlModifier)) {\n                        txt = true\n\n                        x = -100\n                        y = -100\n                        visible = false\n                        focus = true\n                        enabled = false\n                        opacity = 0\n                        textEdit1.focus = false\n\n                        event.accepted = true\n                    }\n\n                    if (event.key == Qt.Key_Escape) {\n                        txt = false\n                        x = -100\n                        y = -100\n                        visible = false\n                        focus = true\n                        enabled = false\n                        opacity = 0\n                        textEdit1.focus = false\n\n                        event.accepted = true\n                    }\n                }\n\n                TextArea {\n                    id: textEdit1\n                    objectName: \"textEdit1\"\n                    anchors.fill: parent\n                    clip: true\n                    textFormat: Text.AutoText\n                    visible: true\n                    font.pixelSize: 12\n                    z: 99\n                }\n            }\n            Flickable {\n                id: scview\n                objectName: \"scview\"\n                anchors.fill: parent\n                boundsBehavior: Flickable.OvershootBounds\n                flickableDirection: Flickable.VerticalFlick\n                pixelAligned: true\n                //verticalScrollBarPolicy: Qt.ScrollBarAlwaysOn\n                //horizontalScrollBarPolicy: Qt.ScrollBarAlwaysOff\n                //highlightOnFocus: false\n                //frameVisible: true\n                contentHeight: contentItem.childrenRect.height\n            }\n        }\n\n        Rectangle {\n            id: mainView\n            objectName: \"mainView\"\n            clip: false\n            visible: true\n            Layout.minimumWidth: 300\n            Layout.fillWidth: false\n\n            Button {\n                id: button1\n                objectName: \"btnAdd\"\n                x: 8\n                y: 8\n                text: qsTr(\"Add\")\n                onClicked: sv.addLst(\"fail\")\n            }\n\n            Button {\n                id: button2\n                x: 8\n                y: 43\n                text: qsTr(\"Remove\")\n                objectName: \"btnRem\"\n            }\n\n            Button {\n                id: button3\n                x: 8\n                y: 78\n                text: qsTr(\"Button \")\n                objectName: \"btnMem\"\n            }\n        }\n    }\n    /*\n    menuBar: MenuBar {\n        Menu {\n            title: \"&File\"\n            MenuItem {\n                text: \"Close\"\n                shortcut: StandardKey.Quit\n            }\n        }\n        Menu {\n            title: \"&Edit\"\n            MenuItem {\n                text: \"quick edit\"\n                objectName: \"mnuEdit\"\n            }\n        }\n\n        Menu {\n            title: \"Window\"\n\n            MenuItem {\n                text: \"Display\"\n                objectName: \"mnuDisplay\"\n            }\n        }\n\n        Menu {\n\n            MenuItem {\n                text: \"&help\"\n            }\n        }\n    }\n*/\n    Menu {\n        objectName: \"mnuCtx\"\n        title: \"new image...\"\n        MenuItem {\n            objectName: \"mnuImgPick\"\n            text: \"new Image...\"\n            onTriggered: imgpicker.open()\n        }\n    }\n}\n\x00\x00\x00\x1a[Controls]\nStyle=Material\n\x00\x00\r\rimport QtQuick 2.3\nimport QtQuick.Controls 1.3 as Quick\nimport QtQuick.Layouts 1.0\n\nRectangle {\n    id: itm\n    height: 100\n    anchors.right: parent.right\n    anchors.left: parent.left\n    Quick.SplitView {\n        id: splitView\n        anchors.fill: parent\n        Rectangle {\n            id: half1\n            height: 100\n            Layout.fillWidth: true\n            Layout.minimumWidth: 100\n            Rectangle {\n                objectName: \"cellRect\"\n                property int index: 0\n                anchors.fill: parent\n                border.width: 2\n                border.color: \"black\"\n                Text {\n                    id: cellText\n                    enabled: true\n                    objectName: \"cellText\"\n                    text: \"itm.model.text\"\n                    renderType: Text.NativeRendering\n                    clip: true\n                    wrapMode: Text.WrapAtWordBoundaryOrAnywhere\n                    anchors.fill: parent\n                    anchors.right: parent.right\n                    anchors.rightMargin: 0\n                    anchors.left: parent.left\n                    anchors.leftMargin: 2\n\n                    MouseArea {\n                        id: cellMouse\n                        hoverEnabled: true\n                        enabled: true\n                        objectName: \"cellMouse\"\n                        anchors.fill: parent\n                        acceptedButtons: Qt.AllButtons\n\n                        onMouseXChanged: cellHover()\n                        onExited: focusChanged(focus)\n\n                        function cellHover() {\n                            if (containsMouse) {\n                                parent.parent.border.color = \"skyblue\"\n                                parent.parent.color = \"darkblue\"\n                                parent.color = \"white\"\n                            } else if (focus) {\n                                parent.color = \"black\"\n                            }\n                        }\n\n                        function notSelected() {\n\n                            parent.parent.border.color = \"black\"\n                            parent.parent.color = \"white\"\n                            parent.color = \"black\"\n                            cellHover()\n                        }\n\n                        function selected() {\n                            parent.parent.border.color = \"blue\"\n                            parent.color = \"black\"\n                            parent.parent.color = \"gainsboro\"\n                            cellHover()\n                        }\n                    }\n                }\n            }\n        }\n        Rectangle {\n            id: half3\n            Layout.fillWidth: false\n            Layout.minimumWidth: 100\n            Layout.maximumHeight: 400\n            onXChanged: {\n                var temp = (cellImg.sourceSize.height / cellImg.sourceSize.width) * half3.width\n                if (temp < 100) {\n                    temp = 100\n                }\n                if (temp > 150) {\n                    temp = 150\n                }\n\n                itm.height = temp\n            }\n\n            Img {\n                id: cellImg\n                anchors.fill: parent\n                objectName: \"cellImg\"\n                source: \"itm.model.imageSource\"\n            }\n        }\n    }\n}\n\x00\x00\x13\x86//https://gist.github.com/elpuri/3753756\nimport QtQuick 2.4\n\nItem {\n    id: rt\n    property ListElement def: ListElement {\n        property string cellText: \"DOCTYPE html PUBLIC\"\n        property int collectionIndex: 0\n        property string imageSource: \"image://images/list:;cell:\"\n    }\n\n    Component.onCompleted: addLst(\"Haha :-P\")\n    height: lst.contentHeight\n    anchors.right: parent.right\n    anchors.left: parent.left\n    anchors.leftMargin: 0\n    function remove(List, index) {\n        lst.subCount--\n        lst.model[List].subItems.remove(index, 1)\n    }\n\n    function pop(List) {\n        lst.subCount--\n        lst.model[List].subItems.remove(lst.model[List].subItems.count - 1, 1)\n    }\n    function newdef(index, txt, src) {\n        var item = Object.create(def)\n        item.collectionIndex = index\n        item.title = txt\n        item.imageSource = src\n        return item\n    }\n    function remLst() {\n        lst.model.remove(lst.model.count - 1, 1)\n    }\n\n    function apppend(List, obj) {\n        lst.subCount++\n        lst.model[List].subItems.append(obj)\n    }\n\n    function insert(List, index, obj) {\n        lst.subCount++\n        lst.model[List].subItems.insert(index, obj)\n    }\n\n    function get(List, index) {\n        return lst.model[List].subItems.get(index)\n    }\n\n    function set(List, index, obj) {\n        lst.model[List].subItems.set(index, obj)\n    }\n\n    function addLst(str) {\n        var newCollection\n        var count = 2\n        var i = 0\n        var temp = Qt.createComponent(\"Sublist.qml\").createObject(rt, {\n\n                                                                  })\n\n        newCollection = temp.get(0)\n        newCollection.title = str\n        newCollection.subItems.clear()\n        for (i = 0; i < count; i++) {\n            newCollection.subItems.append(newdef(lst.model.count, \"idiot\"))\n        }\n\n        lst.model.append(newCollection)\n    }\n\n    ListView {\n        id: lst\n        anchors.fill: parent\n        y: 0\n        interactive: false\n        model: nestedModel\n        delegate: Component {\n            id: categoryDelegate\n            Column {\n                anchors.right: parent.right\n                anchors.left: parent.left\n\n                Rectangle {\n                    id: categoryItem\n                    anchors.right: parent.right\n                    anchors.left: parent.left\n                    border.color: \"black\"\n                    border.width: 5\n                    color: \"white\"\n                    height: 50\n\n                    Text {\n                        anchors.verticalCenter: parent.verticalCenter\n                        x: 15\n                        font.pixelSize: 24\n                        text: index + ' ' + cellText\n                        clip: true\n                        anchors.left: parent.left\n                        anchors.right: parent.right\n                        anchors.rightMargin: 15\n                        anchors.leftMargin: 5\n                    }\n\n                    Rectangle {\n                        color: \"red\"\n                        width: 30\n                        height: 30\n                        anchors.right: parent.right\n                        anchors.rightMargin: 15\n                        anchors.verticalCenter: parent.verticalCenter\n\n                        MouseArea {\n                            anchors.fill: parent\n\n                            // Toggle the 'collapsed' property\n                            onClicked: {\n                                lst.model.setProperty(index, \"collapsed\",\n                                                      !collapsed)\n                            }\n                        }\n                    }\n                }\n\n                Loader {\n                    id: subItemLoader\n\n                    // This is a workaround for a bug/feature in the Loader element. If sourceComponent is set to null\n                    // the Loader element retains the same height it had when sourceComponent was set. Setting visible\n                    // to false makes the parent Column treat it as if it's height was 0.\n                    visible: !collapsed\n                    property variant subItemModel: subItems\n                    sourceComponent: subItemColumnDelegate\n                    onStatusChanged: if (status == Loader.Ready) {\n                                         item.model = subItemModel\n                                     }\n                }\n            }\n        }\n    }\n\n    Component {\n        id: subItemColumnDelegate\n        Column {\n            property alias model: subItemRepeater.model\n\n            width: rt.width\n            Repeater {\n                anchors.right: parent.right\n                anchors.left: parent.left\n                id: subItemRepeater\n                objectName: \"repeater\"\n                delegate: Cell {\n                }\n            }\n        }\n    }\n    ListModel {\n        id: nestedModel\n        objectName: \"nestedModel\"\n    }\n}\n\x00\x03\x00\x00x<\x00q\x00m\x00l\x00\v\x00\x12&\\\x00S\x00u\x00b\x00l\x00i\x00s\x00t\x00.\x00q\x00m\x00l\x00\a\x00:X\x9c\x00I\x00m\x00g\x00.\x00q\x00m\x00l\x00\f\x00K\xc3\\\x00S\x00o\x00n\x00g\x00E\x00d\x00i\x00t\x00.\x00q\x00m\x00l\x00\x0f\x00\xd2\f\\\x00f\x00i\x00l\x00e\x00D\x00i\x00a\x00l\x00o\x00g\x00s\x00.\x00q\x00m\x00l\x00\v\x02\x1eD\xdc\x00D\x00i\x00s\x00p\x00l\x00a\x00y\x00.\x00q\x00m\x00l\x00\b\b\x01^\\\x00M\x00a\x00i\x00n\x00.\x00q\x00m\x00l\x00\x15\b\x1e\x16f\x00q\x00t\x00q\x00u\x00i\x00c\x00k\x00c\x00o\x00n\x00t\x00r\x00o\x00l\x00s\x002\x00.\x00c\x00o\x00n\x00f\x00\b\f/a\x1c\x00C\x00e\x00l\x00l\x00.\x00q\x00m\x00l\x00\v\x0f'Ǽ\x00S\x00e\x00r\x00v\x00i\x00c\x00e\x00.\x00q\x00m\x00l\x00\x00\x00\x00\x00\x02\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x00\x00\x02\x00\x00\x00\t\x00\x00\x00\x02\x00\x00\x00\f\x00\x00\x00\x00\x00\x01\x00\x00\x00\x00\x00\x00\x00(\x00\x00\x00\x00\x00\x01\x00\x00\x02:\x00\x00\x00<\x00\x00\x00\x00\x00\x01\x00\x00\x04;\x00\x00\x00Z\x00\x00\x00\x00\x00\x01\x00\x00\"m\x00\x00\x00~\x00\x00\x00\x00\x00\x01\x00\x00#\\\x00\x00\x00\x9a\x00\x00\x00\x00\x00\x01\x00\x00$\xce\x00\x00\x00\xb0\x00\x00\x00\x00\x00\x01\x00\x007\x0e\x00\x00\x00\xe0\x00\x00\x00\x00\x00\x01\x00\x007,\x00\x00\x00\xf6\x00\x00\x00\x00\x00\x01\x00\x00D="
