// Copyright 2015 Google Inc. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package common

import (
	"blueprint"
)

var (
	pctx = blueprint.NewPackageContext("android/soong/common")

	cpPreserveSymlinks = pctx.VariableConfigMethod("cpPreserveSymlinks",
		Config.CpPreserveSymlinksFlags)

	// A phony rule that is not the built-in Ninja phony rule.  The built-in
	// phony rule has special behavior that is sometimes not desired.  See the
	// Ninja docs for more details.
	Phony = pctx.StaticRule("Phony",
		blueprint.RuleParams{
			Command:     "# phony $out",
			Description: "phony $out",
		})

	// GeneratedFile is a rule for indicating that a given file was generated
	// while running soong.  This allows the file to be cleaned up if it ever
	// stops being generated by soong.
	GeneratedFile = pctx.StaticRule("GeneratedFile",
		blueprint.RuleParams{
			Command:     "# generated $out",
			Description: "generated $out",
			Generator:   true,
		})

	// A copy rule.
	Cp = pctx.StaticRule("Cp",
		blueprint.RuleParams{
			Command:     "cp $cpPreserveSymlinks $cpFlags $in $out",
			Description: "cp $out",
		},
		"cpFlags")

	// A symlink rule.
	Symlink = pctx.StaticRule("Symlink",
		blueprint.RuleParams{
			Command:     "ln -f -s $fromPath $out",
			Description: "symlink $out",
		},
		"fromPath")
)
