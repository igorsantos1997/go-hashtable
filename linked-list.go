package main

type LinkedList[T any] struct {
	Current T
	Next    *LinkedList[T]
}

func (ll *LinkedList[T]) Insert(value T) {
	if ll.Next != nil {
		ll.Next.Insert(value)
		return
	}

	ll.Next = &LinkedList[T]{
		Current: value,
		Next:    nil,
	}
}

// Search busca um elemento na lista
func (ll *LinkedList[T]) Search(predicate func(T) bool) *LinkedList[T] {
	if predicate(ll.Current) {
		return ll
	}

	if ll.Next == nil {
		return nil
	}

	return ll.Next.Search(predicate)
}

// Delete remove um elemento da lista
func (ll *LinkedList[T]) Delete(predicate func(T) bool) bool {
	if ll.Next == nil {
		return false
	}

	if predicate(ll.Next.Current) {
		ll.Next = ll.Next.Next
		return true
	}

	return ll.Next.Delete(predicate)
}

// Size retorna o tamanho da lista
func (ll *LinkedList[T]) Size() int {
	if ll.Next == nil {
		return 1
	}
	return 1 + ll.Next.Size()
}
