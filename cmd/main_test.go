package main_test

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/alecthomas/assert/v2"
)

var (
	binName = "temp"
)

func TestMain(m *testing.M) {
	fmt.Println("Building tool...")

	build := exec.Command("go", "build", "-o", binName)

	if err := build.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Cannot build tool %s: %s", binName, err)
		os.Exit(1)
	}

	fmt.Println("Running tests....")
	result := m.Run()

	fmt.Println("Cleaning up...")
	os.Remove(binName)

	os.Exit(result)
}

func TestTempConverter(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	cmdPath := filepath.Join(dir, binName)

	t.Run("Converts Fahrenheit to Celcius", func(t *testing.T) {
		cmd := exec.Command(cmdPath)

		cmdStdIn, err := cmd.StdinPipe()
		assert.NoError(t, err)

		io.WriteString(cmdStdIn, "C ")
		io.WriteString(cmdStdIn, "32")
		cmdStdIn.Close()

		out, err := cmd.CombinedOutput()
		fahrenheit := string(out)

		assert.Contains(t, fahrenheit, "64")
	})
}

//Press C to convert from Fahrenheit to Celsius.
//Press F to convert from Celsius to Fahrenheit.
//Your choice: C
//
//Enter the temperature in Fahrenheit: 32
//The temperature in Celsius is: 0
