package export

import (
	"encoding/csv"
	"encoding/json"
	"io/ioutil"
	"os"

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

type Export interface {
	CSVHeader() []string
	Data() [][]string
}

func CSV(filePath string, data Export) error {
	csvfile, err := os.Create(filePath + ".csv")
	if err != nil {
		return errors.Wrap(err, "failed creating file")
	}

	csvwriter := csv.NewWriter(csvfile)

	// csv header
	err = csvwriter.Write(data.CSVHeader())
	if err != nil {
		return errors.Wrap(err, "failed writing header")
	}

	for _, row := range data.Data() {
		err = csvwriter.Write(row)
		if err != nil {
			return errors.Wrap(err, "failed writing csv file")
		}
	}

	csvwriter.Flush()
	csvfile.Close()

	return nil
}
