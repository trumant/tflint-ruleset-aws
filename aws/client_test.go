package aws

import (
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
	awsbase "github.com/hashicorp/aws-sdk-go-base"
	"github.com/mitchellh/go-homedir"
)

func Test_getBaseConfig(t *testing.T) {
	home, err := homedir.Expand("~/")
	if err != nil {
		t.Fatal(err)
	}

	cases := []struct {
		Name     string
		Creds    Credentials
		Expected *awsbase.Config
	}{
		{
			Name: "static credentials",
			Creds: Credentials{
				AccessKey: "AWS_ACCESS_KEY",
				SecretKey: "AWS_SECRET_KEY",
				Region:    "us-east-1",
			},
			Expected: &awsbase.Config{
				AccessKey: "AWS_ACCESS_KEY",
				SecretKey: "AWS_SECRET_KEY",
				Region:    "us-east-1",
			},
		},
		{
			Name: "shared credentials",
			Creds: Credentials{
				Profile:   "default",
				CredsFile: "~/.aws/creds",
				Region:    "us-east-1",
			},
			Expected: &awsbase.Config{
				Profile:       "default",
				CredsFilename: filepath.Join(home, ".aws", "creds"),
				Region:        "us-east-1",
			},
		},
	}

	for _, tc := range cases {
		base, err := getBaseConfig(tc.Creds)
		if err != nil {
			t.Fatalf("Failed `%s` test: Unexpected error occurred: %s", tc.Name, err)
		}
		if !cmp.Equal(tc.Expected, base) {
			t.Fatalf("Failed `%s` test: Diff=%s", tc.Name, cmp.Diff(tc.Expected, base))
		}
	}
}