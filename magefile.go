//go:build mage
// +build mage

package main

import (
	"fmt"
	"os"

	//	"github.com/magefile/mage/mg" // mg contains helpful utility functions, like Deps
	"github.com/magefile/mage/sh"
)

var Default = Test

// Dep ensures that we've got our dependencies
func Dep() error {
	fmt.Println("Installing Deps...")
	if err := sh.RunV("go", "mod", "download"); err != nil {
		return err
	}

	return nil
}

// Test runs the tests on this repository
func Test() error {
	if os.Getenv("CI") == "" {
		if err := sh.Run("golangci-lint", "run"); err != nil {
			return err
		}
	}

	return TestFast()
}

// TestFast runs the tests on this repository without golangci-lint, for fast cycles
func TestFast() error {
	return sh.RunV("go", "test", "-coverpkg=./...", "-coverprofile=c.out", "./...")
}

// Clean cleans up all build artifacts
func Clean() {
	fmt.Println("Cleaning...")
	os.RemoveAll("bin")
}
