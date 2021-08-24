package main

import (
	"context"
	"fmt"
	"time"
)
func main() {
	// 设置上下文10s的超时时间
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second * 10)
	go func() {
		// 1s后主动取消上下文
		<-time.After(time.Second)
		cancelFunc()
	}()
	<-ctx.Done()
	err := ctx.Err()
	fmt.Printf("err == nil ? %v\n", err == nil)
	// 输出 err == nil ? false
}

func worker(ctx context.Context){
	for {
		select {
		case <- ctx.Done():
			fmt.Println("下班咯~~~")
			return
		default:
			fmt.Println("认真摸鱼中，请勿打扰...")
		}
		time.Sleep(1*time.Second)
	}
}