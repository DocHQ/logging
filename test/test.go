package test

import (
	"fmt"
)

type Logger struct{}

func (t Logger) Log(i interface{}, fields interface{}, level uint32, verbose bool) {
	fmt.Printf("%v %v %v %v \n", i, fields, level, verbose)
}
