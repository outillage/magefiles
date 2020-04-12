package magefiles

import "github.com/magefile/mage/sh"

// Test runs go test ./... -v -race
func Test() error {
	return sh.RunV("go", "test", "./...", "-v", "-race")
}
