package main

import (
	"itgit.emea.cegedim.grp/plerouge/itcare-go-client/pkg/itcare"
)

func main() {
	client := itcare.ITCareClient{}
	client.Connect()
	client.GetCI("MAMACHINE01")

}
