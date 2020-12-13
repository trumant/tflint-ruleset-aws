// This file generated by `generator/main.go`. DO NOT EDIT

package api

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
    "github.com/terraform-linters/tflint-ruleset-aws/aws"
)

// AwsLaunchConfigurationInvalidIAMProfileRule checks whether attribute value actually exists
type AwsLaunchConfigurationInvalidIAMProfileRule struct {
	resourceType  string
	attributeName string
	data          map[string]bool
	dataPrepared  bool
}

// NewAwsLaunchConfigurationInvalidIAMProfileRule returns new rule with default attributes
func NewAwsLaunchConfigurationInvalidIAMProfileRule() *AwsLaunchConfigurationInvalidIAMProfileRule {
	return &AwsLaunchConfigurationInvalidIAMProfileRule{
		resourceType:  "aws_launch_configuration",
		attributeName: "iam_instance_profile",
		data:          map[string]bool{},
		dataPrepared:  false,
	}
}

// Name returns the rule name
func (r *AwsLaunchConfigurationInvalidIAMProfileRule) Name() string {
	return "aws_launch_configuration_invalid_iam_profile"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsLaunchConfigurationInvalidIAMProfileRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsLaunchConfigurationInvalidIAMProfileRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsLaunchConfigurationInvalidIAMProfileRule) Link() string {
	return ""
}

// Check checks whether the attributes are included in the list retrieved by ListInstanceProfiles
func (r *AwsLaunchConfigurationInvalidIAMProfileRule) Check(rr tflint.Runner) error {
    runner := rr.(*aws.Runner)

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		if !r.dataPrepared {
			log.Print("[DEBUG] invoking ListInstanceProfiles")
			var err error
			r.data, err = runner.AwsClient.ListInstanceProfiles()
			if err != nil {
				err := &tflint.Error{
					Code:    tflint.ExternalAPIError,
					Level:   tflint.ErrorLevel,
					Message: "An error occurred while invoking ListInstanceProfiles",
					Cause:   err,
				}
				log.Printf("[ERROR] %s", err)
				return err
			}
			r.dataPrepared = true
		}

		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if !r.data[val] {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" is invalid IAM profile name.`, val),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
