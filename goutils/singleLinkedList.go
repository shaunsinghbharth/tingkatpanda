package goutils

import (
	"errors"
	"fmt"
)

type strNode struct {
	item string
	next *strNode
}

type singleLinkedList struct {
	head *strNode
	size int
}

func GetList() *singleLinkedList {
	return &singleLinkedList{nil, 0}
}

func (p *singleLinkedList) AddNode(name string) error {
	newNode := &strNode{
		item: name,
		next: nil,
	}
	if p.head == nil {
		p.head = newNode
	} else {
		currentNode := p.head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = newNode
	}
	p.size++
	return nil
}

func (p *singleLinkedList) PrintAllNodes() error {
	currentNode := p.head
	if currentNode == nil {
		fmt.Println("Linked list is empty.")
		return nil
	}
	fmt.Printf("%+v\n", currentNode.item)
	for currentNode.next != nil {
		currentNode = currentNode.next
		fmt.Printf("%+v\n", currentNode.item)
	}
	return nil
}

func (p *singleLinkedList) Remove(index int) error {
	currentNode := p.head
	for i := 0; i < index; i++ {
		if i == index-1 {
			currentNode.next = currentNode.next.next
			return nil
		}
		currentNode = currentNode.next
	}
	return errors.New("node exists")
}

func (p *singleLinkedList) AddAtPos(index int, s string) error {
	currentNode := p.head
	temp := &strNode{s, nil}

	for i := 0; i <= index; i++ {
		currentNode = currentNode.next
		var holder *strNode = nil
		if currentNode.next.next != nil {
			holder = currentNode.next.next
		}

		if i == index-2 {
			currentNode.next = temp
		}

		if i == index-1 {
			currentNode.next = holder

			return nil
		}
	}
	return errors.New("could not add")
}

func (p *singleLinkedList) Get(index int) string {
	currentNode := p.head
	for i := 0; i < index; i++ {
		if i == index-1 {
			return currentNode.item
		}
		currentNode = currentNode.next
	}
	return ""
}
