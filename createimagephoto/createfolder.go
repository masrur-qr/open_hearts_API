package createimagephoto

import (
	"fmt"
	"os"
)


func CreateFolder(folder_Name string) error {
	err := os.Mkdir("Statics"+"/"+folder_Name, os.ModePerm)
	if err != nil {
		if os.IsExist(err) {
			fmt.Println("Папка уже существует")
		}
		return err
	}
	return nil
}
