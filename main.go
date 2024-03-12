package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// Get the path from the environment variable
	path := os.Getenv("PLUGIN_PATH")
	// path := "./opt" // For testing purposes only!

	// Check if the path is provided
	if path == "" {
		fmt.Println("Error: No path provided")
		os.Exit(1)
	}

	// Change the current working directory
	err := os.Chdir(path)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Read the contents of the directory
	files, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Print the names of the files
	for _, file := range files {
		fmt.Println(file.Name())
	}
}
