package test

import (
	"testing"
	"path/filepath"
	"io/ioutil"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestTerraformPlanToFile(t *testing.T) {
	t.Parallel()

	// Set up Terraform options
	terraformOptions := &terraform.Options{
		// Set the path to your Terraform code that will be tested.
		TerraformDir: "../",
	}

	// Run `terraform init` and `terraform plan` to generate the plan
	parentDir := ".."
	filePath := filepath.Join(parentDir,"terraform.plan")
	plan, err := terraform.InitandPlan(t, terraformOptions)
	
        }
