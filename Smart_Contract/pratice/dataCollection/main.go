package main

import (
	"dataCollection/util"
	"log"
)

func main() {
	// Get the latest known block
	number, err := util.GetCurrentBlockNumber()
	if err != nil {
		log.Fatal("Get Current blockchain info error:", err)
		return
	}
	//util.GetClient().
	//util.GetEvent(util.TimeLess, number)
}
