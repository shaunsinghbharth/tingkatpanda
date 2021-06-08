package goutils

type PriorityNode struct {
	Item     string
	next     *PriorityNode
	prev     *PriorityNode
	priority int
}

type PriorityQueue struct {
	Front *PriorityNode
	Back  *PriorityNode
	Size  int
}

func (p *PriorityQueue) Enqueue(name string, priority int) error {
	temp := p.Front

	newNode := &PriorityNode{
		Item:     name,
		next:     nil,
		prev:     nil,
		priority: priority,
	}

	for i := 10; i > 0 && temp != nil; i-- {
		for j := 0; j < p.Size; j++ {
			if priority < i {
				temp.next = newNode
				p.Size++

				return nil
			}
			temp = temp.next
		}
	}

	if p.Front == nil {
		p.Front = newNode
	} else {
		p.Back.next = newNode
	}
	p.Back = newNode
	p.Size++

	return nil
}

func (p *PriorityQueue) Dequeue() string {
	var item string
	if p.Front == nil {
		return ""
	}

	item = p.Front.Item
	if p.Size == 1 {
		p.Front = nil
		p.Back = nil
	} else {
		p.Front = p.Front.next
	}
	p.Size--

	return item
}
