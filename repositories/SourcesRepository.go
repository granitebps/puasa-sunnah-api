package repositories

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/granitebps/puasa-sunnah-api/types"
)

func SourcesReadFile() ([]types.Source, error) {
	data := []types.Source{}

	filename := "data/sources.json"
	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Printf("failed to open json file: %s, error: %v", filename, err)
		return data, err
	}
	defer jsonFile.Close()

	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Printf("failed to read json file, error: %v", err)
		return data, err
	}

	if err := json.Unmarshal(jsonData, &data); err != nil {
		fmt.Printf("failed to unmarshal json file, error: %v", err)
		return data, err
	}

	return data, nil
}
