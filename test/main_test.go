package test

import (
	"testing"
	//"os/exec"
	//"bufio"
	"os"
	"fmt"
	//"strings"
	"io/ioutil"

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
	content, err :=ioutil.ReadFile(PlanFileName")
	if err != nil {
            fmt.Println("Unable to read file")
	    os.Exit(1)
        }
	fmt.Println(string(content))
}
