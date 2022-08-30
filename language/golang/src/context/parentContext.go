package main

import (
	"context"
	"fmt"
	"time"
)

//context中的通道结束之前,如果没有协程监听,那么这个消息永远不会得到
func main() {
	ctx := context.Background()
	sonctx, cancel := context.WithTimeout(ctx, 10*time.Millisecond)
	go func() {
		select {
		case <-sonctx.Done():
			fmt.Println("when preGoroute ending,the sonGoroute is running")
		}
	}()
	cancel()
	time.Sleep(10 * time.Second)
	defer cancel()
}
