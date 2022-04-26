package lab2

import (
	"math/rand"
	"sync"
	"time"
)

type Deal struct {
	mutex sync.Mutex
	Payer int
	Payee int
}

var interested [2]bool
var turn, wrong1, wrong2, wrong3 int
var times1, times2, times3 int
var wg sync.WaitGroup

const waitTime = 1000

func (deal *Deal) MakeDeal() {
	times1++
	defer wg.Done()
	dealAmount := rand.Intn(1000000)
	deal.Payer += dealAmount
	time.Sleep(waitTime * time.Nanosecond)
	deal.Payee -= dealAmount

	if deal.Payee != -deal.Payer {
		wrong1++
	}
}

func (deal *Deal) SyncDeal(num int) {
	times2++
	defer wg.Done()

	other := 1 - num
	interested[num] = true
	turn = num
	for turn == num && interested[other] {
	}

	dealAmount := rand.Intn(1000000)
	deal.Payer += dealAmount
	time.Sleep(waitTime * time.Nanosecond)

	deal.Payee -= dealAmount
	if deal.Payee != -deal.Payer {
		wrong2++
	}
	interested[num] = false //退出

}

func (deal *Deal) MutexDeal() {
	times3++
	defer wg.Done()

	mu := &deal.mutex
	mu.Lock()
	defer mu.Unlock()

	dealAmount := rand.Intn(1000000)
	deal.Payer += dealAmount
	time.Sleep(waitTime * time.Nanosecond)

	deal.Payee -= dealAmount
	if deal.Payee != -deal.Payer {
		wrong3++
	}
}
