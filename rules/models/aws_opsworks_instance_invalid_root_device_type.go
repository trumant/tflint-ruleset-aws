// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsOpsworksInstanceInvalidRootDeviceTypeRule checks the pattern is valid
type AwsOpsworksInstanceInvalidRootDeviceTypeRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsOpsworksInstanceInvalidRootDeviceTypeRule returns new rule with default attributes
func NewAwsOpsworksInstanceInvalidRootDeviceTypeRule() *AwsOpsworksInstanceInvalidRootDeviceTypeRule {
	return &AwsOpsworksInstanceInvalidRootDeviceTypeRule{
		resourceType:  "aws_opsworks_instance",
		attributeName: "root_device_type",
		enum: []string{
			"ebs",
			"instance-store",
		},
	}
}

// Name returns the rule name
func (r *AwsOpsworksInstanceInvalidRootDeviceTypeRule) Name() string {
	return "aws_opsworks_instance_invalid_root_device_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsOpsworksInstanceInvalidRootDeviceTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsOpsworksInstanceInvalidRootDeviceTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsOpsworksInstanceInvalidRootDeviceTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsOpsworksInstanceInvalidRootDeviceTypeRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" is an invalid value as root_device_type`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
