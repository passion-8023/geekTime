package trylock

import (
	"sync"
	"sync/atomic"
	"unsafe"
)

const (
	mutexLocked = 1 << iota // mutex is locked
	mutexWoken
	mutexStarving
	mutexWaiterShift = iota
)

type Mutex struct {
	sync.Mutex
}

func (m *Mutex) TryLock() bool {
	//如果能够成功抢到锁
	if atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&m.Mutex)), 0, mutexLocked) {
		return true
	}

	//如果处于唤醒、饥饿或者加锁状态，这次请求就不参与竞争了
	old := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	if old&(mutexLocked|mutexWoken|mutexStarving) != 0 {
		return false
	}

	//尝试在竞争的状态下请求锁
	new := old|mutexLocked
	return atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&m.Mutex)), old, new)

}