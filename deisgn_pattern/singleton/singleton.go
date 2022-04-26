package singleton

import (
	"math/rand"
	"sync"
)

var n int = 3

type IMultiton interface {
	GetInstance() Multiton
	Random(int) int
}
type Multiton struct {
	array []Multiton
}

func (mul *Multiton) GetInstance() Multiton {
	//懒汉式双重检测
	mutex := sync.Mutex{}
	if len(mul.array) == n {
		return mul.array[mul.Random(n)]
	}
	mutex.Lock()
	if len(mul.array) < n {
		mul.array = append(mul.array, Multiton{})
	}
	mutex.Unlock()
	return mul.array[mul.Random(len(mul.array))]
}
func (mul *Multiton) Random(n int) int {
	return rand.Intn(n)
}
