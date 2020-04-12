//+build mage

package main

import (
	"github.com/outillage/magefiles"
)

func Test() error {
	return magefiles.Test()
}

func GoModTidy() error {
	return magefiles.GoModTidy()
}
