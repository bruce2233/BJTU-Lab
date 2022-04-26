package main

import (
	"fmt"
	"testing"
)

func TestProtypeShallow(t *testing.T) {

	// ori := BMWShallow{Tyre: Tyre{"my"}}
	// var i ICar = &ori
	// o := i.(*BMWShallow)
	// t.Log(ori.tyre)
	// o.tyre = "another"
	// t.Log(ori.tyre)

	bmwShallow := BMWShallow{
		Tyre: Tyre{"a bmw"},
	}
	c := bmwShallow.Clone()
	clone := c.(*BMWShallow)
	clone.tyre = "a benz"
	fmt.Printf("%p\n", &bmwShallow)
	fmt.Printf("%p\n", clone)
	fmt.Println(bmwShallow.tyre)
}

func TestProtypeDeep(t *testing.T) {
	bmwShallow := BMWDeep{
		Tyre: Tyre{tyre: "a bmw"},
	}
	clonedBmwDeep := bmwShallow.Clone().(BMWDeep)
	clonedBmwDeep.tyre = "a benz"
	fmt.Println(bmwShallow.tyre)
}
