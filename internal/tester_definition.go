package internal

import (
	testerutils "github.com/codecrafters-io/tester-utils"
)

var testerDefinition = testerutils.TesterDefinition{
	AntiCheatStages:    []testerutils.Stage{},
	ExecutableFileName: "script.sh",
	Stages: []testerutils.Stage{
		{
			Number:                  1,
			Slug:                    "init",
			Title:                   "The first stage",
			TestFunc:                testInit,
			ShouldRunPreviousStages: true,
		},
		{
			Number:                  2,
			Slug:                    "second",
			Title:                   "The second stage",
			TestFunc:                testInit,
			ShouldRunPreviousStages: true,
		},
		{
			Number:                  3,
			Slug:                    "ext1-stage1",
			Title:                   "Start with ext1",
			TestFunc:                testInit,
			ShouldRunPreviousStages: true,
		},
		{
			Number:                  4,
			Slug:                    "ext1-stage2",
			Title:                   "Finish with ext1",
			TestFunc:                testInit,
			ShouldRunPreviousStages: true,
		},
		{
			Number:                  5,
			Slug:                    "ext2-stage1",
			Title:                   "Start with ext2",
			TestFunc:                testInit,
			ShouldRunPreviousStages: true,
		},
		{
			Number:                  6,
			Slug:                    "ext2-stage2",
			Title:                   "Finish with ext1 + ext2",
			TestFunc:                testInit,
			ShouldRunPreviousStages: true,
		},
	},
}
