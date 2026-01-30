package lister

import (
	"bytes"
	"fmt"
	"os"
)

type Lister struct {
	head *ListNode
	tail *ListNode
	size int
}

type ListNode struct {
	prev *ListNode
	next *ListNode
	data *string
}

func (ln *ListNode) Data() string {
	return *ln.data
}

func (l *Lister) Append(data string) (ListNode, error) {
	l.size++
	if l.head == nil {
		l.head = &ListNode{
			data: &data,
		}
		l.tail = l.head
		return *l.head, nil
	} else if l.tail == l.head {
		l.tail = &ListNode{
			prev: l.head,
			data: &data,
		}
		l.head.next = l.tail

		return *l.tail, nil
	} else {
		n := &ListNode{
			prev: l.tail,
			next: nil,
			data: &data,
		}
		l.tail.next = n
		l.tail = n

		return *n, nil
	}
}

func (l *Lister) Size() int {
	return l.size
}

func (l *Lister) Show() {
	buff := bytes.NewBufferString("")

	for curr := l.head; curr != nil; curr = curr.next {
		switch curr {
		case l.head:
			buff.WriteString(*curr.data)
			if l.tail == l.head {
				fmt.Fprintf(buff, ", Size: %d", l.size)
			}
		case l.tail:
			fmt.Fprintf(buff, " -> %s, Size: %d", *curr.data, l.size)
		default:
			fmt.Fprintf(buff, " -> %s", *curr.data)
		}
	}

	if buff.Len() == 0 {
		fmt.Printf("<- Empty List ->, Size: %d\n", l.size)
	} else {
		fmt.Fprintln(os.Stdout, buff)
	}

}

func NewList() Lister {
	return Lister{}
}
