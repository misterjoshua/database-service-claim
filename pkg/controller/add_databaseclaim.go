package controller

import (
	"github.com/misterjoshua/db-claim-operator/pkg/controller/databaseclaim"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, databaseclaim.Add)
}
