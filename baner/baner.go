package baner

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
)

func ImageFunc(ByneryImage string, Idage string,folder string) string{

	bynaryData, err := base64.StdEncoding.DecodeString(ByneryImage)
	if err != nil {
		log.Printf("Error image : %v", err)
	}

	errWriteFile := ioutil.WriteFile("./Statics/"+folder+"/"+Idage, bynaryData, 0644)
	if errWriteFile != nil {
		log.Printf("Error write: %v", errWriteFile)
	}
	return fmt.Sprintf(Idage)
}
