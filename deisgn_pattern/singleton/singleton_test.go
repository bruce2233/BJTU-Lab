package singleton

import (
	"testing"
)

func TestSingleton(t *testing.T) {
	manager := &Multiton{}
	manager.GetInstance()
	t.Log(len(manager.array))
	manager.GetInstance()
	t.Log(len(manager.array))
	manager.GetInstance()
	t.Log(len(manager.array))
	manager.GetInstance()
	t.Log(len(manager.array))
}
