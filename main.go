package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/namsral/flag"
	"gopkg.in/yaml.v2"
)

type Configuration map[string]string

func main() {
	var (
		configFile   string
		outputFormat string
	)

	flag.StringVar(&configFile, "input", "./config.json", "path to the config file")
	flag.StringVar(&outputFormat, "output", "export", "output format, one of 'export' (default) or 'vars'")
	flag.Parse()

	configuration := Configuration{}

	file, err := os.Open(configFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = yaml.Unmarshal(data, &configuration)
	if err != nil {
		fmt.Println(err)
		return
	}

	switch outputFormat {
	case "export":
		printExport(configuration)
	case "vars":
		printVars(configuration)
	}
}

func printExport(configuration Configuration) {
	for envKey, envVal := range configuration {
		fmt.Printf("export %s=\"%s\"\n", envKey, envVal)
	}
}

func printVars(configuration Configuration) {
	for envKey, envVal := range configuration {
		fmt.Printf("%s=\"%s\"\n", envKey, envVal)
	}
}
