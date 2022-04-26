package factorymethod

import "testing"

func TestClient(t *testing.T) {

	var factory IFactory
	var fruit IFruit
	factory = NewIFactory("apple")
	fruit = factory.produce()
	fruit.Eat()
	factory = NewIFactory("banana")
	fruit = factory.produce()
	fruit.Eat()
	factory = NewIFactory("orange")
	fruit = factory.produce()
	fruit.Eat()
}
