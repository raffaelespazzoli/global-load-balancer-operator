package controller

import (
	"github.com/redhat-cop/global-load-balancer-operator/pkg/controller/globalroutediscovery"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, globalroutediscovery.Add)
}
