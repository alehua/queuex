package priority_queue

import "errors"

// Compare 可比较类型
type Compare[T any] func(src T, dst T) int

var (
	ErrOutOfCapacity = errors.New("queue: 超出最大容量限制")
	ErrEmptyQueue    = errors.New("queue: 队列为空")
)

type IPriorityQueue[T any] interface {
	Len() int            // 队列长度
	Top() (T, error)     // 获取队头
	Peek() (T, error)    // 弹出队头
	Enqueue(t T) error   // 入队
	Dequeue() (T, error) // 出队
}
