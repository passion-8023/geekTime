package relock

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

type RecursiveMutex struct {
	sync.Mutex
	owner int64 //当前持有锁的goroutine id
	recursive int32 //这个goroutine 重入的次数
}

func (m *RecursiveMutex) Lock()  {
	gid := GoID()
	if atomic.LoadInt64(&m.owner) == gid {
		m.recursive++
		return
	}
	m.Mutex.Lock()
	atomic.StoreInt64(&m.owner, gid)
	m.recursive = 1
}

func (m *RecursiveMutex) Unlock()  {
	gid := GoID()
	if atomic.LoadInt64(&m.owner) != gid {
		panic(fmt.Sprintf("wrong the owner(%d): %d!", m.owner, gid))
	}
	m.recursive--
	if m.recursive != 0 {
		return
	}
	atomic.StoreInt64(&m.owner, -1)
	m.Mutex.Unlock()
}

func GoID() int64 {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return int64(id)
}


