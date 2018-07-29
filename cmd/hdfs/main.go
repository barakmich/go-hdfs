package main

import "github.com/barakmich/go-hdfs/cmd/hdfs/cmd"

var version string

func main() {
	cmd.Execute(version)
}
