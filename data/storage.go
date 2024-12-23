package data

import (
	"encoding/json"
	"fmt"
	"os"
)

func SaveFinanceData(data *FinanceData) error {
	file, err := os.Create(data.FilePath)
	if err != nil {
		return fmt.Errorf("could not create file: %v", err)
	}
	defer file.Close()

	enconder := json.NewEncoder(file)
	err = enconder.Encode(data)
	if err != nil {
		return fmt.Errorf("could not encode data: %v", err)
	}

	return nil
}

func LoadFinanceData(filename string) (*FinanceData, error) {
	file, err := os.Open(filename)
	if os.IsNotExist(err) {
		newFinanceData := &FinanceData{FilePath: filename}
		err = SaveFinanceData(newFinanceData)
		if err != nil {
			return nil, err
		}
		return newFinanceData, nil
	}

	if err != nil {
		return &FinanceData{}, fmt.Errorf("could not open file: %v", err)
	}
	defer file.Close()

	var data FinanceData
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&data)
	if err != nil {
		return &FinanceData{}, fmt.Errorf("could not decode data: %v", err)
	}

	return &data, nil
}
