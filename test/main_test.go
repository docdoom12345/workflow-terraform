package test

import (
	"testing"
	"encoding/json"
	"io/ioutil"
	"strings"
	"fmt"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformVMNames(t *testing.T) {
	t.Parallel()

	// Path to your Terraform code directory
	terraformOptions := &terraform.Options{
		TerraformDir: ".",
	}

	// Run 'terraform init' to initialize the Terraform configuration
	terraform.Init(t, terraformOptions)

	// Run 'terraform plan' to create the plan
	planFilePath := "test.tfplan"
	terraform.Plan(t, terraformOptions, planFilePath)

	// Parse the Terraform plan file and extract the Azure VM names
	vmNames, err := extractAzureVMNamesFromPlan(planFilePath)
	if err != nil {
		t.Fatalf("Failed to extract Azure VM names from plan: %v", err)
	}

	// Set your expected Azure VM names here
	expectedVMNames := []string{"example-machine"}

	// Check if the Azure VM names match the expected values
	assert.ElementsMatch(t, expectedVMNames, vmNames, "Azure VM names should match expected values")
}

func extractAzureVMNamesFromPlan(planFilePath string) ([]string, error) {
	// Read the contents of the plan file
	planJSON, err := ioutil.ReadFile(planFilePath)
	if err != nil {
		return nil, err
	}

	// Parse the JSON content to a Terraform Plan object
	plan := &terraform.Plan{}
	if err := json.Unmarshal(planJSON, plan); err != nil {
		return nil, err
	}

	// Find the resources "azurerm_virtual_machine" in the plan and extract their names
	var vmNames []string
	for _, resourceChange := range plan.ResourceChanges {
		if resourceChange.Type == "azurerm_windows_virtual_machine" {
			// Get the name attribute of the resource change
			if nameVal, ok := resourceChange.Change.After["name"]; ok {
				vmName := nameVal.(string)
				vmNames = append(vmNames, vmName)
			}
		}
	}

	if len(vmNames) == 0 {
		// If "azurerm_virtual_machine" resource is not found, return an error
		return nil, fmt.Errorf("azurerm_virtual_machine resource not found in the plan")
	}

	return vmNames, nil
}
