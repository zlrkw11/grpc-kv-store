package store

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
