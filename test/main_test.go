package test

import (
	"testing"
	//"os/exec"
	//"bufio"
	"//os"
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
	PlanFilePath := "../terraform.plan"
	terraform.InitandPlanWithConfig(t, terraformOptions, PlanFilePath)

	// Get the plan using `terraform plan -out` command.
	
}
