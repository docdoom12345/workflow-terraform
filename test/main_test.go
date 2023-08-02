package test

import (
	"testing"
	"os/exec"
	"os"
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

	// Run `terraform init` and `terraform plan` to generate the plan.
	terraform.InitAndApply(t, terraformOptions)

	// Get the plan using `terraform plan -out` command.
	planFile, err := ioutil.TempFile("", "terraform-plan")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(planFile.Name())

	cmd := exec.Command("terraform", "plan", "-out", planFile.Name())
	cmd.Dir = terraformOptions.TerraformDir
	err = cmd.Run()
	if err != nil {
		t.Fatalf("Failed to run terraform plan command: %v", err)
	}

	// At this point, the plan has been saved to a temporary file. You can process or assert on the plan as needed.
	// For example, you can copy the plan file to a specific location or compare it against an expected plan.
	// For simplicity, we'll just print the plan file path here:
	t.Logf("Terraform plan saved to: %s", planFile.Name())
}
