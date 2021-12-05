package utils

import (
	"fmt"
)

func Inspect(objects ...interface{}) {
	for _, object := range objects {
		fmt.Printf("%+v\n", object)
	}
}
