# sync
https://golang.google.cn/pkg/sync/


## sync.Pool

[type Pool](https://golang.google.cn/pkg/sync/#Pool)

- 临时对象
- 自动移除
- 当这个对象的引用只有sync.Pool持有时，这个对象内存会被释放
- goroutine安全
- 目的就是缓存并重用对象，减少GC的压力
- 自动扩容、缩容
- 不要去拷贝pool，也就是说最好单例

### 什么情况下适合使用sync.Pool呢？
### sync.Pool的对象什么时候会被回收呢？
### sync.Pool是如何实现线程安全的？