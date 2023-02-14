package importer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"time"
)

const datePattern = `(\d{4}-\d{2}-\d{2})`

func JSON(filePath string, data interface{}) error {
	b, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalln(err)
	}

	err = json.Unmarshal(b, data)
	if err != nil {
		log.Fatalln(err)
	}

	return nil
}

func OrderedList(rePattern string, limit int, target string) ([]string, error) {
	dateRe := regexp.MustCompile(datePattern)
	re := regexp.MustCompile(rePattern)

	files, err := ioutil.ReadDir(target)
	if err != nil {
		return nil, err
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if re.MatchString(file.Name()) {
			paths = append(paths, filepath.Join(target, file.Name()))
		}
	}

	sort.Slice(paths, func(i, j int) bool {
		t1, err := time.Parse("2006-01-02", dateRe.FindStringSubmatch(paths[i])[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		t2, err := time.Parse("2006-01-02", dateRe.FindStringSubmatch(paths[j])[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		return t1.After(t2)
	})

	if limit > 0 && limit <= len(paths) {
		paths = paths[:limit]
	}

	return paths, nil
}
