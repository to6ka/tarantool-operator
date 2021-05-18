package controller

import (
	"github.com/to6ka/tarantool-operator/pkg/controller/cluster"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, cluster.Add)
}
