// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsIAMUserSSHKeyInvalidStatusRule checks the pattern is valid
type AwsIAMUserSSHKeyInvalidStatusRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsIAMUserSSHKeyInvalidStatusRule returns new rule with default attributes
func NewAwsIAMUserSSHKeyInvalidStatusRule() *AwsIAMUserSSHKeyInvalidStatusRule {
	return &AwsIAMUserSSHKeyInvalidStatusRule{
		resourceType:  "aws_iam_user_ssh_key",
		attributeName: "status",
		enum: []string{
			"Active",
			"Inactive",
		},
	}
}

// Name returns the rule name
func (r *AwsIAMUserSSHKeyInvalidStatusRule) Name() string {
	return "aws_iam_user_ssh_key_invalid_status"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsIAMUserSSHKeyInvalidStatusRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsIAMUserSSHKeyInvalidStatusRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsIAMUserSSHKeyInvalidStatusRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsIAMUserSSHKeyInvalidStatusRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as status`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
