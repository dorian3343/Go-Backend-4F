package shop

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

const filePermission = 0644

type Product struct {
	Id    string
	Price int
	Name  string
}

func (p Product) Setup(path string) error {
	filePath := filepath.Join(path, p.Id+".json")

	f, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
		}
	}(f)

	productJSON, err := json.MarshalIndent(p, "", "    ")
	if err != nil {
		return fmt.Errorf("error marshaling JSON: %v", err)
	}

	err = os.WriteFile(f.Name(), productJSON, filePermission)
	if err != nil {
		return fmt.Errorf("error writing to file: %v", err)
	}

	return nil
}

func ExtractJSONShop(path string) (Product, error) {
	var data Product

	content, err := os.ReadFile(path)
	if err != nil {
		log.Printf("error reading file: %v", err)
		return Product{}, err
	}

	err = json.Unmarshal(content, &data)
	if err != nil {
		return Product{}, fmt.Errorf("error unmarshaling JSON: %v", err)
	}

	return data, nil
}
