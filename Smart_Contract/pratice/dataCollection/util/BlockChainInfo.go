//package util
//
//import (
//	"fmt"
//	"github.com/cockroachdb/errors"
//	"github.com/ethereum/go-ethereum"
//	"github.com/ethereum/go-ethereum/common"
//	"github.com/ethereum/go-ethereum/core/types"
//	"golang.org/x/net/context"
//	"math/big"
//	"strings"
//	"sync"
//	"sync/atomic"
//	"time"
//)
//
//type workState int32
//
//var (
//	timeout workState = 3
//	success workState = 1
//	failed  workState = 2
//	ongoing workState = 0
//)
//
//var TimeLess time.Duration = 1<<63 - 1
//
//const MaxQueryBlockSize int64 = 2000
//
//const defaultMallocCap int64 = 1024
//const MaxConcurrentNumber = 8e4
//const MaxWorkNumber = MaxConcurrentNumber / MaxQueryBlockSize
//const EmergencyRecovery = 100
//const SmoothRecoverRatio = 0.25
//
//var DefaultSmoothRecoverTimes = time.Second
//
//func GetCurrentBlockNumber() (uint64, error) {
//	return GetClient().BlockNumber(context.Background())
//}
//
//func GetEvent(timeout time.Duration, from int64, to int64, address []common.Address, topics [][]common.Hash) (stream *LogsStream, err error) {
//	info := newGlobalInfo(timeout, from, to, address, topics)
//	var workNumber = info.workNumber
//	var i int32 = 0
//	for ; i < workNumber; i++ {
//		newLogsWork(info).handler()
//	}
//	info.group.Wait()
//	ok := atomic.CompareAndSwapInt32((*int32)(&info.state), 0, 1)
//	if !ok {
//		return nil, fmt.Errorf("get event error: %v", info.err)
//	}
//	logs := info.arrangeLogs()
//	stream = &LogsStream{
//		logs: logs,
//	}
//	return stream, nil
//}
//func (g *globalInfo) arrangeLogs() []types.Log {
//	var i int32 = 0
//	var result = make([]types.Log, 0, defaultMallocCap)
//	for ; i < g.currentId; i++ {
//		result = append(result, g.queue[i].returnValue...)
//	}
//	return result
//}
//
//func newGlobalInfo(timeout time.Duration, from int64, to int64, address []common.Address, topics [][]common.Hash) (g *globalInfo) {
//	var workNumber = (to - from) / MaxQueryBlockSize
//	if MaxQueryBlockSize*workNumber+from != to {
//		workNumber++
//	}
//	g = &globalInfo{end: to, errTrigger: sync.Once{}, mutex: sync.Mutex{}, workNumber: int32(workNumber), address: address, topics: topics, offset: from, timeout: timeout, queue: make([]*logsWork, workNumber), group: sync.WaitGroup{}}
//	var chanNumber int64 = workNumber
//	if chanNumber > MaxWorkNumber {
//		chanNumber = MaxWorkNumber
//	}
//	g.workMutex = sync.Mutex{}
//	g.controlPanel = controlPanel{cond: sync.NewCond(&g.workMutex), recoverSignal: make(chan int32, 1)}
//	g.workChan = make(chan int8, chanNumber)
//	var i int64
//	for ; i < chanNumber; i++ {
//		g.workChan <- 1
//	}
//	g.group.Add(int(workNumber))
//	return g
//}
//
//type globalInfo struct {
//	address      []common.Address
//	topics       [][]common.Hash
//	currentId    int32
//	queue        []*logsWork
//	workNumber   int32
//	timeout      time.Duration
//	offset       int64
//	end          int64
//	group        sync.WaitGroup
//	state        workState  //state 0 is cantWork 1 is success 2 is failed 3 is timeout
//	mutex        sync.Mutex //err mutex
//	err          error
//	errTrigger   sync.Once
//	workMutex    sync.Mutex
//	workChan     chan int8
//	retryTimes   int32
//	controlPanel controlPanel
//	//smooth     int32
//}
//type logsWork struct {
//	id          int32
//	returnValue []types.Log
//	shareInfo   *globalInfo
//	done        chan struct{}
//	filter      ethereum.FilterQuery
//}
//
//func newLogsWork(global *globalInfo) (result *logsWork) {
//	var barrier int32 = 0
//	atomic.LoadInt32(&barrier)
//	value := atomic.AddInt32(&global.currentId, 1)
//	atomic.StoreInt32(&barrier, 1)
//	id := value - 1
//	end := int64(id+1)*MaxQueryBlockSize - 1 + global.offset
//	if end > global.end {
//		end = global.end
//	}
//	result = &logsWork{
//		id:        id,
//		done:      make(chan struct{}, 1),
//		shareInfo: global,
//		filter:    ethereum.FilterQuery{Topics: global.topics, Addresses: global.address, FromBlock: big.NewInt(int64(id)*MaxQueryBlockSize + global.offset), ToBlock: big.NewInt(int64(id+1)*MaxQueryBlockSize - 1 + global.offset)},
//	}
//	result.done <- struct{}{}
//	global.queue[id] = result
//	return result
//}
//
//type controlPanel struct {
//	cond          *sync.Cond
//	state         int32 //state 0: 正常 state 1: 熔断,平滑过度
//	failedTimes   int32
//	sumTimes      int32
//	recoverSignal chan int32
//}
//
//func (cp *controlPanel) smoothRecover() {
//	time.Sleep(DefaultSmoothRecoverTimes)
//	var timeGap = time.NewTicker(DefaultSmoothRecoverTimes)
//	for {
//		select {
//		case <-timeGap.C:
//			cp.cond.Signal()
//		case <-cp.recoverSignal:
//			println("recover")
//			cp.cond.Broadcast()
//			break
//		}
//	}
//	timeGap.Stop()
//}
//func (cp *controlPanel) recover() {
//	atomic.CompareAndSwapInt32(&cp.state, 1, 0)
//	println("recover")
//	cp.sumTimes = 0
//	cp.failedTimes = 0
//	cp.recoverSignal <- 1
//}
//
///**
//后期做平滑过度，目前有点过于消耗CPU资源
//*/
//
//func (work *logsWork) handler() {
//	go func() {
//		<-work.shareInfo.workChan
//		defer func() {
//			work.shareInfo.workChan <- 0
//		}()
//		defer work.shareInfo.group.Done()
//		state := atomic.LoadInt32((*int32)(&work.shareInfo.state))
//		if state == 2 || state == 3 {
//			return
//		}
//		//timer
//		timer := time.NewTimer(work.shareInfo.timeout)
//		for {
//			select {
//			case <-work.done:
//			retryGet:
//				work.shareInfo.workMutex.Lock()
//				//平滑过度
//				if atomic.LoadInt32(&work.shareInfo.controlPanel.state) != 0 {
//					work.shareInfo.controlPanel.cond.Wait()
//				}
//				work.shareInfo.workMutex.Unlock()
//				if atomic.LoadInt32(&work.shareInfo.controlPanel.state) == 1 {
//					atomic.AddInt32(&work.shareInfo.controlPanel.sumTimes, 1)
//					if atomic.LoadInt32(&work.shareInfo.controlPanel.sumTimes) == 0 || float64(work.shareInfo.controlPanel.failedTimes/work.shareInfo.controlPanel.sumTimes) < SmoothRecoverRatio {
//						work.shareInfo.controlPanel.recover()
//					}
//				}
//				logs, err := GetClient().FilterLogs(context.Background(), work.filter)
//				if err != nil {
//					if work.shareInfo.retryTimes >= EmergencyRecovery {
//						if atomic.CompareAndSwapInt32(&work.shareInfo.controlPanel.state, 0, 1) {
//							fmt.Println(work.shareInfo.retryTimes)
//							work.shareInfo.retryTimes = 0
//							work.shareInfo.controlPanel.smoothRecover()
//						}
//					}
//					if strings.Contains(err.Error(), "429 Too Many Requests") {
//						if atomic.LoadInt32(&work.shareInfo.controlPanel.state) == 0 {
//							atomic.AddInt32(&work.shareInfo.retryTimes, 1)
//						} else if atomic.LoadInt32(&work.shareInfo.controlPanel.state) == 1 {
//							atomic.AddInt32(&work.shareInfo.controlPanel.failedTimes, 1)
//						}
//						goto retryGet
//					}
//					//atomic.SwapInt32((*int32)(&work.state), 2)
//					work.shareInfo.errTrigger.Do(func() {
//						work.shareInfo.mutex.Lock()
//						atomic.SwapInt32((*int32)(&work.shareInfo.state), 2)
//						work.shareInfo.err = errors.New("failed")
//						work.shareInfo.mutex.Unlock()
//					})
//					work.shareInfo.mutex.Lock()
//					if work.shareInfo.err != nil {
//						work.shareInfo.err = fmt.Errorf("%v \n %v", work.shareInfo.err, err)
//					}
//					work.shareInfo.mutex.Unlock()
//					return
//				}
//				//atomic.SwapInt32((*int32)(&work.state), 1)
//				work.returnValue = logs
//				return
//			case <-timer.C:
//				//_ = atomic.CompareAndSwapInt32((*int32)(&work.state), 0, 3)
//				work.shareInfo.mutex.Lock()
//				ok := atomic.CompareAndSwapInt32((*int32)(&work.shareInfo.state), 0, 3)
//				if ok {
//					work.shareInfo.err = errors.New("From %s block to %s block search timeout error")
//				}
//				work.shareInfo.mutex.Unlock()
//				return
//			//monitor the global state ,in order to exit in error
//			default:
//				state = atomic.LoadInt32((*int32)(&work.shareInfo.state))
//				if state == 2 || state == 3 {
//					return
//				}
//			}
//		}
//	}()
//}
//
//// logs filter Stream
//
//type LogsStream struct {
//	logs []types.Log
//	err  error
//}
//type FilterFunc func([]types.Log) error
//
//func NewLogsStream(log []types.Log) *LogsStream {
//	return &LogsStream{logs: log}
//}
//
//func (l *LogsStream) FilterLog(filter FilterFunc) *LogsStream {
//	if l.err != nil {
//		return l
//	}
//	err := filter(l.logs)
//	if err != nil {
//		l.err = err
//	}
//	return l
//}
//func (l *LogsStream) Done() (logs []types.Log, err error) {
//	if l.err != nil {
//		return nil, l.err
//	}
//	return l.logs, nil
//}
package util

import (
	"fmt"
	"github.com/cockroachdb/errors"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"golang.org/x/net/context"
	"math/big"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

type workState int32

var (
	timeout workState = 3
	success workState = 1
	failed  workState = 2
	ongoing workState = 0
)

var TimeLess time.Duration = 1<<63 - 1

const MaxQueryBlockSize int64 = 2000

const defaultMallocCap int64 = 1024
const MaxConcurrentNumber = 8e4
const MaxWorkNumber = MaxConcurrentNumber / MaxQueryBlockSize
const EmergencyRecovery = 100

func GetCurrentBlockNumber() (uint64, error) {
	return GetClient().BlockNumber(context.Background())
}

func GetEvent(timeout time.Duration, from int64, to int64, address []common.Address, topics [][]common.Hash) (stream *LogsStream, err error) {
	info := newGlobalInfo(timeout, from, to, address, topics)
	var workNumber = info.workNumber
	var i int32 = 0
	for ; i < workNumber; i++ {
		newLogsWork(info).handler()
	}
	info.group.Wait()
	ok := atomic.CompareAndSwapInt32((*int32)(&info.state), 0, 1)
	if !ok {
		return nil, fmt.Errorf("get event error: %v", info.err)
	}
	logs := info.arrangeLogs()
	stream = &LogsStream{
		logs: logs,
	}
	return stream, nil
}
func (g *globalInfo) arrangeLogs() []types.Log {
	var i int32 = 0
	var result = make([]types.Log, 0, defaultMallocCap)
	for ; i < g.currentId; i++ {
		result = append(result, g.queue[i].returnValue...)
	}
	return result
}

func newGlobalInfo(timeout time.Duration, from int64, to int64, address []common.Address, topics [][]common.Hash) (g *globalInfo) {
	var workNumber = (to - from) / MaxQueryBlockSize
	if MaxQueryBlockSize*workNumber+from != to {
		workNumber++
	}
	g = &globalInfo{end: to, errTrigger: sync.Once{}, mutex: sync.Mutex{}, workNumber: int32(workNumber), address: address, topics: topics, offset: from, timeout: timeout, queue: make([]*logsWork, workNumber), group: sync.WaitGroup{}}
	var chanNumber int64 = workNumber
	if chanNumber > MaxWorkNumber {
		chanNumber = MaxWorkNumber
	}
	g.workChan = make(chan int8, chanNumber)
	var i int64
	for ; i < chanNumber; i++ {
		g.workChan <- 1
	}
	g.group.Add(int(workNumber))
	return g
}

type globalInfo struct {
	address    []common.Address
	topics     [][]common.Hash
	currentId  int32
	queue      []*logsWork
	workNumber int32
	timeout    time.Duration
	offset     int64
	end        int64
	group      sync.WaitGroup
	state      workState //state 0 is cantWork 1 is success 2 is failed 3 is timeout
	mutex      sync.Mutex
	err        error
	errTrigger sync.Once
	workChan   chan int8
	//retryTimes int32
	//smooth     int32
}
type logsWork struct {
	id          int32
	returnValue []types.Log
	shareInfo   *globalInfo
	done        chan struct{}
	filter      ethereum.FilterQuery
}

func newLogsWork(global *globalInfo) (result *logsWork) {
	var barrier int32 = 0
	atomic.LoadInt32(&barrier)
	value := atomic.AddInt32(&global.currentId, 1)
	atomic.StoreInt32(&barrier, 1)
	id := value - 1
	end := int64(id+1)*MaxQueryBlockSize - 1 + global.offset
	if end > global.end {
		end = global.end
	}
	result = &logsWork{
		id:        id,
		done:      make(chan struct{}, 1),
		shareInfo: global,
		filter:    ethereum.FilterQuery{Topics: global.topics, Addresses: global.address, FromBlock: big.NewInt(int64(id)*MaxQueryBlockSize + global.offset), ToBlock: big.NewInt(int64(id+1)*MaxQueryBlockSize - 1 + global.offset)},
	}
	result.done <- struct{}{}
	global.queue[id] = result
	return result
}

/**
后期做平滑过度，目前有点过于消耗CPU资源
*/

func (work *logsWork) handler() {
	go func() {
		<-work.shareInfo.workChan
		defer func() {
			work.shareInfo.workChan <- 0
		}()
		defer work.shareInfo.group.Done()
		state := atomic.LoadInt32((*int32)(&work.shareInfo.state))
		if state == 2 || state == 3 {
			return
		}
		//timer
		timer := time.NewTimer(work.shareInfo.timeout)
		for {
			select {
			case <-work.done:
			retryGet:
				logs, err := GetClient().FilterLogs(context.Background(), work.filter)
				if err != nil {
					//if work.shareInfo.retryTimes > EmergencyRecovery {
					//	work.shareInfo.retryTimes = 0
					//}
					if strings.Contains(err.Error(), "429 Too Many Requests") {
						//work.shareInfo.retryTimes++
						goto retryGet
					}
					//atomic.SwapInt32((*int32)(&work.state), 2)
					work.shareInfo.errTrigger.Do(func() {
						work.shareInfo.mutex.Lock()
						atomic.SwapInt32((*int32)(&work.shareInfo.state), 2)
						work.shareInfo.err = errors.New("failed")
						work.shareInfo.mutex.Unlock()
					})
					work.shareInfo.mutex.Lock()
					if work.shareInfo.err != nil {
						work.shareInfo.err = fmt.Errorf("%v \n %v", work.shareInfo.err, err)
					}
					work.shareInfo.mutex.Unlock()
					return
				}
				//atomic.SwapInt32((*int32)(&work.state), 1)
				work.returnValue = logs
				return
			case <-timer.C:
				//_ = atomic.CompareAndSwapInt32((*int32)(&work.state), 0, 3)
				work.shareInfo.mutex.Lock()
				ok := atomic.CompareAndSwapInt32((*int32)(&work.shareInfo.state), 0, 3)
				if ok {
					work.shareInfo.err = errors.New("From %s block to %s block search timeout error")
				}
				work.shareInfo.mutex.Unlock()
				return
			//monitor the global state ,in order to exit in error
			default:
				state = atomic.LoadInt32((*int32)(&work.shareInfo.state))
				if state == 2 || state == 3 {
					return
				}
			}
		}
	}()
}

type LogsStream struct {
	logs []types.Log
	err  error
}
type FilterFunc func([]types.Log) error

func NewLogsStream(log []types.Log) *LogsStream {
	return &LogsStream{logs: log}
}

func (l *LogsStream) FilterLog(filter FilterFunc) *LogsStream {
	if l.err != nil {
		return l
	}
	err := filter(l.logs)
	if err != nil {
		l.err = err
	}
	return l
}
func (l *LogsStream) Done() (logs []types.Log, err error) {
	if l.err != nil {
		return nil, l.err
	}
	return l.logs, nil
}
