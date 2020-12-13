// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsCodebuildSourceCredentialInvalidServerTypeRule checks the pattern is valid
type AwsCodebuildSourceCredentialInvalidServerTypeRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsCodebuildSourceCredentialInvalidServerTypeRule returns new rule with default attributes
func NewAwsCodebuildSourceCredentialInvalidServerTypeRule() *AwsCodebuildSourceCredentialInvalidServerTypeRule {
	return &AwsCodebuildSourceCredentialInvalidServerTypeRule{
		resourceType:  "aws_codebuild_source_credential",
		attributeName: "server_type",
		enum: []string{
			"GITHUB",
			"BITBUCKET",
			"GITHUB_ENTERPRISE",
		},
	}
}

// Name returns the rule name
func (r *AwsCodebuildSourceCredentialInvalidServerTypeRule) Name() string {
	return "aws_codebuild_source_credential_invalid_server_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCodebuildSourceCredentialInvalidServerTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCodebuildSourceCredentialInvalidServerTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCodebuildSourceCredentialInvalidServerTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCodebuildSourceCredentialInvalidServerTypeRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as server_type`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
