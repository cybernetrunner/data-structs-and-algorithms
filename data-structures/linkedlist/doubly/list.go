package doubly

type node struct {
	data int
	prev *node
	next *node
}

type LinkedList struct {
	Head   *node
	Tail   *node
	Length int
}

func (ll *LinkedList) AddFirst(data int) {
	temp := ll.Head
	n := &node{
		data: data,
		next: temp,
	}

	temp.prev = n
	temp.next.prev = n
	ll.Head = n
}

func (ll *LinkedList) AddLast(d int) {
	n := &node{data: d, next: nil}
	ll.Length++

	if ll.Head == nil {
		ll.Head, ll.Tail = n, n
		return
	}

	ll.Tail.next, n.prev, ll.Tail = n, ll.Tail, n
}

func (ll *LinkedList) RemoveFirst(d int) (ok bool) {
	if ll.Head.data == d {
		ll.Head = ll.Head.next
		ll.Length--

		return true
	}

	return ll.removeFirst(ll.Head.next, d)
}

func (ll *LinkedList) RemoveLast(d int) (ok bool) {
	if ll.Tail.data == d {
		ll.Tail = ll.Tail.prev
		ll.Length--

		return true
	}

	return ll.removeLast(ll.Tail.prev, d)
}

func (ll *LinkedList) Contains(d int) (ok bool) {
	if ll.Head.data == d {
		return true
	}

	return ll.findFirst(ll.Head.next, d)
}

func (ll *LinkedList) Clear() {
	ll.Head = nil
	ll.Tail = nil
	ll.Length = 0
}

func (ll *LinkedList) CopyToSlice() []int {
	arr := make([]int, ll.Length)
	currentNode := ll.Head

	for i := range arr {
		arr[i] = currentNode.data

		if currentNode.next == nil {
			break
		}

		currentNode = currentNode.next
	}

	return arr
}

func (ll *LinkedList) removeFirst(n *node, d int) (ok bool) {
	if ll.Tail == n {
		return false
	} else if n.next.data == d {
		n.next, n.next.next.prev = n.next.next, n
		ll.Length--

		return true
	}

	return ll.removeFirst(n.next, d)
}

func (ll *LinkedList) removeLast(n *node, d int) (ok bool) {
	if ll.Tail == n {
		return false
	} else if n.prev.data == d {
		n.prev = n.prev.next
		ll.Length--

		return true
	}

	return ll.removeLast(n.prev, d)
}

func (ll *LinkedList) findFirst(n *node, d int) (ok bool) {
	if n.data == d {
		return true
	}

	return ll.findFirst(n.next, d)
}
