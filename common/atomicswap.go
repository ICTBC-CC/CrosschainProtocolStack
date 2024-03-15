// coding:utf-8
// 原子跨链资源结构
package common

// 资源条目
type Item struct {
	ChainName    string // 链名
	ResourceName string // 资源名
	Count        int    // 交换的资源数量
}

// 跨链原子交换结构
type AtomicSwap struct {
	HashLock     []byte // 链下用户协调商量的哈希锁
	TimeLock     int64  // 时间锁,秒,到这个时间就触发回退
	LockResource []Item // 自己锁定的资源
	WantResource []Item // 自己想要交换的资源
}
