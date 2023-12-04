package services

import (
	"encoding/json"
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

	return f, nil
}

func Update(path string, flowers []*model.Flower) {
	contents, _ := json.Marshal(flowers)

	os.WriteFile(path, contents, 0)
}
