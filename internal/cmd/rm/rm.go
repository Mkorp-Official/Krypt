// Copyright © 2021 Rak Laptudirm <raklaptudirm@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package rm

import (
	"fmt"

	"github.com/MakeNowJust/heredoc"
	"github.com/raklaptudirm/krypt/internal/auth"
	"github.com/raklaptudirm/krypt/internal/cmdutil"
	"github.com/raklaptudirm/krypt/pkg/pass"
	"github.com/spf13/cobra"
)

type RmOptions struct {
	Checksum []byte
	Creds    *auth.Creds
	Pass     pass.Manager
}

func NewCmd(c *cmdutil.Context) *cobra.Command {
	opts := &RmOptions{
		Creds: c.Creds,
		Pass:  c.PassManager,
	}

	var cmd = &cobra.Command{
		Use:   "rm [name]",
		Short: "remove a password from krypt",
		Args:  cobra.ExactArgs(1),
		Long: heredoc.Doc(`
			Logout clears the file which stores your database key,
			so that accessing the passwords requires logging in with
			the master password.
		`),
		RunE: func(cmd *cobra.Command, args []string) error {
			pass, err := pass.GetS(opts.Pass, args[0], c.Creds.Key)
			if err != nil {
				return err
			}

			opts.Checksum = pass.Checksum
			return rm(opts)
		},
	}

	return cmd
}

func rm(opts *RmOptions) error {
	if !opts.Creds.LoggedIn() {
		return cmdutil.ErrNoLogin
	}

	err := opts.Pass.Delete(opts.Checksum)
	if err != nil {
		return err
	}

	fmt.Println("Deleted password.")
	return nil
}
