package controller

import (
	"github.com/integr8ly/3scale-operator/pkg/controller/threescale"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, threescale.Add)
}
