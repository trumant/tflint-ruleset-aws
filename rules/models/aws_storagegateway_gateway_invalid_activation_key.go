// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsStoragegatewayGatewayInvalidActivationKeyRule checks the pattern is valid
type AwsStoragegatewayGatewayInvalidActivationKeyRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsStoragegatewayGatewayInvalidActivationKeyRule returns new rule with default attributes
func NewAwsStoragegatewayGatewayInvalidActivationKeyRule() *AwsStoragegatewayGatewayInvalidActivationKeyRule {
	return &AwsStoragegatewayGatewayInvalidActivationKeyRule{
		resourceType:  "aws_storagegateway_gateway",
		attributeName: "activation_key",
		max:           50,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsStoragegatewayGatewayInvalidActivationKeyRule) Name() string {
	return "aws_storagegateway_gateway_invalid_activation_key"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsStoragegatewayGatewayInvalidActivationKeyRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsStoragegatewayGatewayInvalidActivationKeyRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsStoragegatewayGatewayInvalidActivationKeyRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsStoragegatewayGatewayInvalidActivationKeyRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"activation_key must be 50 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"activation_key must be 1 characters or higher",
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
