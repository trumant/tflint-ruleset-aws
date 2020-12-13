// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsMacieS3BucketAssociationInvalidMemberAccountIDRule checks the pattern is valid
type AwsMacieS3BucketAssociationInvalidMemberAccountIDRule struct {
	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAwsMacieS3BucketAssociationInvalidMemberAccountIDRule returns new rule with default attributes
func NewAwsMacieS3BucketAssociationInvalidMemberAccountIDRule() *AwsMacieS3BucketAssociationInvalidMemberAccountIDRule {
	return &AwsMacieS3BucketAssociationInvalidMemberAccountIDRule{
		resourceType:  "aws_macie_s3_bucket_association",
		attributeName: "member_account_id",
		pattern:       regexp.MustCompile(`^[0-9]{12}$`),
	}
}

// Name returns the rule name
func (r *AwsMacieS3BucketAssociationInvalidMemberAccountIDRule) Name() string {
	return "aws_macie_s3_bucket_association_invalid_member_account_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsMacieS3BucketAssociationInvalidMemberAccountIDRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsMacieS3BucketAssociationInvalidMemberAccountIDRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsMacieS3BucketAssociationInvalidMemberAccountIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsMacieS3BucketAssociationInvalidMemberAccountIDRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[0-9]{12}$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
