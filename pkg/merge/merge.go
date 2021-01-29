package merge

import "fmt"

func Do(filePaths []string, resultFilePath string) error {
	specs, err := loadSpecs(filePaths)
	if err != nil {
		return err
	}

	res := merge(specs)

	if resultFilePath != "" {
		return save(resultFilePath, res)
	}

	data, err := StructToByte(res)
	if err != nil {
		return err
	}

	fmt.Println(string(data))

	return nil
}

func merge(specs []OpenAPI) OpenAPI {
	if len(specs) < 1 {
		return OpenAPI{}
	}

	start := specs[0]
	res := OpenAPI{
		Openapi: start.Openapi,
		Info:    start.Info,
		Servers: start.Servers,
		Paths:   map[string]interface{}{},
		Components: Components{
			Schemas:         map[string]interface{}{},
			SecuritySchemes: start.Components.SecuritySchemes,
		},
	}

	for _, spec := range specs {
		res.Tags = append(res.Tags, spec.Tags...)

		for path, data := range spec.Paths {
			res.Paths[path] = data
		}

		for key, value := range spec.Components.Schemas {
			res.Components.Schemas[key] = value
		}
	}

	return res
}

func loadSpecs(filePaths []string) ([]OpenAPI, error) {
	specs := make([]OpenAPI, len(filePaths))

	for index, filePath := range filePaths {
		data, err := ReadFromFile(filePath)
		if err != nil {
			return nil, err
		}

		spec, err := ParseToStruct(data)
		if err != nil {
			return nil, err
		}

		specs[index] = spec
	}

	return specs, nil
}

func save(outFileName string, res OpenAPI) error {
	d, err := StructToByte(res)
	if err != nil {
		return err
	}

	return WriteToFile(outFileName, d)
}
