package magefiles

import "github.com/magefile/mage/sh"

// GoModTidy will run go mod tidy and check if there are any pending changes in git. Will fail if changes are pending.
func GoModTidy() error {
	err := sh.RunV("go", "mod", "tidy", "-v")
	if err != nil {
		return err
	}
	return sh.RunV("git", "diff-index", "--quiet", "HEAD")
}
