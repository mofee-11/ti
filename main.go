package main

import (
	"strconv"
	"ti/cmd"
	"ti/conf"
	"ti/db"
)

func main() {
	file, cmd, args := cmd.Execute()
	config, err := conf.Parse(file)
	if err != nil {
		panic(err)
	}

	coll := db.NewDB(config.DBPath)
	switch cmd {
	case "info":
		coll.Info()
	case "get":
		i, _ := strconv.Atoi(args[0])
		coll.Get(i)
	case "find":
		coll.Find(args)
	}
}
