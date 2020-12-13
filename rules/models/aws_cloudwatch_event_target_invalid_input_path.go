// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsCloudwatchEventTargetInvalidInputPathRule checks the pattern is valid
type AwsCloudwatchEventTargetInvalidInputPathRule struct {
	resourceType  string
	attributeName string
	max           int
}

// NewAwsCloudwatchEventTargetInvalidInputPathRule returns new rule with default attributes
func NewAwsCloudwatchEventTargetInvalidInputPathRule() *AwsCloudwatchEventTargetInvalidInputPathRule {
	return &AwsCloudwatchEventTargetInvalidInputPathRule{
		resourceType:  "aws_cloudwatch_event_target",
		attributeName: "input_path",
		max:           256,
	}
}

// Name returns the rule name
func (r *AwsCloudwatchEventTargetInvalidInputPathRule) Name() string {
	return "aws_cloudwatch_event_target_invalid_input_path"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCloudwatchEventTargetInvalidInputPathRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCloudwatchEventTargetInvalidInputPathRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCloudwatchEventTargetInvalidInputPathRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCloudwatchEventTargetInvalidInputPathRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"input_path must be 256 characters or less",
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
