package config

import (
	"fmt"

	"github.com/B0go/octopus/external/format"
	"github.com/B0go/octopus/external/system"

	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
)

func init() {
	log.SetHandler(cli.Default)
}

//Project represents a project to be managed
type Project struct {
	Name       string `yaml:"name"`
	Repository string `yaml:"repository"`
	Team       string `yaml:"team"`
}

//Config represents the configuration containing the projects that will be managed
type Config struct {
	Projects []Project `yaml:"projects"`
}

//Load loads the configuratin file at {UserHome}/.octopus/config.yaml
func Load(usrRetriever system.UserRetriever, fsReader system.FileSystemReader, yamlManipuler format.YamlManipuler) (*Config, error) {
	config := Config{}

	usr, err := usrRetriever.Current()
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("%s/.octopus/config.yaml", usr.HomeDir)

	if _, err := fsReader.FileStatus(path); fsReader.IsNotExist(err) {
		log.WithField("path", path).
			Errorf("Octopus is not configured. config.yaml not found!")
		return nil, err
	}

	err = yamlManipuler.ReadYaml(&config, path)
	if err != nil {
		return nil, err
	}
	return &config, err

}
