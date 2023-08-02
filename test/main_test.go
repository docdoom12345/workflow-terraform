package test

import (
	"testing"
	"os/exec"
	//"os"
	//"strings"
	//"io/ioutil"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestTerraformPlanToFile(t *testing.T) {
	t.Parallel()

	// Set up Terraform options
	terraformOptions := &terraform.Options{
		// Set the path to your Terraform code that will be tested.
		TerraformDir: "../",
	}

	// Run `terraform init` and `terraform plan` to generate the plan.
	terraform.Init(t, terraformOptions)

	// Get the plan using `terraform plan -out` command.
	planFilePath := "../terraform.plan"
	cmd := exec.Command("terraform", "plan", "-out", planFilePath)
	cmd.Dir = terraformOptions.TerraformDir
	err := cmd.Run()
	if err != nil {
		t.Fatalf("Failed to run terraform plan command: %v", err)
	}
	vmName := ""
	scanner := bufio.NewScanner(planFile)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "azurerm_virtual_machine.example_vm") {
			// Assuming "azurerm_virtual_machine.example_vm" is the resource name in your Terraform configuration.
			// Adjust it accordingly based on your actual resource name.
			// The VM name will be in the line following the resource name.
			scanner.Scan()
			vmName = scanner.Text()
			break
		}
	}

	if err := scanner.Err(); err != nil {
		t.Fatalf("Error while scanning plan file: %v", err)
	}

	// At this point, the VM name should be stored in the "vmName" variable.
	t.Logf("VM Name: %s", vmName)
}
