package export

import (
	"encoding/json"
	"io/ioutil"

	"github.com/pkg/errors"
)

func JSON(filePath string, data interface{}) error {
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return errors.Wrap(err, "failed to marshal data")
	}

	err = ioutil.WriteFile(filePath+".json", file, 0644)
	if err != nil {
		return errors.Wrap(err, "failed to write json file")
	}

	return nil
}
