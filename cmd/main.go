package main

import (
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
	"github.com/therecipe/qt/widgets"
)

func init() {
	CustomTreeModel_QmlRegisterType2("CustomQmlTypes", 1, 0, "CustomTreeModel")
}

const something = 1<<31 - 1

const (
	FirstName = int(core.Qt__UserRole) + 1<<iota
	LastName
)

type Collection struct {
	title       string
	_childItems []CollectionItem
}

type CollectionItem struct {
	_itemData []string
}

type ROOT []Collection

func (r *ROOT) add(t Collection) {
	*r = append(*r, t)
}

func (r *ROOT) Len() int {
	return len(*r)
}

func (r *ROOT) CollectionLen(i int) int {
	return len((*r)[i]._childItems)
}

func (r *ROOT) Collection(index int) Collection {
	return (*r)[index]
}

func (r *ROOT) CollectionItem(index int, Cindex int) CollectionItem {
	return (*r)[index]._childItems[Cindex]
}

func NewCollectionItem(data []string) CollectionItem {
	return CollectionItem{
		_itemData: data,
	}
}

type service interface {
	add(t Collection)
	Len() int
	Collection(index int) Collection
	CollectionLen(i int) int
	CollectionItem(index int, listIndex int) CollectionItem
}

type CustomTreeModel struct {
	core.QAbstractItemModel

	_ func() `constructor:"init"`

	// _ func()                                  `signal:"remove,auto"`
	// _ func(item []*core.QVariant)             `signal:"add,auto"`
	// _ func(firstName string, lastName string) `signal:"edit,auto"`

	rootItem service // ROOT
}

func (m *CustomTreeModel) init() {
	m.rootItem = &ROOT{
		Collection{
			title: "FirstName LastName",
			_childItems: []CollectionItem{
				CollectionItem{
					_itemData: []string{"john", "doe"},
				},
			},
		},
	}

	m.rootItem.add(Collection{
		title: "FirstName LastName2",
		_childItems: []CollectionItem{
			CollectionItem{
				_itemData: []string{"john", "bob"},
			},
			CollectionItem{
				_itemData: []string{"jim", "bob"},
			},
			CollectionItem{
				_itemData: []string{"jimmy", "bob"},
			},
		},
	})

	m.ConnectIndex(m.index)
	m.ConnectParent(m.parent)
	m.ConnectRowCount(m.rowCount)
	m.ConnectColumnCount(m.columnCount)
	m.ConnectData(m.data)

	m.ConnectRoleNames(m.roleNames)
}

func (m *CustomTreeModel) index(row int, column int, parent *core.QModelIndex) *core.QModelIndex {
	if !m.HasIndex(row, column, parent) {
		return core.NewQModelIndex()
	}

	if !parent.IsValid() {
		return m.CreateIndex2(row, column, something)
	} else {
		return m.CreateIndex2(row, column, uintptr(parent.Row()))
	}

	return core.NewQModelIndex()
}

func (m *CustomTreeModel) parent(index *core.QModelIndex) *core.QModelIndex {
	if !index.IsValid() {
		return core.NewQModelIndex()
	}
	id := int(index.InternalId())

	if id == something {
		return core.NewQModelIndex()
	}

	return m.CreateIndex2(id, 0, something)
}

func (m *CustomTreeModel) roleNames() map[int]*core.QByteArray {
	return map[int]*core.QByteArray{
		FirstName: core.NewQByteArray2("FirstName", -1),
		LastName:  core.NewQByteArray2("LastName", -1),
	}
}

func (m *CustomTreeModel) rowCount(parent *core.QModelIndex) int {
	if !parent.IsValid() {
		return m.rootItem.Len()
	}
	parentId := int32(parent.InternalId())

	if parentId == something {
		return m.rootItem.CollectionLen(parent.Row())
	}

	return 0
}

func (m *CustomTreeModel) columnCount(parent *core.QModelIndex) int {
	if !parent.IsValid() {
		return 1
	}
	parentId := int32(parent.InternalId())

	if parentId == something {
		return 1 //len(r[parent.Row()].title)
	}

	return 0
}

func (m *CustomTreeModel) data(index *core.QModelIndex, role int) *core.QVariant {
	if !index.IsValid() {
		return core.NewQVariant()
	}

	if int32(index.InternalId()) == something {
		return core.NewQVariant17(m.rootItem.Collection(index.Row()).title)
	}
	switch role {
	case FirstName:
		return core.NewQVariant17(m.rootItem.CollectionItem(int(index.InternalId()), index.Row())._itemData[0])
	case LastName:
		return core.NewQVariant17(m.rootItem.CollectionItem(int(index.InternalId()), index.Row())._itemData[1])
	}
	return core.NewQVariant()
}

// func (m *CustomTreeModel) remove() {
// 	if m.rootItem.childCount() == 0 {
// 		return
// 	}
// 	m.BeginRemoveRows(core.NewQModelIndex(), len(m.rootItem._childItems)-1, len(m.rootItem._childItems)-1)
// 	m.rootItem._childItems = m.rootItem._childItems[:len(m.rootItem._childItems)-1]
// 	m.EndRemoveRows()
// }

// func (m *CustomTreeModel) add(item []*core.QVariant) {
// 	m.BeginInsertRows(core.NewQModelIndex(), len(m.rootItem._childItems), len(m.rootItem._childItems))
// 	m.rootItem.appendChild(NewTreeItem([]string{item[0].ToString(), item[1].ToString()}))
// 	m.EndInsertRows()
// }

// func (m *CustomTreeModel) edit(firstName string, lastName string) {
// 	if m.rootItem.childCount() == 0 {
// 		return
// 	}
// 	m.BeginRemoveRows(core.NewQModelIndex(), len(m.rootItem._childItems)-1, len(m.rootItem._childItems)-1)
// 	m.BeginInsertRows(core.NewQModelIndex(), len(m.rootItem._childItems)-1, len(m.rootItem._childItems)-1)
// 	item := m.rootItem._childItems[len(m.rootItem._childItems)-1]
// 	item._itemData = []string{firstName, lastName}
// 	m.EndRemoveRows()
// 	m.EndInsertRows()

// 	//TODO:
// 	//ideally DataChanged should be used instead, but it doesn't seem to work ...
// 	//if you search for "qml treeview datachanged" online
// 	//it will just lead you to tons of unresolved issues
// 	//m.DataChanged(m.Index(item.row(), 0, core.NewQModelIndex()), m.Index(item.row(), 1, core.NewQModelIndex()), []int{FirstName, LastName})
// 	//feel free to send a PR, if you got it working somehow :)
// }

func main() {
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)
	if !core.QResource_RegisterResource("qml.rcc", "") {
		panic("failure: resource needed")
	}
	// gui.QGuiApplication_Screens()[0].
	app := widgets.NewQApplication(len(os.Args), os.Args)

	view := quick.NewQQuickView(nil)
	view.SetTitle("treeview Example")
	view.SetResizeMode(quick.QQuickView__SizeRootObjectToView)
	view.SetSource(core.NewQUrl3("qrc:/qml/main.qml", 0))
	// view.SetPosition2(posx, posy)
	view.ShowMaximized()

	app.Exec()
}
