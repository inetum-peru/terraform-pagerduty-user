package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestBasicSuccess(t *testing.T) {
	t.Parallel()

	name := "test user"
	email := "luis@yopmail.com"
	mobile := "959196850"

	terraformOptions := &terraform.Options{
		// The path to where your Terraform code is located
		TerraformDir: "pagerduty-user-basic",
		Upgrade:      true,
		Vars: map[string]interface{}{
			"name":   name,
			"email":  email,
			"mobile": mobile,
		},
	}

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// This will run `terraform init` and `terraform apply` and fail the test if there are any errors
	terraform.InitAndApply(t, terraformOptions)
	outputInstance := terraform.Output(t, terraformOptions, "instance")
	assert.NotEmpty(t, outputInstance, outputInstance)
}
