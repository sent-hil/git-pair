package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
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
		names  []string
		emails []string
	)

	for k, v := range p.Authors {
		for _, a := range pairs {
			if k == a {
				split := strings.Split(v, ";")
				names = append(names, strings.TrimSpace(split[0]))
				emails = append(emails, strings.TrimSpace(split[1]))
			}
		}
	}

	if len(names) != len(emails) || len(names) != len(pairs) {
		fmt.Println("One or more pairs not found")
		os.Exit(1)
	}

	nameOutput := strings.Join(names, " and ")
	var emailOutput string

	if len(emails) > 1 {
		emailOutput = fmt.Sprintf("<pair+%s@>", strings.Join(emails, "+"))
	} else {
		emailOutput += emails[0]
	}

	if err = set_git_config("user.email", emailOutput); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err = set_git_config("user.name", nameOutput); err != nil {
		os.Exit(1)
	}
}

func set_git_config(key, value string) error {
	args := []string{"config", key, value}
	_, err := exec.Command("git", args...).CombinedOutput()
	return err
}
