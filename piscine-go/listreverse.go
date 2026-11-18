package piscine

func ListReverse(l *List) {
	if l == nil || l.Head == nil {
		return
	}

	var prev *NodeL
	current := l.Head

	l.Tail = l.Head // old head becomes new tail

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	l.Head = prev // prev becomes the new head
}
