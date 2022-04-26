package simplefactory

import "testing"

func TestSimpleFactory(t *testing.T){
	var myFruit IFruit 
	myFruit = NewIFruit("apple")
	myFruit.Eat()
	myFruit = NewIFruit("banana")
	myFruit.Eat()
	myFruit = NewIFruit("orange")
	myFruit.Eat()
}