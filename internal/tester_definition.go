package internal

import (
	"github.com/codecrafters-io/tester-utils/tester_definition"
)

var testerDefinition = tester_definition.TesterDefinition{
	AntiCheatTestCases: []tester_definition.TestCase{},
	ExecutableFileName: "script.sh",
	TestCases: []tester_definition.TestCase{
		{
			Slug:     "init",
			TestFunc: testInit,
		},
		{
			Slug:     "second",
			TestFunc: testInit,
		},
		{
			Slug:     "ext1-stage1",
			TestFunc: testInit,
		},
		{
			Slug:     "ext1-stage2",
			TestFunc: testInit,
		},
		{
			Slug:     "ext2-stage1",
			TestFunc: testInit,
		},
		{
			Slug:     "ext2-stage2",
			TestFunc: testInit,
		},
	},
}
