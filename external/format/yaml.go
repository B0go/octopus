package format

import (
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

//YamlManipuler manipulates yaml files
type YamlManipuler interface {
	PrintAsYaml(interface{}) error
	ReadYaml(interface{}, string) error
}

//DefaultYamlManipuler is the default manipuler of yaml files
type DefaultYamlManipuler struct {
}

//PrintAsYaml print the passed struct as yaml on console
func (yamlManipuler DefaultYamlManipuler) PrintAsYaml(class interface{}) error {
	yaml, err := yaml.Marshal(&class)
	if err != nil {
		return err
	}

	fmt.Printf("\n%s\n", string(yaml))
	return nil
}

//ReadYaml reads the YAML in the provided path and populates the provided struct
func (yamlManipuler DefaultYamlManipuler) ReadYaml(class interface{}, path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(file, class)
	if err != nil {
		return err
	}
	return nil
}
