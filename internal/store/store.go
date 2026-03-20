package store

import (
	"sync"
)

// TODO: 实现内存 KV 存储
//
// 你需要做的：
// 1. 定义一个 Store struct（用 map + sync.RWMutex）
// 2. 实现 Get(key) / Put(key, value) / Delete(key) / List() 方法
// 3. 所有方法必须并发安全（用读写锁保护 map）
//
// 挑战（选做）：
//   - 支持 TTL：Put 时设置过期时间，Get 时检查是否过期
//   - 提示：可以用 time.Time 记录过期时间，开一个 goroutine 定期清理
//
// 需要的包：
//   - "sync"
//   - "time"（如果做 TTL）

type Store struct {
	mu   sync.RWMutex
	data map[string]string
}

func New() *Store {
	return &Store{
		data: make(map[string]string),
	}
}

func (s *Store) Get(id string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	val, ok := s.data[id]
	return val, ok
}

func (s *Store) Put(id string, val string) string {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[id] = val
	return val
}

func (s *Store) Delete(id string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, ok := s.data[id]
	if !ok {
		return false
	}
	delete(s.data, id)
	return true
}

func (s *Store) List() map[string]string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	res := make(map[string]string, len(s.data))
	for k, v := range s.data {
		res[k] = v
	}
	return res
}
