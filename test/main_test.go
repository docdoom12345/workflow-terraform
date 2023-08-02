package test

import (
	"testing"
	//"os/exec"
	//"bufio"
	"os"
	"fmt"
	//"strings"
	"io/ioutil"
        "path/filepath"
	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestTerraformPlanToFile(t *testing.T) {
	t.Parallel()

	// Set up Terraform options
	terraformOptions := &terraform.Options{
		// Set the path to your Terraform code that will be tested.
		TerraformDir: "../",
		PlanFilePath: "terraform.tfplan",
	}

	// Run `terraform init` and `terraform plan` to generate the plan.
	terraform.Init(t, terraformOptions)

	// Get the plan using `terraform plan -out` command.
	PlanFileName := "../terraform.tfplan"
	terraform.RunTerraformCommand(t, terraformOptions, "plan" ,"-out="+PlanFileName)
	currentDir, err := os.Getwd()
	    if err != nil {
		fmt.Println("Error getting the current directory:", err)
		os.Exit(1)
	    }
	
	    // Navigate to the parent directory
	    parentDir := filepath.Join(currentDir, "..")
	
	    // Define the file path to read
	    filePath := filepath.Join(parentDir, "terraform.tfplan")
	
	    // Read the entire content of the file
	    content, err := ioutil.ReadFile(filePath)
	    if err != nil {
		fmt.Println("Error reading the file:", err)
		os.Exit(1)
	    }
	
	    // Convert the content to a string and print it
	    fmt.Println(string(content))
}
