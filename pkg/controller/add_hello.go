package controller

import (
	"goturkiye-operator/pkg/controller/hello"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, hello.Add)
}
