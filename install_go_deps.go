package magefiles

import "github.com/magefile/mage/sh"

// Install iterates through provided dependency urls and go installs them
func Install(deps []string) error {
	for _, dep := range deps {
		err := sh.RunV("go", "install", dep)
		if err != nil {
			return err
		}
	}
	return nil
}
