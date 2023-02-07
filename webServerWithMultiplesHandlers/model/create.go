package model

import (
	"fmt"
	//"io/ioutil"
	//"log"
)

// Define all vars
type CreateVolume struct {
	zone        string
	application string
	environment string
	COMPONENT   string
	pv          string
}

// end vars

func CreateVol(body []CreateVolume) {
	fmt.Printf("Entro en la funcion de crear\n")
}
