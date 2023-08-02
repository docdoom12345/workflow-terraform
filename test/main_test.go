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
}
