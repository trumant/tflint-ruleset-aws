// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsSsmPatchBaselineInvalidApprovedPatchesComplianceLevelRule checks the pattern is valid
type AwsSsmPatchBaselineInvalidApprovedPatchesComplianceLevelRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsSsmPatchBaselineInvalidApprovedPatchesComplianceLevelRule returns new rule with default attributes
func NewAwsSsmPatchBaselineInvalidApprovedPatchesComplianceLevelRule() *AwsSsmPatchBaselineInvalidApprovedPatchesComplianceLevelRule {
	return &AwsSsmPatchBaselineInvalidApprovedPatchesComplianceLevelRule{
		resourceType:  "aws_ssm_patch_baseline",
		attributeName: "approved_patches_compliance_level",
		enum: []string{
			"CRITICAL",
			"HIGH",
			"MEDIUM",
			"LOW",
			"INFORMATIONAL",
			"UNSPECIFIED",
		},
	}
}

// Name returns the rule name
func (r *AwsSsmPatchBaselineInvalidApprovedPatchesComplianceLevelRule) Name() string {
	return "aws_ssm_patch_baseline_invalid_approved_patches_compliance_level"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSsmPatchBaselineInvalidApprovedPatchesComplianceLevelRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSsmPatchBaselineInvalidApprovedPatchesComplianceLevelRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSsmPatchBaselineInvalidApprovedPatchesComplianceLevelRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSsmPatchBaselineInvalidApprovedPatchesComplianceLevelRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as approved_patches_compliance_level`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
