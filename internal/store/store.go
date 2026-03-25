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
	mu          sync.RWMutex
	data        map[string]string
	subscribers map[string][]chan Event // TODO: key -> 订阅者列表
}

// Event 代表一次数据变化
type Event struct {
	// TODO: 定义字段 — 需要哪些信息来描述"发生了什么"？
	// 提示：Id, Value, Action (比如 "PUT" 或 "DELETE")
	id     string
	val    string
	action string
}

func New() *Store {
	return &Store{
		data:        make(map[string]string),
		subscribers: make(map[string][]chan Event),
	}
}

// Subscribe 注册一个订阅者，返回一个 channel 用来接收事件
// TODO: 你来实现
//  1. 创建一个 chan Event（带 buffer，比如 16）
//  2. 把它 append 到 s.subscribers[id]
//  3. 返回这个 channel
//  4. 需要加锁吗？用什么锁？
func (s *Store) Subscribe(id string) chan Event {
	s.mu.Lock()
	defer s.mu.Unlock()
	ch := make(chan Event, 16)
	s.subscribers[id] = append(s.subscribers[id], ch)
	return nil
}

// Unsubscribe 移除一个订阅者并关闭 channel
// TODO: 你来实现
//  1. 遍历 s.subscribers[id]，找到这个 ch
//  2. 从 slice 中移除它
//  3. close(ch)
//  4. 需要加锁
func (s *Store) Unsubscribe(id string, ch chan Event) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, sub := range s.subscribers[id] {
		if sub == ch {
			s.subscribers[id] = append(s.subscribers[id][:i], s.subscribers[id][i+1:]...)
			break
		}
	}
	close(ch)
}

// notify 通知某个 key 的所有订阅者（内部方法，在 Put/Delete 里调用）
// TODO: 你来实现
//  1. 遍历 s.subscribers[id]
//  2. 用 conc pool 并发地往每个 channel 发 Event
//  3. 调用时已经持有写锁，所以这里不要再加锁
//
// 需要的包：
//   - "github.com/sourcegraph/conc/pool"
func (s *Store) notify(id string, event Event) {
	// TODO: 实现
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
	// TODO: 调用 s.notify(id, Event{...}) 通知订阅者
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
	// TODO: 调用 s.notify(id, Event{...}) 通知订阅者
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
