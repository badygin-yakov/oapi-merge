package merge

import (
	"log"

	"gopkg.in/yaml.v2"
)

func ParseToStruct(data []byte) (OpenAPI, error) {
	res := OpenAPI{}

	err := yaml.Unmarshal(data, &res)
	if err != nil {
		return OpenAPI{}, err
	}

	return res, nil
}

func StructToByte(data OpenAPI) ([]byte, error) {
	d, err := yaml.Marshal(data)
	if err != nil {
		log.Fatalf("error: %v", err)
		return nil, err
	}

	return d, nil
}
