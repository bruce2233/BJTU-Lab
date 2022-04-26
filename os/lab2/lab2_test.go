package lab2

import (
	"testing"
)

func BenchmarkMakeDeal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		wg.Add(2)
		deal := &Deal{}
		go deal.MakeDeal()
		go deal.MakeDeal()
		wg.Wait()
	}
	println("times: ", times1)
	println("wrong: ", wrong1)
}

func BenchmarkSyncDeal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		wg.Add(2)

		deal := &Deal{}
		interested[0] = false
		interested[1] = false
		turn = -1
		go deal.SyncDeal(1)
		go deal.SyncDeal(0)
		wg.Wait()

	}
	println("times: ", times2)
	println("wrong: ", wrong2)

}

func BenchmarkMutexDeal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		wg.Add(2)
		deal := &Deal{}
		go deal.MutexDeal()
		go deal.MutexDeal()
		wg.Wait()
	}
	println("times: ", times3)
	println("wrong: ", wrong3)

}
