package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

const (
	PairFile = ".pairs"
)

type Pairs struct {
	Authors map[string]string `yaml:"pairs"`
	Email   struct {
		Prefix       string `yaml:"prefix"`
		NoSoloPrefix bool   `yaml:"no_solo_prefix"`
		Global       bool   `yaml:"global"`
	}
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Pass in name of pairs seperated by space, ie: 'ij sa' or 'sa'")
		os.Exit(1)
	}

	pairs := os.Args[1:]

	raw, err := ioutil.ReadFile(PairFile)
	if err != nil {
		fmt.Println("File: '%s' not found.")
		os.Exit(1)
	}

	var p Pairs
	err = yaml.Unmarshal(raw, &p)
	if err != nil {
		fmt.Println("Err parsing yaml: '%s'", err)
		os.Exit(1)
	}

	var (
		nameOutput  []string
		emailOutput []string
	)

	for k, v := range p.Authors {
		for _, a := range pairs {
			if k == a {
				split := strings.Split(v, ";")
				nameOutput = append(nameOutput, strings.TrimSpace(split[0]))
				emailOutput = append(emailOutput, strings.TrimSpace(split[1]))
			}
		}
	}

	if len(nameOutput) != len(emailOutput) || len(nameOutput) != len(pairs) {
		fmt.Println("One or more pairs not found")
		os.Exit(1)
	}

	var output string
	output += strings.Join(nameOutput, " and ")
	output += " "

	if len(emailOutput) > 1 {
		output += fmt.Sprintf("<pair+%s@>", strings.Join(emailOutput, "+"))
	} else {
		output += emailOutput[0]
	}

	fmt.Println(output)
}
