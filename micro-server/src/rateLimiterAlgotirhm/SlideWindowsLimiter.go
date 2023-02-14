package rateLimiterAlgotirhm

import (
	"sync"
	"sync/atomic"
	"time"
)

const SmallWindows int64 = 1 << 5
const TimeShift = 1e9
const WindowsSize int64 = 3 * TimeShift
const MaxRequestPerWindows = WindowsSize / TimeShift * 10

// const TimeShift = 1000000000

//const MaxRequestPerWindows = WindowsSize / TimeShift * 10

type slideWindowsLimiter struct {
	permitsPerWindows    int64
	windows              map[int64]int64
	totalCount           int64
	lock                 sync.Mutex
	once                 sync.Once
	timestamp            int64
	smallWindowsDistance int64
	windowsSize          int64
	clearFlag            int32
	cond                 *sync.Cond
}

func init() {
	slideLimiter = &slideWindowsLimiter{
		permitsPerWindows: MaxRequestPerWindows,
		// windows length is  prime number may be can defeat conflict better
		windows:              make(map[int64]int64, SmallWindows+3),
		timestamp:            time.Now().UnixNano(),
		lock:                 sync.Mutex{},
		smallWindowsDistance: WindowsSize / SmallWindows,
		windowsSize:          WindowsSize,
	}
	slideLimiter.cond = sync.NewCond(&slideLimiter.lock)
	slideLimiter.initMapping()

}

func (s *slideWindowsLimiter) initMapping() {
	// init mapping
	s.once.Do(func() {
		var i int64 = 0
		for ; i < s.windowsSize; i++ {
			s.windows[i] = 0
		}
	})
}

var slideLimiter *slideWindowsLimiter

func TryAcquire() bool {
	return slideLimiter.TryAcquire()
}
func (s *slideWindowsLimiter) TryAcquire() bool {
	var diff int64
	s.lock.Lock()
	//for atomic.LoadInt32(&s.clearFlag) != 0 {
	//	s.cond.Wait()
	//}
	diff = time.Now().UnixNano() - s.timestamp
	var index = diff / s.smallWindowsDistance
	if diff <= s.windowsSize {
		if s.totalCount < s.permitsPerWindows {
			s.totalCount++
			s.windows[index]++
			s.lock.Unlock()
			return true
		} else {
			s.lock.Unlock()
			return false
		}
	} else {
		if atomic.CompareAndSwapInt32(&s.clearFlag, 0, 1) {
			go func() {
				s.timestamp += diff
				var i int64 = 0
				for ; i <= index; i++ {
					s.totalCount -= s.windows[i]
					s.windows[i] = 0
				}
				println(atomic.CompareAndSwapInt32(&s.clearFlag, 1, 0))
				//s.cond.Broadcast()
			}()
		}
		s.lock.Unlock()
	}
	return s.TryAcquire()
}
