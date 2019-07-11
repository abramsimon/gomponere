package model

type Team struct {
	Name        string      `yaml:"name"`
	TeamContact TeamContact `yaml:"team-contact"`
	LeadContact TeamContact `yaml:"lead-contact"`
	Display     Display     `yaml:"display"`
}

type TeamContact struct {
	Name  string `yaml:"name"`
	Email string `yaml:"email"`
}

type Display struct {
	BackgroundColor string `yaml:"background-color"`
	ForegroundColor string `yaml:"foreground-color"`
}
