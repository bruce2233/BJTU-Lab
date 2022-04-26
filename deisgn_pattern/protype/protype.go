package main

type ICar interface {
	Clone() ICar
}
type Tyre struct {
	tyre string
}
type SteeringWheel struct {
}
type Motor struct {
}

type BMWShallow struct {
	Tyre
	SteeringWheel
	Motor
}

func (bmw *BMWShallow) Clone() ICar {
	return bmw
}

type BMWDeep struct {
	Tyre
	SteeringWheel
	Motor
}

func (bmw BMWDeep) Clone() ICar {
	return bmw
}
