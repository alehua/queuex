package priority_queue

// 基于container/heap实现的优先级队列

// HeapPriorityQueue to implements the interface of "heap.Interface"
type HeapPriorityQueue[T any] struct {
	Cap  int
	Data []T
	Cmp  Compare[T]
}

func NewHeapPriorityQueue[T any](len int, cmp Compare[T]) *HeapPriorityQueue[T] {
	return &HeapPriorityQueue[T]{
		Cmp:  cmp,
		Data: make([]T, 0, len),
		Cap:  len,
	}
}

// 实现container/heap接口

// Len 队列长度
func (h *HeapPriorityQueue[T]) Len() int { return len(h.Data) }

func (h *HeapPriorityQueue[T]) Less(i, j int) bool {
	v := h.Cmp(h.Data[i], h.Data[j])
	return v < 0
}

func (h *HeapPriorityQueue[T]) Swap(i, j int) { h.Data[i], h.Data[j] = h.Data[j], h.Data[i] }

func (h *HeapPriorityQueue[T]) Push(x T) {
	v := append(h.Data, x)
	h.Data = v
}

func (h *HeapPriorityQueue[T]) Top() (T, error) {
	n := len(h.Data)
	if n <= 0 {
		return *new(T), ErrEmptyQueue
	}
	return h.Data[n-1], nil
}

func (h *HeapPriorityQueue[T]) Pop() T {
	old := h.Data
	n := len(old)
	x := old[n-1]
	h.Data = old[0 : n-1]
	return x
}

func (h *HeapPriorityQueue[T]) Peek() (T, error) {
	old := h.Data
	n := len(old)
	if n <= 0 {
		return *new(T), ErrEmptyQueue
	}
	x := old[n-1]
	h.Data = old[0 : n-1]
	return x, nil
}

func (h *HeapPriorityQueue[T]) Enqueue(x T) error {
	if len(h.Data) >= h.Cap {
		return ErrOutOfCapacity
	}
	h.Push(x)
	return nil
}

func (h *HeapPriorityQueue[T]) Dequeue() (T, error) {
	return h.Peek()
}
