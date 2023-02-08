package model

import (
	"encoding/json"
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

type Response struct {
	Response string
	Code     int
}

// end vars

func CreateVol(body []CreateVolume) ([]byte, error) {
	fmt.Printf("Entro en la funcion de crear\n")
	return Response{Response: "OK", Code: 200}, nil
}
