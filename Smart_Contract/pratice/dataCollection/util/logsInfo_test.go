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
	event, err := GetEvent(TimeLess, 39000000, 39791976, []common.Address{common.HexToAddress("0xC7728354f9fe0e43514B1227162D5B0E40FaD410")}, nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	done, err := event.Done()

	if err != nil {
		return
	}
	fmt.Println(len(done))
}

func TestGetEvent(t *testing.T) {
	var length int
	for i := 39000000; i < 39791976; {
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
