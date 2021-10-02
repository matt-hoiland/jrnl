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

// entryAddCmd represents the add command
var entryAddCmd = &cobra.Command{
	Use:   "add <title> [-t <tag> ...]",
	Short: "Add a new journal by title with optional tags",
	Long: `Entry add (jrnl entry add) creates a new journal entry preconfigured with the
optional tag set, timestamped for the moment of its creation, and titled with
the provide 'title'. The produced file will have time stamped, abbreviated,
sanitized short name derived from the title.`,
	Args: cobra.ExactArgs(1),
	Run:  runEntryAddCmd,
}

func init() {
	entryCmd.AddCommand(entryAddCmd)
}

func runEntryAddCmd(cmd *cobra.Command, args []string) {
	fmt.Printf("TODO: implement entry add, args: %v, tags: %v\n", args, tags)
}
