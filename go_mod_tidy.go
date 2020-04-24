package magefiles

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/magefile/mage/sh"
)

// GoModTidy will run go mod tidy and check if there are any pending changes in git. Will fail if changes are pending.
func GoModTidy() error {
	err := sh.RunV("go", "mod", "tidy", "-v")
	if err != nil {
		return err
	}

	gitRepo, err := git.PlainOpen(".")

	if err != nil {
		return err
	}

	worktree, err := gitRepo.Worktree()

	if err != nil {
		return err
	}

	status, err := worktree.Status()

	if err != nil {
		return err
	}

	if !status.IsClean() {

		return fmt.Errorf("git status is not clean: \n%s", status.String())
	}

	return nil
}
