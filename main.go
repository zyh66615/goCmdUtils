package main

import (
	"goCmdUtils/cmd"
	"log"
)

func main() {
	log.SetFlags(log.Ltime | log.Lshortfile)
	cmd.ParseFromArgs()
}
