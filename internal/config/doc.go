// Package config is responsible for handling the rules passed to the program
// as YAML files through the --rules flag.
//
// It allows to redefine which parameters the package checks should consider.
//
// If no rules are passed (as to say, no YAML files are provided), the program will stick
// to its default values, located in defuaults.go.
package config
