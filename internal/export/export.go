package export

import (
	"encoding/csv"
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/jonfriesen/symbol-list/internal/model"
	"github.com/pkg/errors"
)

func JSON(filePath string, data *model.Export) error {
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

func CSV(filePath string, data *model.Export) error {
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

	for _, row := range data.Data {
		err = csvwriter.Write(row.Row())
		if err != nil {
			return errors.Wrap(err, "failed writing csv file")
		}
	}

	csvwriter.Flush()
	csvfile.Close()

	return nil
}
