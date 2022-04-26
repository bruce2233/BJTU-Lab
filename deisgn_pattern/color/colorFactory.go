package factory

import (
	"strconv"
)

func ColorStrFactory(s string, front int, background int, special int) (string, string) {

	rtn := "\033[" + strconv.Itoa(front) + ";" + strconv.Itoa(background) + ";" + strconv.Itoa(special) + "m%s\033[0m"
	return rtn, s
}
