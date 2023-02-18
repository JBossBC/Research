package rateLimiter

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestSlideWindows(t *testing.T) {
	println(TryAcquire())
}
func BenchmarkTryAcquire(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TryAcquire()
	}
}

func TestConcurrencyTryAcquire(t *testing.T) {
	const times = 10000
	group := sync.WaitGroup{}
	group.Add(times)

	var smoothTimes time.Duration = 3 * time.Millisecond
	var resultInt int64
	var cycle int64 = 0
	for i := 0; i < times; i++ {
		go func() {
			defer group.Done()
			time.Sleep(smoothTimes)
			if cycle >= 100 && smoothTimes < 3*time.Second {
				smoothTimes += smoothTimes * 10
				cycle = 0
			} else {
				atomic.AddInt64(&cycle, 1)
			}
			result := TryAcquire()
			if slideLimiter.totalCount > slideLimiter.permitsPerWindows {
				panic(any("slide windows invalid"))
			}
			if result {
				atomic.AddInt64(&resultInt, 1)
			}

		}()
	}
	group.Wait()
	println(resultInt)
	for key, value := range slideLimiter.windows {
		println(key, ":", value)
	}
}
