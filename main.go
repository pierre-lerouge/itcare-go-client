package main

import (
	"itgit.emea.cegedim.grp/plerouge/itcare-go-client/pkg/itcare"
)

func main() {
	client := itcare.ITCareClient{
		ClientID:     "cgdm-sa-xoxo",
		ClientSecret: "xoxo",
	}
	client.Connect()
	client.GetCI("PEBCNRSWEB02")
}
