// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsLightsailStaticIPAttachmentInvalidInstanceNameRule checks the pattern is valid
type AwsLightsailStaticIPAttachmentInvalidInstanceNameRule struct {
	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAwsLightsailStaticIPAttachmentInvalidInstanceNameRule returns new rule with default attributes
func NewAwsLightsailStaticIPAttachmentInvalidInstanceNameRule() *AwsLightsailStaticIPAttachmentInvalidInstanceNameRule {
	return &AwsLightsailStaticIPAttachmentInvalidInstanceNameRule{
		resourceType:  "aws_lightsail_static_ip_attachment",
		attributeName: "instance_name",
		pattern:       regexp.MustCompile(`^\w[\w\-]*\w$`),
	}
}

// Name returns the rule name
func (r *AwsLightsailStaticIPAttachmentInvalidInstanceNameRule) Name() string {
	return "aws_lightsail_static_ip_attachment_invalid_instance_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsLightsailStaticIPAttachmentInvalidInstanceNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsLightsailStaticIPAttachmentInvalidInstanceNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsLightsailStaticIPAttachmentInvalidInstanceNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsLightsailStaticIPAttachmentInvalidInstanceNameRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^\w[\w\-]*\w$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
