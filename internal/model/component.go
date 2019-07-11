package model

type Component struct {
	Name           string   `yaml:"name"`
	Description    string   `yaml:"description"`
	Git            string   `yaml:"git"`
	ReleaseDate    string   `yaml:"release-date"`
	LevelKey       string   `yaml:"level"`
	TypeKey        string   `yaml:"type"`
	TeamKey        string   `yaml:"team"`
	AreaKey        string   `yaml:"area"`
	DependencyKeys []string `yaml:"dependencies"`
}
