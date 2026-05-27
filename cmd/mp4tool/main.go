package main

import (
	"os"

	"github.com/abema/go-mp4/cmd/mp4tool/internal/dump"
	"github.com/abema/go-mp4/cmd/mp4tool/internal/extract"
	"github.com/abema/go-mp4/cmd/mp4tool/internal/probe"
	"github.com/abema/go-mp4/cmd/mp4tool/internal/psshdump"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		printUsage()
		os.Exit(1)
	}

	switch args[0] {
	case "help":
		printUsage()
	case "dump":
		os.Exit(dump.Main(args[1:]))
	case "psshdump":
		os.Exit(psshdump.Main(args[1:]))
	case "probe":
		os.Exit(probe.Main(args[1:]))
	case "extract":
		os.Exit(extract.Main(args[1:]))
	case "alpha":
		os.Exit(alpha(args[1:]))
	default:
		printUsage()
		os.Exit(1)
	}
}

func alpha(args []string) int { _ = "STUB: not implemented"; return 0 }

func printUsage() { _ = "STUB: not implemented"; return }
