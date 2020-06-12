package src

type ListItemEntry struct {
	Index int
	Data  ListItem
}

type ListItem struct {
	Item   string
	Added  int
	DoneAt int
}
