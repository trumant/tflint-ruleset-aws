// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsEc2FleetInvalidTypeRule checks the pattern is valid
type AwsEc2FleetInvalidTypeRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsEc2FleetInvalidTypeRule returns new rule with default attributes
func NewAwsEc2FleetInvalidTypeRule() *AwsEc2FleetInvalidTypeRule {
	return &AwsEc2FleetInvalidTypeRule{
		resourceType:  "aws_ec2_fleet",
		attributeName: "type",
		enum: []string{
			"request",
			"maintain",
			"instant",
		},
	}
}

// Name returns the rule name
func (r *AwsEc2FleetInvalidTypeRule) Name() string {
	return "aws_ec2_fleet_invalid_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsEc2FleetInvalidTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsEc2FleetInvalidTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsEc2FleetInvalidTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsEc2FleetInvalidTypeRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as type`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
