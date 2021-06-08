package goutils

type Stack struct {
	Top  *Node
	Size int
}

func (p *Stack) Push(name string) error {
	newNode := &Node{
		Item: name,
		next: nil,
	}
	if p.Top == nil {
		p.Top = newNode
	} else {
		newNode.next = p.Top
		p.Top = newNode
	}
	p.Size++
	return nil
}

func (p *Stack) Pop() string {
	var item string
	if p.Top == nil {
		return ""
	}
	item = p.Top.Item

	if p.Size == 1 {
		p.Top = nil
	} else {
		p.Top = p.Top.next
	}
	p.Size--
	return item
}
