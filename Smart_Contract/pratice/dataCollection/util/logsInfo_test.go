package util

import (
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"golang.org/x/net/context"
	"math/big"
	"testing"
)

func TestGetLogs(t *testing.T) {
	//number, _ := GetCurrentBlockNumber()
	event, err := GetEvent(TimeLess, 39700000, 39791976, []common.Address{common.HexToAddress("0xC7728354f9fe0e43514B1227162D5B0E40FaD410")}, nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	done, err := event.Done()
	if err != nil {
		return
	}
	fmt.Println(len(done))
	fmt.Println(done[0].Index)
	fmt.Println(done[0].Data)
	fmt.Println(done[0].TxHash)
	fmt.Println(len(done[0].Topics))
	fmt.Println(len(done[0].Data))
}

func TestGetEvent(t *testing.T) {
	var length int
	for i := 39700000; i < 39791976; {
		to := i + 2000
		if to > 39791976 {
			to = 39791976
		}
		query := ethereum.FilterQuery{FromBlock: big.NewInt(int64(i)), ToBlock: big.NewInt(int64(to)), Addresses: []common.Address{common.HexToAddress("0xC7728354f9fe0e43514B1227162D5B0E40FaD410")}, Topics: nil}
		logs, err := client.FilterLogs(context.Background(), query)
		if err != nil {
			print("error")
			return
		}
		length += len(logs)
		i = to
	}
}

func BenchmarkLogConcurrent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := GetEvent(TimeLess, 39000000, 39791976, []common.Address{common.HexToAddress("0xC7728354f9fe0e43514B1227162D5B0E40FaD410")}, nil)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
}
func BenchmarkLog(b *testing.B) {
	var target = 39700000
	for i := 0; i < b.N; i++ {
		for target < 39791976 {
			to := target + 2000
			if to > 39791976 {
				to = 39791976
			}
			query := ethereum.FilterQuery{FromBlock: big.NewInt(int64(i)), ToBlock: big.NewInt(int64(to)), Addresses: []common.Address{common.HexToAddress("0xC7728354f9fe0e43514B1227162D5B0E40FaD410")}, Topics: nil}
			_, err := client.FilterLogs(context.Background(), query)
			if err != nil {
				print("error")
				return
			}
			target = to
		}
	}
}
