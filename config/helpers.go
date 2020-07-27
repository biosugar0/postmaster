package config

import (
	"os"
	"os/user"
	"path/filepath"

	"github.com/rs/zerolog/log"
)

const (
	// DefaultDirMod default unix perms for p8r directory.
	DefaultDirMod os.FileMode = 0755
	// DefaultFileMod default unix perms for p8r files.
	DefaultFileMod os.FileMode = 0600
)

// EnsurePath ensures a directory exist from the given path.
func EnsurePath(path string, mod os.FileMode) {
	dir := filepath.Dir(path)
	EnsureFullPath(dir, mod)
}

// EnsureFullPath ensures a directory exist from the given path.
func EnsureFullPath(path string, mod os.FileMode) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err = os.MkdirAll(path, mod); err != nil {
			log.Fatal().Msgf("Unable to create dir %q %v", path, err)
		}
	}
}

func mustP8rHome() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal().Err(err).Msg("Die on retrieving user home")
	}
	return usr.HomeDir
}

// MustP8rUser establishes current user identity or fail.
func MustP8rUser() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal().Err(err).Msg("Die on retrieving user info")
	}
	return usr.Username
}
