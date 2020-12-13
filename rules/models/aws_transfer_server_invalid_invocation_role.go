// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsTransferServerInvalidInvocationRoleRule checks the pattern is valid
type AwsTransferServerInvalidInvocationRoleRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsTransferServerInvalidInvocationRoleRule returns new rule with default attributes
func NewAwsTransferServerInvalidInvocationRoleRule() *AwsTransferServerInvalidInvocationRoleRule {
	return &AwsTransferServerInvalidInvocationRoleRule{
		resourceType:  "aws_transfer_server",
		attributeName: "invocation_role",
		max:           2048,
		min:           20,
		pattern:       regexp.MustCompile(`^arn:.*role/.*$`),
	}
}

// Name returns the rule name
func (r *AwsTransferServerInvalidInvocationRoleRule) Name() string {
	return "aws_transfer_server_invalid_invocation_role"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsTransferServerInvalidInvocationRoleRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsTransferServerInvalidInvocationRoleRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsTransferServerInvalidInvocationRoleRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsTransferServerInvalidInvocationRoleRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"invocation_role must be 2048 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"invocation_role must be 20 characters or higher",
					attribute.Expr,
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^arn:.*role/.*$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
