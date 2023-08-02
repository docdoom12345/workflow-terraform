package test

import (
	"testing"
	"io/ioutil"
	"os"
	"fmt"
)

func TestVMNameInTerraformPlan(t *testing.T) {
        cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting the current working directory:", err)
		return
	}

	// Read the contents of the current directory.
	files, err := ioutil.ReadDir(cwd)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	// Loop through the list of files and print their names.
	fmt.Println("Files in the current directory:")
	for _, file := range files {
		// Check if it is a regular file (not a directory).
		if file.Mode().IsRegular() {
			fmt.Println(file.Name())
		}
	}
}
