package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
)

// ExecuteElixirCode takes Elixir code as a string, executes it, and returns the result or an error
func ExecuteElixirCode(code string) (string, error) {
	// Create a temporary Elixir script
	//script := fmt.Sprintf("elixir  \"%s\"", code)

	// Execute the Elixir script
	cmd := exec.Command("elixir", code)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	return string(output), nil
}

// ReadElixirCodeFromFile reads the content of a file and returns it as a string
func ReadElixirCodeFromFile(filePath string) (string, error) {
	// Read the file content
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	// Return the content as a string
	return string(content), nil
}
func main() {
	//code := `IO.puts("Hello from Elixir!")`
	filePath := "hello.exs"
	// Read the Elixir code from the file
	code, err := ReadElixirCodeFromFile(filePath)
	if err != nil {
		log.Fatalf("Failed to read Elixir code: %v", err)
	}
	// Print the Elixir code
	fmt.Println("Elixir Code:")
	fmt.Println(code)
	result, err := ExecuteElixirCode(filePath)
	if err != nil {
		fmt.Printf("Error executing Elixir code: %v\n", err)
	} else {
		fmt.Printf("Result: %s\n", result)
	}
}
