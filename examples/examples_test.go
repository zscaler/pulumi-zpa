// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.

package examples

import (
	"os"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

func skipNoZPACreds(t *testing.T) {
	token := os.Getenv("ZPA_CLIENT_ID")
	if token == "" {
		t.Skipf("Skipping test due to missing ZPA_CLIENT_ID variable")
	}
	baseUrl := os.Getenv("ZPA_CLIENT_SECRET")
	if baseUrl == "" {
		t.Skipf("Skipping test due to missing ZPA_CLIENT_SECRET variable")
	}
	orgName := os.Getenv("ZPA_CUSTOMER_ID")
	if orgName == "" {
		t.Skipf("Skipping test due to missing ZPA_CUSTOMER_ID variable")
	}
}

func getCwd(t *testing.T) string {
	cwd, err := os.Getwd()
	if err != nil {
		t.FailNow()
	}

	return cwd
}

func getBaseOptions() integration.ProgramTestOptions {
	return integration.ProgramTestOptions{
		ExpectRefreshChanges: true,
	}
}
