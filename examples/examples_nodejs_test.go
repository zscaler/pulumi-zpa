package examples

import (
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

func getJSBaseOptions(t *testing.T) integration.ProgramTestOptions {
	base := getBaseOptions()
	baseJS := base.With(integration.ProgramTestOptions{
		Dependencies: []string{
			"@bdzscaler/pulumi-zpa",
		},
	})

	return baseJS
}

// func TestAccZPASegmentGroupTS(t *testing.T) {
// 	skipNoZPACreds(t)
// 	test := getJSBaseOptions(t).
// 		With(integration.ProgramTestOptions{
// 			Dir: path.Join(getCwd(t), "./ts/zpa_segment_group"),
// 		})

// 	integration.ProgramTest(t, &test)
// }
