package test

import (
	"testing"
	"encoding/json"
	"io/ioutil"
        "os"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestVMNameInTerraformPlan(t *testing.T) {
	t.Parallel()

	// Set the Terraform options with the path to the Terraform code directory.
	terraformOptions := &terraform.Options{
		// Set the path to the Terraform code that will be tested.
		TerraformDir: ".",
		PlanFilePath: "./terraform.plan"
	}

	// Run `terraform init` and `terraform plan`. The plan output will be captured in the `planOutput` variable.
	terraform.InitAndPlan(t, terraformOptions)

	// Save the Terraform plan output to a file.

	// Read the Terraform plan file.
	planData, err := ioutil.ReadFile("terraform.tfplan")
	if err != nil {
		t.Fatalf("Error reading plan file: %v", err)
	}

	// Parse the JSON data from the plan file into a map.
	var planMap map[string]interface{}
	err = json.Unmarshal(planData, &planMap)
	if err != nil {
		t.Fatalf("Error unmarshaling JSON data: %v", err)
	}

	// Extract the VM name from the plan data.
	vmName := extractVMNameFromPlan(planMap)

	// Assert that the VM name exists and has the expected value.
	expectedVMName := "example-machine"
	assert.Equal(t, expectedVMName, vmName, "VM name does not match the expected value")
}

func extractVMNameFromPlan(planMap map[string]interface{}) string {
	// Extract the root module from the plan.
	rootModule := planMap["planned_values"].(map[string]interface{})["root_module"].(map[string]interface{})

	// Find the resources list within the root module.
	resources := rootModule["resources"].([]interface{})

	// Iterate through the resources to find the VM resource.
	for _, resource := range resources {
		resourceMap := resource.(map[string]interface{})
		resourceType := resourceMap["type"].(string)

		// Assuming the VM resource is "azurerm_windows_virtual_machine".
		if resourceType == "azurerm_windows_virtual_machine" {
			vmName := resourceMap["values"].(map[string]interface{})["name"].(string)
			return vmName
		}
	}

	// Return an empty string if the VM resource is not found.
	return ""
}
