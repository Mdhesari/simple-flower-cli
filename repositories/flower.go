package repositories

import (
	"encoding/json"
	"fmt"
	"mdhesari/coralflora/model"
	"mdhesari/coralflora/services"
	"os"
)

func FindById(flowers []*model.Flower, id int) (*model.Flower, int) {
	for i, x := range flowers {
		if x.Id == id {
			return x, i
		}
	}

	// Return nil pointer when not found
	return nil, 0
}

func FindByName(flowers []*model.Flower, name string) (*model.Flower, int) {
	for i, x := range flowers {
		if x.Name == name {
			return x, i
		}
	}

	// Return nil pointer when not found
	return nil, 0
}

func GetItems() []*model.Flower {
	var items []*model.Flower

	contents, err := services.ReadFromFile(getPath())
	if err != nil {
		fmt.Println(err)
	}

	json.Unmarshal(contents, &items)

	return items
}

func UpdateItems(flowers []*model.Flower) {
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
