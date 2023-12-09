package services

import (
	"encoding/json"
	"fmt"
	"mdhesari/coralflora/model"
	"os"
)

var (
	file []byte
	err  error
)

func ReadFromFile(path string) ([]byte, error) {
	file, err = os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func CreateFile(path string) (*os.File, error) {
	f, ferr := os.Create(path)
	if ferr != nil {

		return nil, ferr
	}
	defer f.Close()

	return f, nil
}

func Update(path string, flowers []*model.Flower) {
	contents, _ := json.Marshal(flowers)

	fmt.Println(path)

	err = os.WriteFile(path, contents, 0644)

	if err != nil {
		fmt.Println(err)
	}

	contents, err = ReadFromFile(path)

	if err != nil {
		fmt.Println(err)
	}
}
