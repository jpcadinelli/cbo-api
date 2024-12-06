package dataprovider

import (
	"cbo-api/domain"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

func GetListCBO() ([]domain.CBO, error) {
	file, err := os.Open("cbo.json")
	if err != nil {
		return []domain.CBO{}, fmt.Errorf("unable to read cbo.json: %v", err)
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			log.Println("Error closing file:", err)
		}
	}(file)

	content, err := io.ReadAll(file)
	if err != nil {
		return []domain.CBO{}, fmt.Errorf("error reading cbo.json: %v", err)
	}

	var cbos []domain.CBO
	err = json.Unmarshal(content, &cbos)
	if err != nil {
		return []domain.CBO{}, fmt.Errorf("error parsing cbo.json: %v", err)
	}

	return cbos, nil
}
