package main

import (
	"fmt"
	"math/rand"
	"sync"
)

// 使用Mutex编写100个人的投票程序
// 要求一旦有足够的同意或拒绝票后程序终止

func main() {
	exitSign := false
	yes := 0
	no := 0
	var mu sync.Mutex
	cond := sync.NewCond(&mu)
	for i := 0; i < 100; i++ {
		go func() {
			mu.Lock()
			if exitSign {
				return
			}
			defer mu.Unlock()
			if requestVote() {
				yes++
			} else {
				no++
			}
			if yes >= 50 {
				fmt.Printf("enough votes for yes(yes: %d, no: %d)\n", yes, no)
				exitSign = true
			} else if no >= 50 {
				fmt.Printf("enough votes for no(yes: %d, no: %d)\n", yes, no)
				exitSign = true
			}
			cond.Broadcast()
		}()
	}

	mu.Lock()
	for yes < 50 && no < 50 {
		// fmt.Printf("no enough votes(yes:%d, no:%d)\n", yes, no)
		cond.Wait()
	}
	mu.Unlock()
}

func requestVote() bool {
	return rand.Int()%2 == 1
}
