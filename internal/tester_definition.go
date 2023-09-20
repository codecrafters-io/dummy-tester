package internal

import (
	testerutils "github.com/codecrafters-io/tester-utils"
)

var testerDefinition = testerutils.TesterDefinition{
	AntiCheatTestCases: []testerutils.TestCase{},
	ExecutableFileName: "script.sh",
	TestCases: []testerutils.TestCase{
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
