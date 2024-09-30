package main

import (
	"fmt"
	"github.com/robotjoosen/go-config"
)

type configuration struct {
	Name string `mapstructure:"NAME"`
}

func Example_Usage() {
	var cnf configuration
	if _, err := config.Load(&cnf, map[string]any{
		"NAME": "unknown",
	}); err != nil {
		return
	}

	fmt.Println(cnf.Name)

	// output:
	// unknown
}
