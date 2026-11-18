package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	if l == nil {
		return nil
	}

	changed := true
	for changed {
		changed = false
		current := l
		for current.Next != nil {
			if current.Data > current.Next.Data {
				current.Data, current.Next.Data = current.Next.Data, current.Data
				changed = true
			}
			current = current.Next
		}
	}

	return l
}
