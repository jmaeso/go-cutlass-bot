package yaml

import (
	"errors"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type testData struct {
	Foo struct {
		Bar string `yaml:"bar"`
	} `yaml:"foo"`
}

func TestYaml_Load(t *testing.T) {
	Convey("when the file does not exist", t, func() {
		filePath := "test_data/non_existent.yml"
		v := new(testData)

		expectedErr := errors.New("yaml_load: error reading file. err: open test_data/non_existent.yml: no such file or directory")

		Convey("it returns an error", func() {
			err := Load(filePath, v)

			So(expectedErr.Error(), ShouldEqual, err.Error())
		})
	})

	Convey("when fails unmarshalling data", t, func() {
		filePath := "test_data/data_invalid.yml"
		v := new(testData)

		expectedErr := errors.New("yaml_load: error unmarshalling YAML data into *yaml.testData. err: yaml: line 1: mapping values are not allowed in this context")

		Convey("it returns an error", func() {
			err := Load(filePath, v)

			So(expectedErr.Error(), ShouldEqual, err.Error())
		})
	})

	Convey("when the loading is successful", t, func() {
		filePath := "test_data/data.yml"
		v := new(testData)

		Convey("it returns a nil error", func() {
			err := Load(filePath, v)

			So(err, ShouldBeNil)
			So(v.Foo.Bar, ShouldNotBeEmpty)
		})
	})
}
