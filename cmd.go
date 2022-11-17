// Copyright 2022 bonzai
// SPDX-License-Identifier: Apache-2.0

// Package example provides the Bonzai command branch of the same name.
package example

import (
	"github.com/rwx-yxu/bonzai/own"
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/conf"
	"github.com/rwxrob/help"
	"github.com/rwxrob/vars"
)

// Most Cmds that make use of the conf and vars branches will want to
// call SoftInit in order to create the persistence layers or whatever
// else is needed to initialize their use. This cannot be done
// automatically from these imported modules because Cmd authors may
// with to change the default values before calling SoftInit and
// committing them.

func init() {
	Z.Conf.SoftInit()
	Z.Vars.SoftInit()
}

// Cmd provides a Bonzai branch command that can be composed into Bonzai
// trees or used as a standalone with light wrapper (see cmd/).
var Cmd = &Z.Cmd{

	Name:      `bonzai`,
	Summary:   `an example of Bonzai composite command tree`,
	Version:   `v0.4.1`,
	Copyright: `Copyright 2021 Yongle Xu`,
	License:   `Apache-2.0`,
	Site:      `yonglexu.dev`,
	Source:    `git@github.com:rwx-yxu/bonzai.git`,
	Issues:    `github.com/rwx-yxu/bonzai/issues`,

	// Composite commands, local and external, all have their own names
	// that are added to the command tree depending on where they are
	// composed.

	Commands: []*Z.Cmd{

		// standard external branch imports (see rwxrob/{help,conf,vars})
		help.Cmd, conf.Cmd, vars.Cmd,

		// local commands (in this module)
		own.Cmd,
	},

	// Add custom BonzaiMark template extensions (or overwrite existing ones).

	Description: `
		The {{cmd .Name}} command sets all cached variables to their initial
		values. Any variable name from {{cmd "conf"}} will be used to
		initialize if defined.  Otherwise, the following hard-coded package
		globals will be used instead:
			`,
}
