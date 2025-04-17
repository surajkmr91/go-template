package config

import (
	"fmt"
	"os"

	"github.com/surajkmr91/go-template/commons/flags"

	"gopkg.in/yaml.v3"
)

type Config struct {
	env         map[string]string
	application map[string]interface{}
}

var config *Config

func InitConfig(directory, appConfigName string) error {
	config = new(Config)
	fileName := getFileName(directory, appConfigName)
	err := FromYamlFile(fileName, &config.application)
	if err != nil {
		return err
	}
	return nil
}

func getFileName(directory, configName string) string {
	filePathWithName := fmt.Sprintf("%s/%s/%s", directory, flags.Env(), configName)
	return filePathWithName
}

func FromYamlFile(path string, v interface{}) error {
	// read YAML text file into a string
	yml, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	// unmarshal from string to struct
	return FromYaml(yml, v)
}

// FromYaml extracts settings from a YAML string.
func FromYaml(yml []byte, v interface{}) error {
	return yaml.Unmarshal(yml, v)
}

// // Loads config information from a JSON file
// func LoadConfig(filename string) *Config {
// 	result := newConfig()
// 	result.filename = filename
// 	err := result.parse()
// 	if err != nil {
// 		log.Fatalf("error loading config file %s: %s", filename, err)
// 	}
// 	return result
// }

// // Loads config information from a JSON string
// func LoadConfigString(s string) *Config {
// 	result := newConfig()
// 	err := json.Unmarshal([]byte(s), &result.application)
// 	if err != nil {
// 		log.Fatalf("error parsing config string %s: %s", s, err)
// 	}
// 	return result
// }

// func (c *Config) StringMerge(s string) {
// 	next := LoadConfigString(s)
// 	c.merge(next.application)
// }

// func (c *Config) LoadMerge(filename string) {
// 	next := LoadConfig(filename)
// 	c.merge(next.application)
// }

// func (c *Config) merge(napplication map[string]interface{}) {
// 	for k, v := range napplication {
// 		c.application[k] = v
// 	}
// }

// func (c *Config) parse() error {
// 	f, err := os.Open(c.filename)
// 	if err != nil {
// 		return err
// 	}
// 	defer f.Close()
// 	b := new(bytes.Buffer)
// 	_, err = b.ReadFrom(f)
// 	if err != nil {
// 		return err
// 	}
// 	err = json.Unmarshal(b.Bytes(), &c.application)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// Returns a string for the config variable key
func (c *Config) GetString(key string) string {
	result, present := c.application[key]
	if !present {
		return ""
	}
	return result.(string)
}

// Returns an int for the config variable key
func (c *Config) GetInt(key string) int {
	x, ok := c.application[key]
	if !ok {
		return -1
	}
	return int(x.(float64))
}

// Returns a float for the config variable key
func (c *Config) GetFloat(key string) float64 {
	x, ok := c.application[key]
	if !ok {
		return -1
	}
	return x.(float64)
}

// Returns a bool for the config variable key
func (c *Config) GetBool(key string) bool {
	x, ok := c.application[key]
	if !ok {
		return false
	}
	return x.(bool)
}

// Returns an array for the config variable key
func (c *Config) GetArray(key string) []interface{} {
	result, present := c.application[key]
	if !present {
		return []interface{}(nil)
	}
	return result.([]interface{})
}

// Returns an map for the config variable key
func (c *Config) GetMap(key string) map[string]interface{} {
	result, present := c.application[key]
	if !present {
		return map[string]interface{}(nil)
	}
	return result.(map[string]interface{})
}

func Deafult() *Config {
	return config
}
