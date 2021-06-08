package goutils

type QueueNode struct {
	Item string
	next *QueueNode
}

type Queue struct {
	Front *QueueNode
	Back  *QueueNode
	Size  int
}

func (p *Queue) Enqueue(name string) error {
	newNode := &QueueNode{
		Item: name,
		next: nil,
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

func (p *Queue) Dequeue() string {
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
