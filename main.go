package main

import (
	"os"

	"github.com/biosugar0/postmaster/cmd"
	"github.com/biosugar0/postmaster/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	config.EnsurePath(config.P8rLogs, config.DefaultDirMod)
}
func main() {
	mod := os.O_CREATE | os.O_APPEND | os.O_WRONLY
	file, err := os.OpenFile(config.P8rLogs, mod, config.DefaultFileMod)
	if err != nil {
		panic(err)
	}
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: file})
	cmd.Execute()
}
