package list

type Node struct {
	value int
	next  *Node
}

type List struct {
	head *Node
}

func New() *List {
	return &List{}
}

func (l *List) AddToStart(value int) {
	if l.head == nil {
		l.head = &Node{value: value}
		return
	}

	n := &Node{value: value}
	n.next = l.head
	l.head = n
}

func (l *List) isEmpty() bool {
	return l.head == nil
}

// Head return head value and true
// or 0 and false if list is empty.
func (l *List) Head() (int, bool) {
	if l.head != nil {
		return l.head.value, true
	}

	return 0, false
}

func (l *List) GetAll() []int {
	values := make([]int, 0)
	for cur := l.head; cur != nil; cur = cur.next {
		values = append(values, cur.value)
	}

	return values
}

// Reverse returns new list with reversed values.
// Input:  1 -> 2 -> 3 -> 4 -> 5
// Output: 5 -> 4 -> 3 -> 2 -> 1
func (l *List) Reverse() *List {
	rl := New()

	if l.isEmpty() {
		return rl
	}

	rl.head = &Node{value: l.head.value}

	for cur := l.head.next; cur != nil; cur = cur.next {
		n := &Node{value: cur.value}
		n.next = rl.head
		rl.head = n
	}

	return rl
}
