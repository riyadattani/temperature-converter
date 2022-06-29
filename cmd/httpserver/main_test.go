package main_test

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"

	"temperature-converter/pkg/temphttp"

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

func TestHttp_ConvertsTemperatures(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	cmdPath := filepath.Join(dir, binName)
	cmd := exec.Command(cmdPath)
	assert.NoError(t, cmd.Start())

	waitForServer()
	client := temphttp.NewClient("http://localhost:8080", &http.Client{})

	t.Run("Converts Celsius to Fahrenheit", func(t *testing.T) {
		temp, err := client.Convert("F", 32)
		assert.NoError(t, err)
		assert.Equal(t, "64", temp)
	})

	t.Run("Converts Fahrenheit to Celsius", func(t *testing.T) {
		temp, err := client.Convert("C", 32)
		assert.NoError(t, err)
		assert.Equal(t, "0", temp)
	})
}

func waitForServer() {
	for i := 0; i < 10; i++ {
		conn, _ := net.Dial("tcp", net.JoinHostPort("localhost", "8080"))
		if conn != nil {
			conn.Close()
			break
		}
		time.Sleep(100 * time.Millisecond)
	}
}
