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

package list

import (
	"fmt"

	"github.com/MakeNowJust/heredoc"
	"github.com/raklaptudirm/krypt/internal/auth"
	"github.com/raklaptudirm/krypt/pkg/cmdutil"
	"github.com/raklaptudirm/krypt/pkg/pass"
	"github.com/spf13/cobra"
)

type ListOptions struct {
	Auth    *auth.Auth
	Filters []pass.Filter
}

func NewCmd(f *cmdutil.Factory) *cobra.Command {
	opts := &ListOptions{
		Auth: f.Auth,
	}

	var cmd = &cobra.Command{
		Use:   "list [name]",
		Short: "un-encrypt and fetch a password from krypt using the filters",
		Args:  cobra.NoArgs,
		Long: heredoc.Doc(`
			List all the passwords which match the provided filters. If no filters
			are provided, all the passwords are listed.
		`),
		RunE: func(cmd *cobra.Command, args []string) error {
			// TODO: map flags to password filters
			return list(opts)
		},
	}

	return cmd
}

func list(opts *ListOptions) error {
	passwords, err := pass.Get(opts.Auth.Key, opts.Filters...)
	if err != nil {
		return err
	}

	for _, password := range passwords {
		fmt.Println(password.String())
	}
	return nil
}