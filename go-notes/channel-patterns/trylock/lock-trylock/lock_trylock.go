package lock_trylock

import (
	"sync"
	"sync/atomic"
	"unsafe"
)

/*
我们知道， Go的标准库sync有Mutex,可以用来作为锁，但是Mutex却没有实现TryLock方法。

我们对于TryLock的定义是当前goroutine尝试获得锁， 如果成功，则获得了锁，返回true, 否则返回false。
我们可以使用这个方法避免在获取锁的时候当前goroutine被阻塞住。

本来，这是一个常用的功能，在一些其它编程语言中都有实现，为什么Go中没有实现的？issue#6123有详细的讨论，
在我看来，Go核心组成员本身对这个特性没有积极性，并且认为通过channel可以实现相同的方式。

*/

const mutexLocked  =  1 << iota

type Mutex struct {
	mu sync.Mutex
}

func (m *Mutex)Lock()  {
	m.mu.Lock()
}

func (m *Mutex)Unlock()  {
	m.mu.Unlock()
}

func (m *Mutex)TryLock() bool  {
	return atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&m.mu)), 0, mutexLocked)
}

func (m *Mutex)Islocked() bool {
	return atomic.LoadInt32((*int32)(unsafe.Pointer(&m.mu))) == mutexLocked
}