package damn_queue

import "context"

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-21 22:00

type Queue[T any] interface {
	EnQueue(ctx context.Context, data T) error
	DeQueue(ctx context.Context) (T, error)
	IsEmpty() (bool, error)
	IsFull() (bool, error)
	Len() (uint64, error)
}
