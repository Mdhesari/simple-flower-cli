package repositories

import (
	"fmt"
	"mdhesari/coralflora/model"
	"mdhesari/coralflora/services"
	"os"
)

func findById() {
	//
}

func updateItems(flowers []*model.Flower) {
	path := getPath()

	// update whole file with new items
	services.Update(path, flowers)
}

func getPath() string {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error")

		return ""
	}

	return pwd + "/flowers.json"
}
