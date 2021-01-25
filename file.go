package dockercompose

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

//File docker compose file
type File struct {
	Version  string             `yaml:"version"`
	Services map[string]Service `yaml:"services"`
	Networks map[string]Network `yaml:"networks,omitempty"`
}

//Network dc object
type Network struct {
	External bool `yaml:"external,omitempty"`
}

//Service dc object
type Service struct {
	Image       string   `yaml:"image,omitempty"`
	Build       string   `yaml:"build,omitempty"`
	Command     string   `yaml:"command,omitempty"`
	Environment []string `yaml:"environment,omitempty"`
	Volumes     []string `yaml:"volumes,omitempty"`
	DependsOn   []string `yaml:"depends_on,omitempty"`
	Ports       []string `yaml:"ports,omitempty"`
	NetworkMode string   `yaml:"network_mode,omitempty"`
	Networks    []string `yaml:"networks,omitempty"`
	Links       []string `yaml:"links,omitempty"`
	Expose      []string `yaml:"expose,omitempty"`
	Scale       int      `yaml:"scale,omitempty"`
}

//SaveFile save as a file
func (f File) SaveFile(path string) error {
	d, err := f.SaveBytes()
	if err != nil {
		return err
	}

	return ioutil.WriteFile(path, d, 0644)
}

//SaveBytes saves file as byes
func (f File) SaveBytes() ([]byte, error) {
	return yaml.Marshal(&f)
}

//LoadFile the docker file
func LoadFile(path string) (File, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return File{}, err
	}

	return LoadBytes(data)
}

//LoadBytes loads the file from byte array
func LoadBytes(data []byte) (File, error) {
	file := File{}

	err := yaml.Unmarshal(data, &file)
	if err != nil {
		return file, err
	}

	return file, nil
}
