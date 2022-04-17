package repositories

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/granitebps/puasa-sunnah-api/types"
)

func SourcesReadFile() ([]types.Source, error) {
	data := []types.Source{}

	filename := "data/sources.json"
	jsonFile, err := os.Open(filename)
	if err != nil {
		return data, err
	}
	defer jsonFile.Close()

	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return data, err
	}

	if err := json.Unmarshal(jsonData, &data); err != nil {
		return data, err
	}

	return data, nil
}
