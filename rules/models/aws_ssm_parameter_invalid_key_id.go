// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsSsmParameterInvalidKeyIDRule checks the pattern is valid
type AwsSsmParameterInvalidKeyIDRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsSsmParameterInvalidKeyIDRule returns new rule with default attributes
func NewAwsSsmParameterInvalidKeyIDRule() *AwsSsmParameterInvalidKeyIDRule {
	return &AwsSsmParameterInvalidKeyIDRule{
		resourceType:  "aws_ssm_parameter",
		attributeName: "key_id",
		max:           256,
		min:           1,
		pattern:       regexp.MustCompile(`^([a-zA-Z0-9:/_-]+)$`),
	}
}

// Name returns the rule name
func (r *AwsSsmParameterInvalidKeyIDRule) Name() string {
	return "aws_ssm_parameter_invalid_key_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSsmParameterInvalidKeyIDRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSsmParameterInvalidKeyIDRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSsmParameterInvalidKeyIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSsmParameterInvalidKeyIDRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"key_id must be 256 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"key_id must be 1 characters or higher",
					attribute.Expr,
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^([a-zA-Z0-9:/_-]+)$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
