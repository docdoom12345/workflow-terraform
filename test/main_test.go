// main_test.go

package test

import (
	"testing"
	"strings"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2021-07-01/compute"
	"github.com/Azure/go-autorest/autorest/azure/auth"
)

func TestAzureVMName(t *testing.T) {
	t.Parallel()

	// Variables for the Azure VM name pattern and the Terraform example module.
	const expectedVMNamePattern = "example-machine"
	const terraformModulePath = "../"

	// Terraform options to run during the test.
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: terraformModulePath,
	})

	// Terraform init, apply, and defer destroy.
	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)

	// Get the VM name from the Terraform output.
	vmName := terraform.Output(t, terraformOptions, "vm_name")

	// Create an Azure compute client to fetch the VM details.
	authorizer, err := auth.NewAuthorizerFromEnvironment()
	if err != nil {
		t.Fatalf("Failed to get Azure authorizer: %v", err)
	}

	vmClient := compute.NewVirtualMachinesClient(terraformOptions.Vars["azure_subscription_id"].(string))
	vmClient.Authorizer = authorizer

	// Get the VM details by resource group and VM name.
	resourceGroup := terraformOptions.Vars["resource_group"].(string)

	vm, err := vmClient.Get(terraformOptions.Vars["azure_vm_resource_group"].(string), resourceGroup, vmName, compute.InstanceView)
	if err != nil {
		t.Fatalf("Failed to get Azure VM details: %v", err)
	}

	// Check if the VM name matches the specified pattern.
	matched, err := MatchPattern(expectedVMNamePattern, *vm.Name)
	if err != nil {
		t.Fatalf("Failed to check if VM name matches pattern: %v", err)
	}

	// Assert the match result.
	assert.True(t, matched, "Azure VM name does not match the pattern: %s", expectedVMNamePattern)
}

// MatchPattern checks if a given string matches the specified pattern.
func MatchPattern(pattern, value string) (bool, error) {
	matched, err := path.Match(pattern, value)
	if err != nil {
		return false, err
	}
	return matched, nil
}
