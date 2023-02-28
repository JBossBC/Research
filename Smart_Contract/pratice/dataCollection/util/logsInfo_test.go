package util

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"testing"
)

func TestGetLogs(t *testing.T) {
	//number, _ := GetCurrentBlockNumber()
	event, err := GetEvent(TimeLess, 39710000, 39791976, []common.Address{common.HexToAddress("0xC7728354f9fe0e43514B1227162D5B0E40FaD410")}, nil)
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
