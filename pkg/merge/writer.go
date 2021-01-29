package merge

import "os"

func WriteToFile(fileName string, data []byte) error {
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}

	defer f.Close()

	_, err = f.Write(data)
	if err != nil {
		return err
	}

	return nil
}
