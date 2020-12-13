// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsSsmAssociationInvalidMaxErrorsRule checks the pattern is valid
type AwsSsmAssociationInvalidMaxErrorsRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsSsmAssociationInvalidMaxErrorsRule returns new rule with default attributes
func NewAwsSsmAssociationInvalidMaxErrorsRule() *AwsSsmAssociationInvalidMaxErrorsRule {
	return &AwsSsmAssociationInvalidMaxErrorsRule{
		resourceType:  "aws_ssm_association",
		attributeName: "max_errors",
		max:           7,
		min:           1,
		pattern:       regexp.MustCompile(`^([1-9][0-9]*|[0]|[1-9][0-9]%|[0-9]%|100%)$`),
	}
}

// Name returns the rule name
func (r *AwsSsmAssociationInvalidMaxErrorsRule) Name() string {
	return "aws_ssm_association_invalid_max_errors"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSsmAssociationInvalidMaxErrorsRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSsmAssociationInvalidMaxErrorsRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSsmAssociationInvalidMaxErrorsRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSsmAssociationInvalidMaxErrorsRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"max_errors must be 7 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"max_errors must be 1 characters or higher",
					attribute.Expr,
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^([1-9][0-9]*|[0]|[1-9][0-9]%|[0-9]%|100%)$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
