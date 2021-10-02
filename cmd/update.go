/*
Copyright © 2021 Matt Hoiland <matthew.proctor.hoiland@gmail.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update metadata store to reflect journal contents",
	Long: `Update (jrnl update) will scrape the journal directory for untracked or out of
date entries, updating the metadata stored in the '.jrnl/' directory.`,
	Args: cobra.NoArgs,
	Run:  runUpdateCmd,
}

func init() {
	rootCmd.AddCommand(updateCmd)
}

func runUpdateCmd(cmd *cobra.Command, args []string) {
	fmt.Printf("TODO: implement update, args: %v\n", args)
}
