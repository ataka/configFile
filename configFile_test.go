package configFile

import (
	"testing"
)

type Config struct {
	Year    int  `json:"year"`
	Verbose bool `json:"verbose"`
	Token   `json:"token"`
}

type Token struct {
	Foo string `json:"foo"`
	Bar string `json:"bar"`
}

func TestRead(t *testing.T) {
	var config Config
	configFile := NewConfigFile("sample.json", &config)
	configFile.Read()

	if config.Year != 2017 {
		t.Errorf("%d is not 2017", config.Year)
	}
	if !config.Verbose {
		t.Errorf("%t is not true", config.Verbose)
	}
	if config.Token.Foo != "abcdefg" {
		t.Errorf("%s is not \"abcdefg\"", config.Token.Foo)
	}
	if config.Bar != "0123-456-789" {
		t.Errorf("%s is not \"0123-456-789\"", config.Bar)
	}
}
