package repos

type Repoyamlfile struct {
	Name     string   `yaml:"name"`
	Description string   `yaml:"description"`
	Homepage    string `yaml:"homepage"`
	Private     bool   `yaml:"private"`
}

