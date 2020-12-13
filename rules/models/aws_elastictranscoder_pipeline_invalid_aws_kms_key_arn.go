// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsElastictranscoderPipelineInvalidAwsKmsKeyArnRule checks the pattern is valid
type AwsElastictranscoderPipelineInvalidAwsKmsKeyArnRule struct {
	resourceType  string
	attributeName string
	max           int
}

// NewAwsElastictranscoderPipelineInvalidAwsKmsKeyArnRule returns new rule with default attributes
func NewAwsElastictranscoderPipelineInvalidAwsKmsKeyArnRule() *AwsElastictranscoderPipelineInvalidAwsKmsKeyArnRule {
	return &AwsElastictranscoderPipelineInvalidAwsKmsKeyArnRule{
		resourceType:  "aws_elastictranscoder_pipeline",
		attributeName: "aws_kms_key_arn",
		max:           255,
	}
}

// Name returns the rule name
func (r *AwsElastictranscoderPipelineInvalidAwsKmsKeyArnRule) Name() string {
	return "aws_elastictranscoder_pipeline_invalid_aws_kms_key_arn"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsElastictranscoderPipelineInvalidAwsKmsKeyArnRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsElastictranscoderPipelineInvalidAwsKmsKeyArnRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsElastictranscoderPipelineInvalidAwsKmsKeyArnRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsElastictranscoderPipelineInvalidAwsKmsKeyArnRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"aws_kms_key_arn must be 255 characters or less",
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
