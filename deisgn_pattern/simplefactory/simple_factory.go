package simplefactory

type IFruit interface {
	Eat()
}
type Apple struct {
}
type Banana struct {
}
type Orange struct {
}

func (a Apple) Eat() {
	print("apple")
}
func (b Banana) Eat() {
	print("banana")
}
func (o Orange) Eat() {
	print("orange")
}
func NewIFruit(t string) IFruit {
	switch t {
	case "apple":
		return Apple{}

	case "banana":
		return Banana{}

	case "orange":
		return Orange{}
	}
	return nil
}
