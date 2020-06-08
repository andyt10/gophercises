package src

var store []ListItem

func GetAll() []ListItem {
	return store
}

func Add(itemData string) int, error {
	store.append(ListItem{item: itemData})
}

func Remove(itemIndex int) {
    store[itemIndex] = store[len(store)-1]
    store[len(store)-1] = nil
    store = store[:len(store)-1]
}
