// magefile.go
//go:build mage
// +build mage

package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

var (
	BINARY_NAME  = "logistics-routes"
	BIN          = "./bin"
	CODE_FOLDERS = "internal/*"
)

func Build() error {
	fmt.Println("Construyendo el proyecto...")
	return runCommand("go", "build", "-o", filepath.Join(BIN, BINARY_NAME), "./...")
}

func InstallDeps() error {
	fmt.Println("Instalando las dependencias...")
	return runCommand("go", "mod", "tidy")
}

func Run() error {
	fmt.Printf("Ejecutando el programa %s...\n", BINARY_NAME)
	return runCommand(filepath.Join(BIN, BINARY_NAME))
}

func Clean() error {
	fmt.Println("Limpiando los binarios...")
	if err := os.RemoveAll(filepath.Join(BIN, BINARY_NAME)); err != nil {
		return fmt.Errorf("failed to remove binary: %v", err)
	}
	return runCommand("go", "clean", "./...")
}

func Check() error {
	fmt.Println("Comprobando sintaxis del proyecto...")
	return runCommand("gofmt", "-l", CODE_FOLDERS)
}

func runCommand(cmd string, args ...string) error {
	command := exec.Command(cmd, args...)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	return command.Run()
}
