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
	path, ok := filepath(filename)
	if !ok {
		log.Fatal("Cannot find Config file: ", path)
	}

	raw, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("Cannot open Config file: ", path)
	}

	json.Unmarshal(raw, &f.v)
}

func filepath(filename string) (string, bool) {
	base := path.Base(filename)

	current := path.Join(".", base)
	if exists(current) {
		return current, true
	}

	home := path.Join(os.Getenv("HOME"), base)
	if exists(home) {
		return home, true
	}

	return filename, false
}

func exists(filepath string) bool {
	_, err := os.Stat(filepath)
	return err == nil
}
