// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsCognitoIdentityProviderInvalidUserPoolIDRule checks the pattern is valid
type AwsCognitoIdentityProviderInvalidUserPoolIDRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsCognitoIdentityProviderInvalidUserPoolIDRule returns new rule with default attributes
func NewAwsCognitoIdentityProviderInvalidUserPoolIDRule() *AwsCognitoIdentityProviderInvalidUserPoolIDRule {
	return &AwsCognitoIdentityProviderInvalidUserPoolIDRule{
		resourceType:  "aws_cognito_identity_provider",
		attributeName: "user_pool_id",
		max:           55,
		min:           1,
		pattern:       regexp.MustCompile(`^[\w-]+_[0-9a-zA-Z]+$`),
	}
}

// Name returns the rule name
func (r *AwsCognitoIdentityProviderInvalidUserPoolIDRule) Name() string {
	return "aws_cognito_identity_provider_invalid_user_pool_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCognitoIdentityProviderInvalidUserPoolIDRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCognitoIdentityProviderInvalidUserPoolIDRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCognitoIdentityProviderInvalidUserPoolIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCognitoIdentityProviderInvalidUserPoolIDRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"user_pool_id must be 55 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"user_pool_id must be 1 characters or higher",
					attribute.Expr,
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[\w-]+_[0-9a-zA-Z]+$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
