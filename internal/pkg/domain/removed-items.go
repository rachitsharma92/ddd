package domain

type RemovedItems struct{
	Name []string
}

func (removedItems *RemovedItems) Add (removedItem string) {
	removedItems.Name = append(removedItems.Name, removedItem)
}