package composite

type IComponent interface {
	Notify()
}

type ILeaf interface {
	IComponent
}
type IContainer interface {
	IComponent
	Add(IComponent)
	Remove(IComponent)
	GetChild(i int) IComponent
}
type ContainerImpl struct {
	childs []IComponent
}
type School struct {
	*ContainerImpl
}
type Class struct {
	*ContainerImpl
}
type Student struct{}

func (container *ContainerImpl) Notify() {
	for _, value := range container.childs {
		value.Notify()
	}
}
func (container *ContainerImpl) Add(com IComponent) {
	container.childs = append(container.childs, com)
}

func (container *ContainerImpl) Remove(com IComponent) {
	for index, value := range container.childs {
		if value == com {
			container.childs = append(container.childs[:index], container.childs[index+1:]...)
		}
	}
}

func (container *ContainerImpl) GetChild(i int) IComponent {
	return container.childs[i]
}

func (student Student) Notify() {
	println("notify student")
}

func NewSchool() *School {
	return &School{&ContainerImpl{}}
}

func NewClass() *Class {
	return &Class{&ContainerImpl{}}
}
func NewStudent() *Student {
	return &Student{}
}
