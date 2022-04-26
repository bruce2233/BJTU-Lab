package factorymethod

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
	print("apple\n")
}
func (b Banana) Eat() {
	print("banana\n")
}
func (o Orange) Eat() {
	print("orange\n")
}

type IFactory interface {
	Produce() IFruit
}
type AppleFactory struct {
}
type BananaFactory struct {
}
type OrangeFactory struct {
}

func (af AppleFactory) Produce() IFruit {
	return Apple{}
}
func (bf BananaFactory) Produce() IFruit {
	return Banana{}
}
func (of OrangeFactory) Produce() IFruit {
	return Orange{}
}
func NewIFactory(t string) IFactory {
	switch t {
	case "apple":
		return AppleFactory{}
	case "banana":
		return BananaFactory{}
	case "orange":
		return OrangeFactory{}
	}
	return nil
}
