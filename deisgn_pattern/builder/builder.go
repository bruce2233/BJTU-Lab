package builder

import "fmt"

type ICar interface {
	Run()
}

type Car struct {
	Tyre       string
	SteerWheel string
	Motor      string
}
type ICarBuilder interface {
	BuildTyre()
	BuildSteerWheel()
	BuildMotor()
	Build() Car
}

type BMWBuilder struct {
	Car
}

func (bmwbuilder BMWBuilder) BuildTyre() {
	bmwbuilder.Car.Tyre = "BMW Tyre"
}
func (bmwbuilder BMWBuilder) BuildSteerWheel() {
	bmwbuilder.Car.SteerWheel = "BMW SteerWheel"
}
func (bmwbuilder BMWBuilder) BuildMotor() {
	bmwbuilder.Car.Motor = "BMW Motor"
}
func (bmwbuilder BMWBuilder) Build() Car{
	return bmwbuilder.Car
}
func (car Car) Run() {
	fmt.Println("I'm running...")
}
