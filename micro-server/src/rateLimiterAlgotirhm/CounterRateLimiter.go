package rateLimiterAlgotirhm

//
//import (
//	"sync"
//	"time"
//)
//
//const TimeShift = 1e9
//const WindowsSize int64 = 3 * TimeShift
//const MaxRequestPerWindows = WindowsSize / TimeShift * 10

//
//type counterRateLimiter struct {
//	permitsPerWindows int64
//	Timestamp         int64
//	counter           int64
//	lock              sync.Mutex
//}
//
//type Options func(limiter *counterRateLimiter)
//
//var counterLimiter *counterRateLimiter = &counterRateLimiter{
//	permitsPerWindows: MaxRequestPerWindows,
//	Timestamp:         time.Now().UnixNano(),
//	lock:              sync.Mutex{},
//}
//
//func TryAcquire() bool {
//	return counterLimiter.TryAcquire()
//}
//func GetCounterLimiter(options ...Options) *counterRateLimiter {
//	var result = &counterRateLimiter{
//		permitsPerWindows: MaxRequestPerWindows,
//		Timestamp:         time.Now().Unix(),
//		lock:              sync.Mutex{},
//	}
//	for i := 0; i < len(options); i++ {
//		options[i](result)
//	}
//	return result
//}
//
//func (c *counterRateLimiter) TryAcquire() bool {
//	var now = time.Now().UnixNano()
//	if now-c.Timestamp < WindowsSize {
//		c.lock.Lock()
//		if c.counter < c.permitsPerWindows {
//			c.counter++
//			c.lock.Unlock()
//			return true
//		} else {
//			c.lock.Unlock()
//			return false
//		}
//	}
//	c.Timestamp = now
//	c.counter = 0
//	return c.TryAcquire()
//}
