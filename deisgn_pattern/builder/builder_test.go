package builder

import (
	"testing"
)

func TestBuilder(t *testing.T) {
	bmwbuilder := BMWBuilder{}
	bmwbuilder.BuildTyre()
	bmwbuilder.BuildTyre()
	bmwbuilder.BuildTyre()
	car := bmwbuilder.Build()
	car.Run()
}
