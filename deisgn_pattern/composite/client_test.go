package composite

import (
	"testing"
)

func TestClient(t *testing.T) {
	school := NewSchool()
	for i := 0; i < 10; i++ {
		class := NewClass()
		for j := 0; j < 30; j++ {
			class.Add(NewStudent())
		}
		school.Add(class)
	}
	school.Add(NewStudent())
	school.Add(NewStudent())
	school.Add(NewStudent())
	school.Notify()
}
