package model

type Job struct {
	Id          int                 `yaml:"id"`
	Name        string              `yaml:"name"`
	Environment []map[string]string `yaml:"environment"`
	Stages      []Stages            `yaml:"stages"`
}

type Stage struct {
	Name string `yaml:"name"`
	Cmd  string `yaml:"cmd"`
}

type Stages struct {
	Stage Stage `yaml:"stage"`
}
