package internal

import (
	"github.com/codecrafters-io/tester-utils/tester_definition"
)

var testerDefinition = tester_definition.TesterDefinition{
	AntiCheatTestCases: []tester_definition.TestCase{},
	ExecutableFileName: "script.sh",
	TestCases: []tester_definition.TestCase{
		{
			Slug:     "ah7",
			TestFunc: testInit,
		},
		{
			Slug:     "lr7",
			TestFunc: testInit,
		},
		{
			Slug:     "qh7",
			TestFunc: testInit,
		},
		{
			Slug:     "wd5",
			TestFunc: testInit,
		},
		{
			Slug:     "ae0",
			TestFunc: testInit,
		},
		{
			Slug:     "um4",
			TestFunc: testInit,
		},
	},
}
