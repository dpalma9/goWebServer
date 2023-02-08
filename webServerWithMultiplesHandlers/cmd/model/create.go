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
	Response string `json:"response"`
	Code     int    `json:"code"`
}

// end vars

func CreateVol(body []CreateVolume) ([]byte, error) {
	fmt.Printf("Entro en la funcion de crear\n")
	res_json := &Response{"OK", 200}
	res, _ := json.Marshal((res_json))
	fmt.Printf("Salgo de la funcion de crear\n")
	return res, nil
}
