package store_test

// TODO: 为你的 Store 写单元测试
//
// 测试用例建议：
// 1. TestPutAndGet — 写入后能读到
// 2. TestGetNotFound — 读不存在的 key
// 3. TestDelete — 删除后读不到
// 4. TestList — 写入多个 key 后全部列出
// 5. TestConcurrency — 多个 goroutine 同时读写不 panic
// 6. TestTTL（选做）— 写入带过期时间的 key，过期后读不到
//
// 需要的包：
//   - "testing"
//   - "sync"（并发测试）
