package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/namsral/flag"
	"gopkg.in/yaml.v2"
)

type JsonConfig map[string]string

var (
	version        = "no version set"
	commit         = "no commit"
	date           = "no date"
	displayVersion bool
	inputFileName  string
	outputFormat   string
)

func main() {
	flag.StringVar(&inputFileName, "input", "", "path to the input JSON file, read from stdin if none provided")
	flag.StringVar(&outputFormat, "output", "export", "output format, one of 'export' (default) or 'vars'")
	flag.BoolVar(&displayVersion, "version", false, "Show version and quit")
	flag.Parse()

	if displayVersion {
		fmt.Printf("%s (%s), build on %s\n\n", version, commit, date)

		flag.Usage()
		os.Exit(0)
	}

	jsonConfig := JsonConfig{}
	inputReader := os.Stdin

	if len(inputFileName) > 0 {
		var err error
		inputReader, err = os.Open(inputFileName)
		if err != nil {
			fmt.Println(err)

			return
		}
		defer inputReader.Close()
	}

	data, err := ioutil.ReadAll(inputReader)
	if err != nil {
		fmt.Println(err)

		return
	}
	err = yaml.Unmarshal(data, &jsonConfig)
	if err != nil {
		fmt.Println(err)

		return
	}

	switch outputFormat {
	case "export":
		printExport(jsonConfig)
	case "vars":
		printVars(jsonConfig)
	}
}

func printExport(jsonConfig JsonConfig) {
	for envKey, envVal := range jsonConfig {
		fmt.Printf("export %s=\"%s\"\n", envKey, envVal)
	}
}

func printVars(jsonConfig JsonConfig) {
	for envKey, envVal := range jsonConfig {
		fmt.Printf("%s=\"%s\"\n", envKey, envVal)
	}
}
