package cmd

import (
	"flag"
	"os"
	"path/filepath"
	"strconv"
)

var upDir string

func init() {
	path, err := os.Executable()
	if err != nil {
		panic(err)
	}

	upDir = filepath.Dir(path)
}

func Execute() (string, string, []string) {
	configPtr := flag.String("c", filepath.Join(upDir, "ti.toml"), "config file")
	flag.Lookup("c").DefValue = "执行文件所在目录/ti.toml"
	flag.Parse()

	args := flag.Args()
	var cmd string
	switch len(args) {
	case 0:
		cmd = "info"
	case 1:
		if _, err := strconv.Atoi(args[0]); err == nil {
			cmd = "get"
		} else {
			cmd = "find"
		}
	default:
		cmd = "find"
	}

	return *configPtr, cmd, args
}
