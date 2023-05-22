package main

import (
	"fmt"

	"github.com/doctori/itcare-go-client/pkg/itcare"
)

func main() {
	client := itcare.ITCareClient{}
	client.Connect()
	ci, err := client.GetCI("MAMACHINE")
	if err != nil {
		panic(err)
	}
	fmt.Printf("The IP of the Instance is %s", ci.IPAddress)

}
