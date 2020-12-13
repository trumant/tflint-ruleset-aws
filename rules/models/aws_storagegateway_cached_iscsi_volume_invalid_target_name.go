// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsStoragegatewayCachedIscsiVolumeInvalidTargetNameRule checks the pattern is valid
type AwsStoragegatewayCachedIscsiVolumeInvalidTargetNameRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsStoragegatewayCachedIscsiVolumeInvalidTargetNameRule returns new rule with default attributes
func NewAwsStoragegatewayCachedIscsiVolumeInvalidTargetNameRule() *AwsStoragegatewayCachedIscsiVolumeInvalidTargetNameRule {
	return &AwsStoragegatewayCachedIscsiVolumeInvalidTargetNameRule{
		resourceType:  "aws_storagegateway_cached_iscsi_volume",
		attributeName: "target_name",
		max:           200,
		min:           1,
		pattern:       regexp.MustCompile(`^[-\.;a-z0-9]+$`),
	}
}

// Name returns the rule name
func (r *AwsStoragegatewayCachedIscsiVolumeInvalidTargetNameRule) Name() string {
	return "aws_storagegateway_cached_iscsi_volume_invalid_target_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsStoragegatewayCachedIscsiVolumeInvalidTargetNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsStoragegatewayCachedIscsiVolumeInvalidTargetNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsStoragegatewayCachedIscsiVolumeInvalidTargetNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsStoragegatewayCachedIscsiVolumeInvalidTargetNameRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"target_name must be 200 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"target_name must be 1 characters or higher",
					attribute.Expr,
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[-\.;a-z0-9]+$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
