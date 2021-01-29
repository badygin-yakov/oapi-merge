package merge

import (
	"io/ioutil"
)

func ReadFromFile(fileName string) ([]byte, error) {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	return content, nil
}
