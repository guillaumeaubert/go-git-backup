package gitbackup

import (
	"github.com/davecgh/go-spew/spew"
	"testing"
)

type testdata struct {
	repositoryName string
	target         Target
	expectedResult bool
}

var tests = []testdata{
	{
		"reponame",
		Target{},
		true,
	},
	{
		"reponame",
		Target{Skip: "^repo"},
		false,
	},
	{
		"reponame",
		Target{Skip: "^test"},
		true,
	},
	{
		"reponame",
		Target{Only: "^test"},
		false,
	},
	{
		"reponame",
		Target{Only: "^repo"},
		true,
	},
}

func TestIncludeRepository(t *testing.T) {
	for _, test := range tests {
		result := includeRepository(
			test.repositoryName,
			test.target,
		)
		if result != test.expectedResult {
			t.Errorf(
				"Repository: %s\nTarget: %s\nExpected: %t\nGot: %t",
				test.repositoryName,
				spew.Sdump(test.target),
				test.expectedResult,
				result,
			)
		}
	}
}
