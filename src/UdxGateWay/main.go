package main

import (
	"flag"
)

var cfgPath string

func init() {
	flag.StringVar(&cfgPath, "dll", "dll", "default path for configuration files")
}

func main() {

}
