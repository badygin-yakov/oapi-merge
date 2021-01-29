package merge

type Components struct {
	Schemas         map[string]interface{} `yaml:"schemas"`
	SecuritySchemes map[string]interface{} `yaml:"securitySchemes"`
}

type OpenAPI struct {
	Openapi string                 `yaml:"openapi"`
	Info    map[string]interface{} `yaml:"info"`
	Servers []interface{}          `yaml:"servers"`
	Tags    []struct {
		Name        string `yaml:"name"`
		Description string `yaml:"description"`
	} `yaml:"tags"`
	Paths      map[string]interface{} `yaml:"paths"`
	Components Components             `yaml:"components"`
}
