package config

type Config struct {
	Repos []Repository `yaml:"repos"`
	// destinations []Destinations
}

type Repository struct {
	Name          string        `yaml:"name"`
	Path          string        `yaml:"path"`
	DefaultBranch string        `yaml:"defaultBranch"`
	Destinations  []Destination `yaml:"destinations"`
}

type Destination struct {
	RemoteName string `yaml:"remoteName"`
}
