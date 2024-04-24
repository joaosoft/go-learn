package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
)

func main() {
	forLoop()
	if err := seeTestCoverage(); err != nil {
		panic(err)
	}
}

func forLoop() {
	// now you can make a range from an integer
	for i := range 10 {
		fmt.Println("index", i)
	}
	fmt.Println("go1.22 has lift-off!")
}

func seeTestCoverage() error {
	workingDirectory, err := os.Getwd()
	if err != nil {
		return err
	}

	cmd := exec.Command("go", "test", "./...", "-cover")
	cmd.Dir = path.Join(workingDirectory, "golang-learn")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	if err = cmd.Run(); err != nil {
		return err
	}

	return nil
}
