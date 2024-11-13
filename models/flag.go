package models

import (
	"flag"
	"log/slog"
)

var (
	Help = flag.Bool("help", false, "Show help messege")
	Dir  = flag.String("dir", "data", "Path to the data directory")
	Port = flag.String("port", "7070", "Port number")

	TotalPrice  = 0.0
	PopularItem = map[string]int{}

	Logger *slog.Logger
)
