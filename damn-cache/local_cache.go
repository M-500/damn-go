//@Author: wulinlin
//@Description:
//@File:  local_cache
//@Version: 1.0.0
//@Date: 2024/03/21 20:57

package damn_cache

import (
	"context"
	"errors"
	"sync"
	"time"
)

type localCache struct {
	data map[string]any
	lock *sync.RWMutex
}

func (l *localCache) Get(ctx context.Context, key string) (any, error) {
	l.lock.RLock()
	defer l.lock.RLock()
	val, ok := l.data[key]
	if !ok {
		return nil, errors.New("Key Not Found")
	}
	return val, nil
}

func (l *localCache) Set(ctx context.Context, key string, value any, expire time.Duration) error {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.data[key] = value
	return nil
}

func (l *localCache) Delete(ctx context.Context, key string) error {
	l.lock.Lock()
	defer l.lock.Unlock()
	delete(l.data, key)
	return nil
}

func (l *localCache) Clear(ctx context.Context) {
	//TODO implement me
	panic("implement me")
}

func NewLocalCache(firstSize uint32) Cache {
	return &localCache{
		data: make(map[string]any, firstSize),
	}
}
