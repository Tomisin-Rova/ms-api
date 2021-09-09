//go:build mage
// +build mage

package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"

	genSchema "github.com/99designs/gqlgen/cmd"
)

// Default target to run when none is specified
// If not set, running mage will list available targets
var Default = GenProto

// A build step that requires additional params, or platform specific steps for example

// Manage your deps, or running package managers.
func GenProto() error {
	matches, _ := filepath.Glob("protos/*.proto")
	wd, _ := os.Getwd()
	p := path.Clean(fmt.Sprintf("%s/protos/", wd))
	bin, _ := exec.Command("which", "protoc").Output()
	protoc := strings.TrimSpace(string(bin))
	for _, filePath := range matches {

		shards := strings.Split(filePath, "/")
		fileName := shards[len(shards)-1]
		folder := fmt.Sprintf("%s/..", wd)

		if _, err := os.Stat(folder); os.IsNotExist(err) {
			_ = os.MkdirAll(folder, os.ModePerm|os.ModeDir)
		}
		commandString := fmt.Sprintf("%s -I%s/ --go_out=plugins=grpc:%s %s/%s", protoc, p, folder, p, fileName)
		println(commandString)
		if _, err := exec.Command("/bin/bash", "-c", commandString).Output(); err != nil {
			fmt.Print(err)
			return err
		}
	}

	return nil
}

func GenSchema() error {
	genSchema.Execute()
	return nil
}
