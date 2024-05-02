package main

import (
	"fmt"
	"ti/cmd"
	"ti/conf"
)

func main() {
	file, cmd, args := cmd.Execute()
	config, err := conf.Parse(file)
	if err != nil {
		panic(err)
	}

	fmt.Println(config, cmd, args)
}
