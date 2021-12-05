package utils

import (
	"fmt"

	"github.com/fatih/color"
)

func Inspect(objects ...interface{}) {
	for i, object := range objects {
		if i == 0 {
			color.Green(fmt.Sprintf("%+v\n", object))
		} else {
			fmt.Printf("%+v\n", object)
		}
	}
}
