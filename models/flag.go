package models

import "flag"

var (
	Help = flag.Bool("help", false, "Show help messege")
	Dir  = flag.String("dir", "data", "Path to the data directory")
	Port = flag.String("port", "7070", "Port number")
)
