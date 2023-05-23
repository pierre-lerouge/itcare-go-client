package main

import (
	"fmt"

	"github.com/pierre-lerouge/itcare-go-client/pkg/itcare"
)

func main() {
	client := itcare.ITCareClient{}
	client.Connect()
	ci, err := client.GetInstance("PEBCNRSWEB01")
	if err != nil {
		panic(err)
	}
	fmt.Printf("The IP of the Instance is %s\n", ci.IPAddress)
	ci, err = client.GetInstanceByID(5220149, true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("The Name of the Instance is %s\n", ci.Name)
	fmt.Printf("The FQDN of the instance is %s\n", ci.Network.DNS[0].Alias)
}
