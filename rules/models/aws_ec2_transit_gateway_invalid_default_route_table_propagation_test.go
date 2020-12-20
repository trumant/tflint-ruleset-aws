// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"testing"

	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_AwsEc2TransitGatewayInvalidDefaultRouteTablePropagationRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "It includes invalid characters",
			Content: `
resource "aws_ec2_transit_gateway" "foo" {
	default_route_table_propagation = "disabled"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewAwsEc2TransitGatewayInvalidDefaultRouteTablePropagationRule(),
					Message: `"disabled" is an invalid value as default_route_table_propagation`,
				},
			},
		},
		{
			Name: "It is valid",
			Content: `
resource "aws_ec2_transit_gateway" "foo" {
	default_route_table_propagation = "disable"
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewAwsEc2TransitGatewayInvalidDefaultRouteTablePropagationRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}