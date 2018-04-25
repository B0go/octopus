package format

import (
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

//YmlManipuler manipulates yml files
type YmlManipuler interface {
	PrintAsYml(interface{}) error
	ReadYml(interface{}, string) error
}

//DefaultYmlManipuler is the default manipuler of yml files
type DefaultYmlManipuler struct {
}

//PrintAsYml print the passed struct as yml on console
func (ymlManipuler DefaultYmlManipuler) PrintAsYml(class interface{}) error {
	yml, err := yaml.Marshal(&class)
	if err != nil {
		return err
	}

	fmt.Printf("\n%s\n", string(yml))
	return nil
}

//ReadYml reads the YML in the provided path and populates the provided struct
func (ymlManipuler DefaultYmlManipuler) ReadYml(class interface{}, path string) error {
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
