package main

import "os"

func commandExit(config *Config, commandParam string) error {
	os.Exit(0)
	return nil
}
