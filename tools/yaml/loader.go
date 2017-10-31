package yaml

import (
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

// Load loads the YAML data from the provided file and unmarshals it into the
// given object. Said object needs to be a pointer.
func Load(filePath string, v interface{}) error {
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("yaml_load: error reading file. err: %s", err.Error())
	}

	if err = yaml.Unmarshal(fileContent, v); err != nil {
		return fmt.Errorf("yaml_load: error unmarshalling YAML data into %T. err: %s", v, err.Error())
	}

	return nil
}
