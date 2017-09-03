package configFile

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path"
)

type ConfigFile struct {
	filename string
	v        interface{}
}

func NewConfigFile(filename string, v interface{}) *ConfigFile {
	r := new(ConfigFile)
	r.filename = filename
	r.v = v
	return r
}

func (f *ConfigFile) Read() {
	filename := f.filename
	pathes, ok := filepathes(filename)
	if !ok {
		log.Fatal("Cannot find Config file: ", filename)
	}

	for _, path := range pathes {
		raw, err := ioutil.ReadFile(path)
		if err != nil {
			log.Fatal("Cannot open Config file: ", path)
		}

		json.Unmarshal(raw, &f.v)
	}
}

func filepathes(filename string) ([]string, bool) {
	pathes := make([]string, 0)
	base := path.Base(filename)

	home := path.Join(os.Getenv("HOME"), base)
	if exists(home) {
		pathes = append(pathes, home)
	}

	current := path.Join(".", base)
	if exists(current) {
		pathes = append(pathes, current)
	}

	return pathes, (len(pathes) > 0)
}

func exists(filepath string) bool {
	_, err := os.Stat(filepath)
	return err == nil
}
