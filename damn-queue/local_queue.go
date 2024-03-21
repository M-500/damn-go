package damn_queue

import (
	"context"
	"errors"
	"sync"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-21 22:04

// LocalQueue
// @Description: 底层数据存取采用双向链表，实现
type LocalQueue struct {
	data   *LinkNode
	count  int // 当前能存多少元素
	cap    int // 一共可以存多少元素
	header *LinkNode
	Tail   *LinkNode
	lock   *sync.RWMutex
}

func (l *LocalQueue) EnQueue(ctx context.Context, data any) error {
	if l.count > l.cap {
		return errors.New("容量不够！")
	}
	l.lock.Lock()
	defer l.lock.Unlock()
	node := &LinkNode{
		prev: l.Tail,
		next: nil,
		data: data,
	}
	l.Tail.next = node
	l.Tail = node
	return nil
}

func (l *LocalQueue) DeQueue(ctx context.Context) (any, error) {
	l.lock.Lock()
	defer l.lock.Unlock()
	if l.count <= 0 || l.header == nil {
		return nil, errors.New("没有数据")
	}
	// 从头部区
	temp := l.header
	l.header = l.header.next
	return temp.data, nil
}

func (l *LocalQueue) IsEmpty() (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (l *LocalQueue) IsFull() (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (l *LocalQueue) Len() (uint64, error) {
	//TODO implement me
	panic("implement me")
}
