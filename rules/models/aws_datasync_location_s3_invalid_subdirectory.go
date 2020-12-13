// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsDatasyncLocationS3InvalidSubdirectoryRule checks the pattern is valid
type AwsDatasyncLocationS3InvalidSubdirectoryRule struct {
	resourceType  string
	attributeName string
	max           int
	pattern       *regexp.Regexp
}

// NewAwsDatasyncLocationS3InvalidSubdirectoryRule returns new rule with default attributes
func NewAwsDatasyncLocationS3InvalidSubdirectoryRule() *AwsDatasyncLocationS3InvalidSubdirectoryRule {
	return &AwsDatasyncLocationS3InvalidSubdirectoryRule{
		resourceType:  "aws_datasync_location_s3",
		attributeName: "subdirectory",
		max:           4096,
		pattern:       regexp.MustCompile(`^[a-zA-Z0-9_\-\+\./\(\)\p{Zs}]*$`),
	}
}

// Name returns the rule name
func (r *AwsDatasyncLocationS3InvalidSubdirectoryRule) Name() string {
	return "aws_datasync_location_s3_invalid_subdirectory"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsDatasyncLocationS3InvalidSubdirectoryRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsDatasyncLocationS3InvalidSubdirectoryRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsDatasyncLocationS3InvalidSubdirectoryRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsDatasyncLocationS3InvalidSubdirectoryRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"subdirectory must be 4096 characters or less",
					attribute.Expr,
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[a-zA-Z0-9_\-\+\./\(\)\p{Zs}]*$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
