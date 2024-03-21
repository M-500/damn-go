package damn_queue

import (
	"context"
	"sync"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-21 22:04

type LocalQueue struct {
	data  []any
	count int
	lock  *sync.RWMutex
}

func (l *LocalQueue) EnQueue(ctx context.Context, data any) error {
	//TODO implement me
	panic("implement me")
}

func (l *LocalQueue) DeQueue(ctx context.Context) (any, error) {
	//TODO implement me
	panic("implement me")
}

func (l *LocalQueue) IsEmpty() (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (l *LocalQueue) IsFull() (bool, error) {
	//TODO implement me
	panic("implement me")
}

func Len() (uint64, error) {
	//TODO implement me
	panic("implement me")
}
